**原文地址**：https://gioui.org/doc/architecture/layout

---

采用 **【英文原文】 $\rightarrow$ 【精准逐字翻译】** $\rightarrow$ **【专业术语与 Gio 布局机制深度剖析】** 的方式，为你拆解 Gio 官方文档的第六章节：**Layout（布局系统）**。

Gio 的 `layout` 包是声明式 UI 的精髓所在。它摒弃了传统坐标系绝对定位，全面转向基于约束传播（Constraint Propagation）**与**闭包组合（Closure Composition）的动态弹性布局系统。



# 第5章：Layout（布局）

#### Putting things where they belong (把事物放在属于它们的位置)



> **【英文原文】** 
>
> Package [`gioui.org/layout`](https://gioui.org/layout) provides support for common layout operations such as spacing, lists and stacks of overlapping widgets.
>
> In the layout examples we’ll use this `ColorBox` widget to visualize layouts:

**【逐字精准翻译】** 

包 [gioui.org/layout](https://gioui.org/layout) 为常见的布局操作提供支持，例如间距、列表和重叠控件的堆叠。

在布局示例中，我们将使用这个 `ColorBox` 组件来可视化布局：

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

**概念剖析：**

- `ColorBox` 是标准的无状态组件范例：裁切区域 $\rightarrow$ 填充颜色 $\rightarrow$ 绘制 $\rightarrow$ 返回 `layout.Dimensions`。



## Inset（内边距）

[`layout.Inset`](https://gioui.org/layout#Inset) adds space around a widget.
`layout.Inset` 在组件周围添加空间（外边距/内边距）。

```go
func inset(gtx layout.Context) layout.Dimensions {
	// 在彼此内部绘制矩形，带有 30dp 的内边距。
	return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return ColorBox(gtx, gtx.Constraints.Max, red)
	})
}
```

- **核心原理：** `Inset.Layout` 会在调用闭包子组件前，先调整并**缩小** `gtx.Constraints`（减少 30dp），并在 `gtx.Ops` 中压入一个 `op.Offset` 平移坐标系，然后再把修改后的 `gtx` 传递给子组件。



## Stack（层叠布局）

> **【英文原文】** 
>
> [`layout.Stack`](https://gioui.org/layout#Stack) lays out overlapping child elements according to the alignment direction. The child of a stack layout can be:
>
> - [`Stacked`](https://gioui.org/layout#Stacked) - which doesn’t have minimum constraints and the maximum constraints passed to Stack.Layout.
> - [`Expanded`](https://gioui.org/layout#Expanded) - which uses the largest Stacked item as the minimum constraint and maximum is the maximum constraints passed to Stack.Layout.
>
> For example, this draws green and blue rectangles on top of a red background:

**【逐字精准翻译】** 

`layout.Stack` 根据对齐方向对重叠的子元素进行布局。层叠布局的子元素可以是：

* **Stacked（自适应层）**：没有最小尺寸约束，其最大尺寸约束为传递给 `Stack.Layout` 的最大约束。(*==自然层叠/松散约束==*)
* **Expanded（拉伸层）**：使用所有 `Stacked` 项中尺寸最大的那个作为其最小尺寸约束，其最大尺寸约束为传递给 `Stack.Layout` 的最大约束。(*==展开填充==*)

例如，这段代码在红色背景上绘制了绿色和蓝色的矩形：

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

- **算法细节：** `Stack` 会先计算所有 `Stacked` 类型的子组件，得出一个最大包围盒（100x100）。然后将这个包围盒大小作为最小约束（`Constraints.Min`）强制传给 `Expanded` 子组件，从而实现“背景自动撑满前景”的层叠效果。
- **计算顺序机制（Two-Pass Strategy）：** 
  `Stack` 会**先评估所有 `Stacked` 子项**，测量出最大的宽高（例如上例中绿宽 100，蓝高 100 $\rightarrow$ 算出 $100 \times 100$）；然后将这个最大值作为 `Min` 约束传递给 **`Expanded` 子项**，实现背景自动铺满前景的效果。



### Background（背景）

> **【英文原文】** 
>
> Because layouting a background for a widget is very frequent there is a more performant implementation for that scenario, which roughly corresponds to:

**【逐字精准翻译】** 

由于给组件布局背景非常频繁，Gio 为该场景提供了一种性能更高的实现，其逻辑大致等价于：

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

> **【英文原文】** 
>
> [`layout.List`](https://gioui.org/layout#List) can display a potentially large list of items. Since `List` also handles scrolling it must be persisted across layouts, otherwise the scrolling position is lost. List handles large numbers of items by only laying out the visible elements. Each frame, the provided closure is invoked only for indicies visible at the current scroll position (and possibly a small number of items above and below the scroll position).

**【逐字精准翻译】** 

`layout.List` 可以高效显示海量数据列表。由于 `List` 同时也负责处理滚动，因此它的状态**必须在多次布局帧之间持久化保存**（声明为全局变量或 Struct 字段），否则滚动位置将会丢失。`List` 处理海量数据的方式是**仅对可见元素进行布局**。在每一帧中，传入的闭包仅针对当前滚动位置可见的索引（以及滚动位置上下相邻的极少数元素）被调用。

```go
// 必须在函数体外定义，保持滚动状态持久
var list = layout.List{}

func listing(gtx layout.Context) layout.Dimensions {
    // 参数分别表示：gtx，元素总数，子项生成闭包
	return list.Layout(gtx, 100, func(gtx layout.Context, i int) layout.Dimensions {
		col := color.NRGBA{R: byte(i * 20), G: 0x20, B: 0x20, A: 0xFF}
		return ColorBox(gtx, image.Pt(20, 100), col)
	})
}
```

- **性能关键机制（Virtualization）：** 
  类似于 Flutter 的 `ListView.builder` 或 Web 的 Virtual List。就算元素总数传 `1_000_000`，Gio 也只会在渲染帧时调用当前视口能容纳的 10~20 次闭包，内存与 GPU 开销为 $O(\text{visible})$。
- **核心算法（虚拟化渲染 / Virtual Windowing）**：Gio 的 `List` 是内置懒加载/虚拟列表的。即便列表有 1,000,000 个元素，闭包也只会触发渲染当前屏幕可视区域内的那 10~20 个，因此性能极高且内存零占用。



## Flex（弹性布局）

> **【英文原文】** 
>
> [`layout.Flex`](https://gioui.org/layout#List) lays out children according to their weights or rigid constraints. First the rigid elements are used to determine the remaining space and then the remaining space is divided among flexed children according to weights.
>
> The children can be:
>
> - [`Rigid`](https://gioui.org/layout#Rigid) - are laid out with as much space left over from other rigid children.
> - [`Flexed`](https://gioui.org/layout#Flexed) - children are sized according to their weights and the space left over from rigid children.

**【逐字精准翻译】** 

`layout.Flex` 根据权重（Weights）或固定/刚性约束（Rigid constraints）来排列子元素。它首先测量所有刚性（Rigid）元素以确定剩余可用空间，然后根据权重将剩余空间分配给弹性（Flexed）子元素。

子元素可以是：

* **Rigid（刚性/固定）：** 根据其他刚性元素剩余的空间来进行布局（尺寸由自身决定，不受比例分配影响）。
* **Flexed（弹性/按比例）：** 尺寸根据其权重以及从刚性元素中剩下的空间进行计算分摊。

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

- **算法逻辑（类似 CSS Flexbox）：**
  1. 计算所有 `Rigid` 子项的总宽度（如 $100 + 100 = 200\text{px}$）。
  2. 计算剩余可用空间 $S_{\text{remain}} = \text{Max.X} - 200$。
  3. 分别将 $0.5 \times S_{\text{remain}}$ 作为精确约束传给两个 `Flexed` 子项。
- **对照概念：** 类似于 CSS 的 Flexbox（`Axis: layout.Horizontal/Vertical`），`Rigid` 相当于 `flex: 0 0 auto`，`Flexed` 相当于 `flex: weight`。



## Spacer（间隔）

> **【英文原文】** 
>
> [`layout.Spacer`](https://gioui.org/layout#Spacer) can be used together with `layout.List` or `layout.Flex` to add empty space between items.

**【逐字精准翻译】** 

`layout.Spacer` 可以与 `layout.List` 或 `layout.Flex` 配合使用，在子项之间添加空白间距。

```go
func spacer(gtx layout.Context) layout.Dimensions {
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Rigid(layout.Spacer{Width: 20}.Layout), // 插入 20px 间距
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

> **【英文原文】** 
>
> Sometimes the builtin layouts are not sufficient. To create a custom layout for widgets there are special functions and structures to manipulate layout.Context. In general, layout code performs the following steps for each sub-widget:
>
> - Use `op.Save`.
> - Set `layout.Context.Constraints`.
> - Set `op.TransformOp`.
> - Call `widget.Layout(gtx, ...)`.
> - Use dimensions returned by widget.
> - Use `StateOp.Load`.

**【逐字精准翻译】** 

有时内置的布局组件并不够用。为了给组件创建自定义布局，存在专门用于操作 `layout.Context` 的函数和结构体。通常，布局代码对每个子组件执行以下步骤：

* 使用 `op.Save`（保存当前状态）。
* 设置 `layout.Context.Constraints`（约束条件）。
* 设置 `op.TransformOp`（平移/转换坐标系）。
* 调用 `widget.Layout(gtx, ...)` 进行子控件布局。
* 使用组件返回的 `Dimensions`（尺寸）。
* 使用 `StateOp.Load`（恢复之前的状态）。

对于复杂的布局，你还需要使用**宏（Macros）**。以 `layout.Flex` 为例，它的内部逻辑大致实现为：

1. 在宏（macros）中录制组件（不立刻绘制）。
2. 计算非刚性（Flexed）组件的尺寸。
3. 通过回放（Replay）录制的 Macro，根据计算好的尺寸将组件绘制出来。

**宏（Macro）的必要性：** 
在 Go 即时模式 UI 中，如果不使用 Macro，调用 `widget.Layout()` 会**直接将画笔指令追加到全局 Ops 中**。而 Flex 布局在不知道 Rigid 元素尺寸前，无法确定 Flexed 元素的坐标偏移。因此通过 `op.Record()` 先“录屏”但不输出，等坐标算清后再 `replay()`，这就是 Gio 高效自定义布局的底层秘诀！

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