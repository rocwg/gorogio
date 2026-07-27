**第5章：Layout（布局）**

**原文地址**：https://gioui.org/doc/architecture/layout

---

### 完整中文翻译

# Layout

# Layout 把东西放在它们该在的地方

包 [gioui.org/layout](https://gioui.org/layout) 为常见的布局操作提供支持，例如间距、列表和重叠控件的堆叠。

在布局示例中，我们将使用这个 `ColorBox` 控件来可视化布局：

```go
// 测试颜色。
var (
	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
)

// ColorBox 创建一个具有指定尺寸和颜色的控件。
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}
```

## Inset（内边距）

[layout.Inset](https://gioui.org/layout#Inset) 在控件周围添加空间。

```go
func inset(gtx layout.Context) layout.Dimensions {
	// 在彼此内部绘制矩形，带有 30dp 的内边距。
	return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return ColorBox(gtx, gtx.Constraints.Max, red)
	})
}
```

## Stack（堆叠）

[layout.Stack](https://gioui.org/layout#Stack) 根据对齐方向布局重叠的子元素。堆叠布局的子元素可以是：

* [Stacked](https://gioui.org/layout#Stacked) —— 没有最小约束，最大约束是传递给 Stack.Layout 的最大约束。
* [Expanded](https://gioui.org/layout#Expanded) —— 使用最大的 Stacked 项作为最小约束，最大约束是传递给 Stack.Layout 的最大约束。

例如，这会在红色背景上绘制绿色和蓝色矩形：

```go
func stacked(gtx layout.Context) layout.Dimensions {
	return layout.Stack{}.Layout(gtx,
		// 强制控件与第二个相同大小。
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			// 这将有 100x100 的最小约束。
			return ColorBox(gtx, gtx.Constraints.Min, red)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 30), green)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(30, 100), blue)
		}),
	)
}
```

### Background（背景）

因为为控件布局背景非常频繁，所以有一个针对该场景的更高效实现，大致相当于：

```go
layout.Stack{Alignment: layout.C}.Layout(gtx,
	layout.Expanded(background),
	layout.Stacked(widget)
)
```

```go
func layoutBackground(gtx layout.Context) layout.Dimensions {
	return layout.Background{}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, background)
			return layout.Dimensions{Size: gtx.Constraints.Min}
		}, func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(30, 100), blue)
		})
}
```

## List（列表）

[layout.List](https://gioui.org/layout#List) 可以显示一个可能很大的项目列表。由于 `List` 还处理滚动，它必须在布局之间持久化，否则滚动位置会丢失。List 通过只布局可见元素来处理大量项目。每一帧，提供的闭包只会对当前滚动位置可见的索引（以及可能在滚动位置上下的少量项目）被调用。

```go
var list = layout.List{}

func listing(gtx layout.Context) layout.Dimensions {
	return list.Layout(gtx, 100, func(gtx layout.Context, i int) layout.Dimensions {
		col := color.NRGBA{R: byte(i * 20), G: 0x20, B: 0x20, A: 0xFF}
		return ColorBox(gtx, image.Pt(20, 100), col)
	})
}
```

## Flex（弹性布局）

[layout.Flex](https://gioui.org/layout#Flex) 根据子元素的权重或刚性约束来布局它们。首先使用刚性元素来确定剩余空间，然后根据权重在弹性子元素之间分配剩余空间。

子元素可以是：

* [Rigid](https://gioui.org/layout#Rigid) —— 用从其他刚性子元素留下的尽可能多的空间进行布局。
* [Flexed](https://gioui.org/layout#Flexed) —— 子元素根据它们的权重和从刚性子元素留下的空间来调整大小。

```go
func flexed(gtx layout.Context) layout.Dimensions {
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, gtx.Constraints.Min, blue)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, gtx.Constraints.Min, green)
		}),
	)
}
```

## Spacer（间隔）

[layout.Spacer](https://gioui.org/layout#Spacer) 可以与 `layout.List` 或 `layout.Flex` 一起使用，在项目之间添加空白空间。

```go
func spacer(gtx layout.Context) layout.Dimensions {
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, gtx.Constraints.Min, blue)
		}),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, gtx.Constraints.Min, green)
		}),
	)
}
```

## Custom（自定义）

有时内置布局不够用。要为控件创建自定义布局，有特殊的函数和结构来操作 layout.Context。一般来说，布局代码对每个子控件执行以下步骤：

* 使用 op.Save。
* 设置 layout.Context.Constraints。
* 设置 op.TransformOp。
* 调用 widget.Layout(gtx, ...)。
* 使用控件返回的尺寸。
* 使用 StateOp.Load。

对于复杂的布局，你还需要使用宏。作为一个例子，请看 [layout.Flex](https://gioui.org/layout#Flex)。它大致实现了：

1. 把控件录制进宏。
2. 为非刚性控件计算尺寸。
3. 通过回放它们的宏，根据计算出的尺寸绘制控件。

---

### 深度解读

**1. 布局的核心思想：约束驱动（Constraint-based）**

Gio 的布局是**父给子约束，子返回实际尺寸**。

- 父布局告诉子：“你最多只能用这么大（Max），至少要这么大（Min）”
- 子布局根据这些约束计算自己真正需要的尺寸，并返回 `layout.Dimensions`
- 父再用这个返回值来决定下一个兄弟放在哪里，或者自己最终占多大

这与 Flutter、SwiftUI 的思路非常接近，但实现更轻量、更底层。

**2. 五大常用布局容器详解**

| 布局           | 主要用途          | 关键特点                           | 是否需要持久化 |
| -------------- | ----------------- | ---------------------------------- | -------------- |
| **Inset**      | 添加内外边距      | 最简单，只改 Constraints           | 否             |
| **Stack**      | 重叠放置          | Expanded 会撑满最大 Stacked 的尺寸 | 否             |
| **Background** | 专门画背景        | Stack 的优化版本，性能更好         | 否             |
| **List**       | 长列表 + 滚动     | **只布局可见项**，支持大量数据     | **必须持久化** |
| **Flex**       | 横向/纵向弹性分配 | 先 Rigid 再按权重分配剩余空间      | 否             |
| **Spacer**     | 固定空白          | 常与 Flex/List 配合                | 否             |

**3. List 是最特殊也最重要的一个**

```go
var list = layout.List{}   // 必须放在结构体或包级变量，跨帧存活
```

为什么必须持久化？
- List 内部保存了当前滚动位置、可见范围等信息
- 如果每帧都 `layout.List{}`，滚动条会一直跳回顶部

List 的性能关键：它**不会**为 10000 个项目都调用你的闭包，只会为当前可见的（+ 缓冲）调用。这让它能轻松处理超长列表。

**4. Flex 的布局算法（简化版）**

1. 先让所有 `Rigid` 子元素按自己想要的尺寸布局
2. 计算剩余空间
3. 按 `Flexed` 的权重比例分配剩余空间
4. 再次布局那些 Flexed 子元素（这次给它们确切的约束）

注意：Flex 默认是横向（Axis: Horizontal）。纵向需要显式设置 `Axis: layout.Vertical`。

**5. 自定义布局的真实难度**

官方给出的步骤看起来简单，但实际写起来并不轻松，因为：

- 你经常需要**先知道子控件想要多大**，才能决定怎么放
- 但要知道大小，就得先调用它的 Layout
- 调用 Layout 又会立刻产生绘制操作

解决办法就是使用 **Macro（录制）**：

1. 先用 `op.Record` 把子控件的绘制操作录下来（不真正画）
2. 拿到它返回的 Dimensions
3. 计算位置
4. 再用变换 + 回放宏的方式真正画出来

这就是 `layout.Flex`、`layout.List` 内部在做的事情。这也是为什么自定义复杂布局时经常会看到 Macro。

**6. 实际开发中的最佳实践**

| 场景          | 推荐                                           |
| ------------- | ---------------------------------------------- |
| 简单边距      | `layout.UniformInset` 或 `layout.Inset`        |
| 背景 + 内容   | 优先用 `layout.Background`，而不是自己写 Stack |
| 横向/纵向排列 | `layout.Flex`                                  |
| 大量数据滚动  | `layout.List`（必须持久化）                    |
| 重叠元素      | `layout.Stack`                                 |
| 固定间距      | `layout.Spacer`                                |
| 完全自定义    | 先研究 Flex 和 List 的源码，再动手             |

**7. 与前后章节的关系**

- 第4章教会我们如何写单个 Widget（接收 Constraints，返回 Dimensions）
- **第5章教会我们如何把多个 Widget 组合起来**
- 第6章 Theme 会用这些布局来实现完整的 Material 控件
- 第7章 Units 解释了为什么 Constraints 是整数像素

布局系统是 Gio 从“能画单个东西”走向“能做完整应用”的关键桥梁。

---

**本章核心记忆点**

> **布局 = 父给约束，子返回实际尺寸**  
> **List 必须跨帧持久化**  
> **Flex 先 Rigid 后按权重分配**  
> **复杂自定义布局几乎都要用 Macro 录制**  
> **Background 是 Stack 的高效特例**

---

准备好后，回复“继续第6章”，我会继续给出 Theme 章节的完整翻译与深度解读。