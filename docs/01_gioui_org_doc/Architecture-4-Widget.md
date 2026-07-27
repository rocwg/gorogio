**第4章：Widget（控件）**

**原文地址**：https://gioui.org/doc/architecture/widget

---

### 完整中文翻译

# Widget

# Widget 可复用且可组合的部分

我们已经多次提到控件（widgets）。原则上，控件是可组合、可绘制的 UI 元素，并且可能对输入做出反应。更具体地说：

* 它们从 [Source](https://gioui.org/app#FrameEvent.Source) 获取输入。
* 它们可能持有一些状态。
* 它们根据约束计算自己的尺寸。
* 它们把自己绘制到一个 [op.Ops](https://gioui.org/op#Ops) 列表中。

按惯例，控件有一个 `Layout` 方法，完成以上所有事情。有些控件有单独的方法来查询它们的状态，或者把事件传递回程序（例如 [Clickable.Clicked](https://gioui.org/widget#Clickable.Clicked)）。

有些控件有多种视觉表现。例如，有状态的 [Clickable](https://gioui.org/widget#Clickable) 被用作 [按钮](https://gioui.org/widget/material#ButtonStyle.Layout) 和 [图标按钮](https://gioui.org/widget/material#IconButtonStyle.Layout) 的基础。实际上，[material 包](https://gioui.org/widget/material) 只实现了 [Material Design](https://material.io)，并打算由其他包来补充实现不同的设计。

## Context（上下文）

为了用这些原语构建更复杂的 UI，我们需要一些结构来以可组合的方式描述布局。

可以静态指定布局，但显示尺寸变化很大，所以我们需要能够动态计算布局——也就是约束可用的显示尺寸，然后计算其余布局。我们还需要一种舒适的方式把事件传递通过组合结构，同样我们也需要一种方式把 [op.Ops](https://gioui.org/op#Ops) 传递通过系统。

[layout.Context](https://gioui.org/layout#Context) 方便地把这些方面捆绑在一起。它携带了几乎所有布局和控件都需要的状态。

总结一下术语：

* [Constraints](https://gioui.org/layout#Context.Constraints) 是控件的“传入”参数。约束持有控件的最大（和最小）尺寸。
* [Ops](https://gioui.org/layout#Context.Ops) 持有生成的绘制操作。
* [Events](https://gioui.org/layout#Context.Events) 持有自上次绘制操作以来生成的事件。

按惯例，接受 `layout.Context` 的函数返回 [layout.Dimensions](https://gioui.org/layout#Dimensions)，它同时提供了已布局控件的尺寸，以及该控件内任何文本内容的基线。

```go
var window app.Window
window.Option(app.Title(title))

var ops op.Ops
for {
	switch e := window.Event().(type) {
	case app.DestroyEvent:
		// 窗口被关闭了。
		return e.Err
	case app.FrameEvent:
		// 为新帧重置 layout.Context。
		gtx := app.NewContext(&ops, e)

		// 根据 e.Queue 中的事件，把状态绘制进 ops。
		draw(gtx)

		// 更新显示。
		e.Frame(gtx.Ops)
	}
}
```

## Custom（自定义）

作为例子，这里是如何实现一个非常简单的按钮。

我们先从绘制它开始：

```go
type ButtonVisual struct {
	pressed bool
}

func (b *ButtonVisual) Layout(gtx layout.Context) layout.Dimensions {
	col := color.NRGBA{R: 0x80, A: 0xFF}
	if b.pressed {
		col = color.NRGBA{G: 0x80, A: 0xFF}
	}
	return drawSquare(gtx.Ops, col)
}

func drawSquare(ops *op.Ops, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{}.Add(ops)
	return layout.Dimensions{Size: image.Pt(100, 100)}
}
```

然后处理指针点击：

```go
type Button struct {
	pressed bool
}

func (b *Button) Layout(gtx layout.Context) layout.Dimensions {
	// 限制指针事件的区域。
	area := clip.Rect(image.Rect(0, 0, 100, 100)).Push(gtx.Ops)

	event.Op(gtx.Ops, b)

	// 这里我们循环处理与这个按钮关联的所有事件。
	for {
		ev, ok := gtx.Event(pointer.Filter{
			Target: b,
			Kinds:  pointer.Press | pointer.Release,
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
			b.pressed = true
		case pointer.Release:
			b.pressed = false
		}
	}

	area.Pop()

	// 绘制按钮。
	col := color.NRGBA{R: 0x80, A: 0xFF}
	if b.pressed {
		col = color.NRGBA{G: 0x80, A: 0xFF}
	}
	return drawSquare(gtx.Ops, col)
}
```

---

### 深度解读

**1. Widget 在 Gio 中的真实定义**

Gio 的 Widget **不是**一个继承自某个基类的对象，也不是必须实现某个接口的结构体。  
它只是一个**约定**：

> 一个接收 `layout.Context`、返回 `layout.Dimensions` 的方法（通常叫 `Layout`）。

这个约定让所有东西都能互相组合：按钮、列表项、整页布局、自定义复杂组件，都可以用同样的方式嵌套。

**2. layout.Context 是真正的核心**

从第1章到第3章，我们一直在手动传递 `*op.Ops` 和 `input.Source`。  
到了第4章，官方正式引入了 `layout.Context`，它把三样最重要的东西打包在一起：

| 字段                | 含义                    | 来源     |
| ------------------- | ----------------------- | -------- |
| `Constraints`       | 父布局给的最小/最大尺寸 | 布局系统 |
| `Ops`               | 当前操作列表            | 第1章    |
| `Events` / `Source` | 输入事件源              | 第3章    |
| `Metric`            | 单位转换（Dp/Sp → Px）  | 第7章    |

标准创建方式是：

```go
gtx := app.NewContext(&ops, e)  // e 是 FrameEvent
```

之后所有控件和布局都只接收这一个 `gtx`。

**3. 返回值 layout.Dimensions 的意义**

```go
type Dimensions struct {
    Size     image.Point  // 实际占用的像素尺寸
    Baseline int          // 文本基线（用于垂直对齐）
}
```

为什么要返回尺寸？因为**父布局需要知道子控件实际占了多大**，才能正确放置下一个兄弟控件，或者计算自己的最终尺寸。

这是约束布局（Constraint-based Layout）的核心：父给约束 → 子返回实际尺寸。

**4. 状态与视觉分离（非常重要的设计）**

官方特别强调：

- `widget.Clickable` 只管理“是否被点击”的状态和事件
- `material.Button`、`material.IconButton` 只负责**怎么画**

这种分离带来巨大好处：
- 同一套状态可以套用不同主题（Material、自定义主题、暗色主题……）
- 你可以自己写完全不同的视觉实现，而不用重新发明点击逻辑

这是 Gio 与很多保留模式框架最大的不同：状态是你自己的，视觉是可插拔的。

**5. 自定义控件的标准模板**

几乎所有自定义控件都遵循这个结构：

```go
func (w *MyWidget) Layout(gtx layout.Context) layout.Dimensions {
    // 1. 定义自己的输入区域（通常用 clip）
    defer clip.Rect{...}.Push(gtx.Ops).Pop()
    
    // 2. 注册输入兴趣
    event.Op(gtx.Ops, w)
    
    // 3. 处理本帧到达的事件
    for {
        ev, ok := gtx.Event(...)
        if !ok { break }
        // 更新 w 的内部状态
    }
    
    // 4. 根据当前状态进行绘制
    // ...
    
    // 5. 返回自己实际占用的尺寸
    return layout.Dimensions{Size: ...}
}
```

**6. 为什么官方示例里先画再处理事件也可以？**

在即时模式里，**顺序其实不重要**。因为：
- 事件是上一帧到这一帧之间积累的
- 绘制的是“当前状态”
- 下一帧才会看到状态变化后的结果

所以“先处理事件再画”或“先画再处理事件”在视觉上几乎没区别（最多延迟一帧）。但推荐先处理事件，逻辑更清晰。

**7. 实际开发中的关键建议**

| 点       | 建议                                          |
| -------- | --------------------------------------------- |
| 状态存放 | 控件结构体自己持有（`*Button`）               |
| Tag      | 直接用控件指针本身（`event.Op(ops, b)`）      |
| 尺寸计算 | 尽量根据 Constraints 计算，而不是写死 100x100 |
| 组合     | 大控件内部调用小控件的 Layout                 |
| 主题     | 视觉部分尽量抽到 material 或自己的主题包      |

**8. 与前后章节的关系**

- 第1章：提供了 `FrameEvent` 和 `op.Ops`
- 第2章：提供了如何画
- 第3章：提供了如何收输入
- **第4章：把画 + 输入 + 尺寸计算打包成可复用的 Widget**
- 第5章：提供各种布局容器，把多个 Widget 组合起来

Widget 是连接底层操作与高层布局的桥梁。

---

**本章核心记忆点**

> **Widget = 接收 layout.Context、处理输入、绘制自己、返回 Dimensions 的约定**  
> **layout.Context 是所有布局和控件的统一上下文**  
> **状态与视觉分离是 Gio 的重要设计**  
> **自定义控件的标准流程：注册输入 → 处理事件 → 绘制 → 返回尺寸**

---

准备好后，回复“继续第5章”，我会继续给出 Layout 章节的完整翻译与深度解读。