继续按照 **【英文原文】 $\rightarrow$ 【精准逐字翻译】** $\rightarrow$ **【专业术语与 Gio 架构机制剖析】** 的架构为你深度拆解 Gio 官方 Learn 教程的第二章：**Split Widget（自定义分割/分栏控件）**。

这一章是 Gio 进阶的石碑，展示了 Gio **如何自定义控件、如何裁剪输入区域（Clip Area）以及如何处理指针拖拽事件（Pointer Events）**。



# Split Widget（自定义分割/分栏控件）

##### Tailoring things to your own needs (按需定制)



## 1. 静态分栏实现 (SplitVisual)

Sometimes there’s a need for writing a custom widget or layout.
有时需要编写自定义控件或布局。

To implement rendering of children, we can use:
为了实现子控件的渲染，我们可以使用：

```go
type SplitVisual struct{}

func (s SplitVisual) Layout(gtx layout.Context, left, right layout.Widget) layout.Dimensions {
	leftsize := gtx.Constraints.Min.X / 2
	rightsize := gtx.Constraints.Min.X - leftsize

	{
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(leftsize, gtx.Constraints.Max.Y))
		left(gtx)
	}

	{
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(rightsize, gtx.Constraints.Max.Y))
		trans := op.Offset(image.Pt(leftsize, 0)).Push(gtx.Ops)
		right(gtx)
		trans.Pop()
	}

	return layout.Dimensions{Size: gtx.Constraints.Max}
}
```

Then we can use the widget like:
然后我们可以像这样使用该控件：

```go
func exampleSplitVisual(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return SplitVisual{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Left", red)
	}, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Right", blue)
	})
}

func FillWithLabel(gtx layout.Context, th *material.Theme, text string, backgroundColor color.NRGBA) layout.Dimensions {
	ColorBox(gtx, gtx.Constraints.Max, backgroundColor)
	return layout.Center.Layout(gtx, material.H3(th, text).Layout)
}
```



## 2. 加入比例调节 (Ratio)

Let’s make the ratio adjustable. We should try to make zero values useful, in this case `0` could mean that it’s split in the center.
让我们让分割比例可调。我们应该尽量让零值具有实际意义，在此例中，0 可以表示居中分割。

```go
type SplitRatio struct {
	// Ratio keeps the current layout.
	// 0 is center, -1 completely to the left, 1 completely to the right.
	Ratio float32
}

func (s SplitRatio) Layout(gtx layout.Context, left, right layout.Widget) layout.Dimensions {
	proportion := (s.Ratio + 1) / 2
	leftsize := int(proportion * float32(gtx.Constraints.Max.X))

	rightoffset := leftsize
	rightsize := gtx.Constraints.Max.X - rightoffset

	{
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(leftsize, gtx.Constraints.Max.Y))
		left(gtx)
	}

	{
		trans := op.Offset(image.Pt(rightoffset, 0)).Push(gtx.Ops)
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(rightsize, gtx.Constraints.Max.Y))
		right(gtx)
		trans.Pop()
	}

	return layout.Dimensions{Size: gtx.Constraints.Max}
}
```

The usage code would look like:
使用代码将如下所示：

```go
func exampleSplitRatio(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return SplitRatio{Ratio: -0.3}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Left", red)
	}, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Right", blue)
	})
}
```



## 3. 交互式拖拽实现 (Interactive & Result)

#### **交互式（Interactive）**

To make it more useful we could make the split draggable.
为了让它更有用，我们可以让分割线可拖拽。

Because we also need to have an area designated for moving the split, let’s add a bar into the center:
因为我们还需要为移动分割线指定一个区域，让我们在中心添加一条分割条（bar）：

```go
bar := gtx.Dp(s.Bar)
if bar <= 1 {
	bar = gtx.Dp(defaultBarWidth)
}

proportion := (s.Ratio + 1) / 2
leftsize := int(proportion*float32(gtx.Constraints.Max.X) - float32(bar))

rightoffset := leftsize + bar
rightsize := gtx.Constraints.Max.X - rightoffset
```

Now we need to store our interactive state:
现在我们需要存储我们的交互状态：

```go
type Split struct {
	// Ratio keeps the current layout.
	// 0 is center, -1 completely to the left, 1 completely to the right.
	Ratio float32
	// Bar is the width for resizing the layout
	Bar unit.Dp

	drag   bool
	dragID pointer.ID
	dragX  float32
}
```

And then we need to handle input events:
接着我们需要处理输入事件：

```go
barRect := image.Rect(leftsize, 0, rightoffset, gtx.Constraints.Max.X)
area := clip.Rect(barRect).Push(gtx.Ops)

// register for input
event.Op(gtx.Ops, s)
pointer.CursorColResize.Add(gtx.Ops)

for {
	ev, ok := gtx.Event(pointer.Filter{
		Target: s,
		Kinds:  pointer.Press | pointer.Drag | pointer.Release | pointer.Cancel,
	})
	if !ok {
		break
	}

	e, ok := ev.(pointer.Event)
	if !ok {
		continue
	}

	switch e.Kind {
	case pointer.Press:
		if s.drag {
			break
		}

		s.dragID = e.PointerID
		s.dragX = e.Position.X
		s.drag = true

	case pointer.Drag:
		if s.dragID != e.PointerID {
			break
		}

		deltaX := e.Position.X - s.dragX
		s.dragX = e.Position.X

		deltaRatio := deltaX * 2 / float32(gtx.Constraints.Max.X)
		s.Ratio += deltaRatio

		if e.Priority < pointer.Grabbed {
			gtx.Execute(pointer.GrabCmd{
				Tag: s,
				ID:  s.dragID,
			})
		}

	case pointer.Release:
		fallthrough
	case pointer.Cancel:
		s.drag = false
	}
}

area.Pop()
```



#### **最终代码（Result）**

Putting the whole widget together:
将整个控件拼装在一起：

```go
type Split struct {
	// Ratio keeps the current layout.
	// 0 is center, -1 completely to the left, 1 completely to the right.
	Ratio float32
	// Bar is the width for resizing the layout
	Bar unit.Dp

	drag   bool
	dragID pointer.ID
	dragX  float32
}


const defaultBarWidth = unit.Dp(10)

func (s *Split) Layout(gtx layout.Context, left, right layout.Widget) layout.Dimensions {
	bar := gtx.Dp(s.Bar)
	if bar <= 1 {
		bar = gtx.Dp(defaultBarWidth)
	}

	proportion := (s.Ratio + 1) / 2
	leftsize := int(proportion*float32(gtx.Constraints.Max.X) - float32(bar))

	rightoffset := leftsize + bar
	rightsize := gtx.Constraints.Max.X - rightoffset

	{ // handle input
		barRect := image.Rect(leftsize, 0, rightoffset, gtx.Constraints.Max.X)
		area := clip.Rect(barRect).Push(gtx.Ops)

		// register for input
		event.Op(gtx.Ops, s)
		pointer.CursorColResize.Add(gtx.Ops)

		for {
			ev, ok := gtx.Event(pointer.Filter{
				Target: s,
				Kinds:  pointer.Press | pointer.Drag | pointer.Release | pointer.Cancel,
			})
			if !ok {
				break
			}

			e, ok := ev.(pointer.Event)
			if !ok {
				continue
			}

			switch e.Kind {
			case pointer.Press:
				if s.drag {
					break
				}

				s.dragID = e.PointerID
				s.dragX = e.Position.X
				s.drag = true

			case pointer.Drag:
				if s.dragID != e.PointerID {
					break
				}

				deltaX := e.Position.X - s.dragX
				s.dragX = e.Position.X

				deltaRatio := deltaX * 2 / float32(gtx.Constraints.Max.X)
				s.Ratio += deltaRatio

				if e.Priority < pointer.Grabbed {
					gtx.Execute(pointer.GrabCmd{
						Tag: s,
						ID:  s.dragID,
					})
				}

			case pointer.Release:
				fallthrough
			case pointer.Cancel:
				s.drag = false
			}
		}

		area.Pop()
	}

	{
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(leftsize, gtx.Constraints.Max.Y))
		left(gtx)
	}

	{
		off := op.Offset(image.Pt(rightoffset, 0)).Push(gtx.Ops)
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(rightsize, gtx.Constraints.Max.Y))
		right(gtx)
		off.Pop()
	}

	return layout.Dimensions{Size: gtx.Constraints.Max}
}
```

And an example:
使用示例：

```go
var split Split

func exampleSplit(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return split.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Left", red)
	}, func(gtx layout.Context) layout.Dimensions {
		return FillWithLabel(gtx, th, "Right", blue)
	})
}
```

---



### 💡 核心设计与 Gio 架构机制剖析

这一章揭示了 Gio 开发中极其关键的 **3 个核心概念**：

#### 1. 约束与变换隔离 (`gtx := gtx` + `op.Offset`)

在渲染左右两个子控件时，代码中两次出现了 `gtx := gtx` 局部变量拷贝：

```go
{
    gtx := gtx // 浅拷贝 Context，避免污染父级的 Constraints 状态
    gtx.Constraints = layout.Exact(image.Pt(leftsize, gtx.Constraints.Max.Y))
    left(gtx)
}
```

- **约束隔离**：通过修改副本的 `gtx.Constraints`，将子控件严格限定在计算出的 `leftsize` 或 `rightsize` 空间内。
- **平移栈 (`op.Offset`)**：绘制右侧控件时，先用 `op.Offset(image.Pt(rightoffset, 0)).Push(gtx.Ops)` 压入全局坐标偏移，绘制完毕后通过 `off.Pop()` 恢复坐标系，**确保右侧控件内部永远使用自己的局部坐标系（0, 0）进行布局**。

#### 2. 输入响应与裁剪区域 (`clip.Rect` + `event.Op` + `pointer.Filter`)

Gio 中**没有任何独立的 Event Listener 机制**，所有的事件绑定都是通过在 `Ops` 压栈操作流中“声明输入响应区域”完成的：

```go
barRect := image.Rect(leftsize, 0, rightoffset, gtx.Constraints.Max.X)
area := clip.Rect(barRect).Push(gtx.Ops) // 1. 在 GPU 指令流中定义响应热区
event.Op(gtx.Ops, s)                     // 2. 将此热区与结构体指针 Tag 's' 绑定
pointer.CursorColResize.Add(gtx.Ops)    // 3. 悬停在此热区时设置鼠标样式为左右双箭头
```

#### 3. 手势抓取与事件抢占 (`pointer.GrabCmd`)

在处理 `pointer.Drag` 时：

```go
if e.Priority < pointer.Grabbed {
    gtx.Execute(pointer.GrabCmd{
        Tag: s,
        ID:  s.dragID,
    })
}
```

- 当鼠标在分割条上按下并拖拽时，鼠标可能会因为移动过快短时间脱离 `barRect` 热区。
- 发送 `pointer.GrabCmd` 命令可以**强制将指定 `PointerID` 的焦点锁定到当前的 `s` Tag 上**，确保即使鼠标快速移出分割条，拖拽事件依然能稳定送达 `s`，防止拖拽断连。

---

