**原文地址**：https://gioui.org/doc/architecture/drawing

---

采用 **【英文原文】 $\rightarrow$ 【精准逐字翻译】 $\rightarrow$ 【专业术语与 Go 底层原理剖析】** 的方式，为你深度解读 Gio 架构文档的第三章节：**Drawing（绘制）**。

这是 Gio 框架中最核心、最精彩的部分！在 Gio 中，绘制并非像传统的 Canvas 那样直接往屏幕像素缓冲区画线条，而是**将“裁剪（Clipping）”、“画笔设置（Paint）”和“矩阵变换（Transform）”组合成指令推入操作列表（`op.Ops`）**，最终由 GPU 统一渲染。



# 第2章：Drawing（绘制）

### Displaying things on the screen (在屏幕上显示内容)



> **【英文原文】** 
>
> The [`paint`](https://gioui.org/op/paint) package provides operations for drawing graphics.
>
> Coordinates are based on the top-left corner, although it’s possible to [transform the coordinate system](https://gioui.org/op#TransformOp). This means `f32.Point{X:0, Y:0}` is the top left corner of the window. All drawing operations use pixel units, see [Units](https://gioui.org/doc/architecture/drawing#units) section for more information.
>
> For example, the following code will draw a 100x100 pixel colored rectangle at the top left corner of the window:

**【逐字精准翻译】** 

[paint](https://gioui.org/op/paint) 包提供了绘制图形的操作。

坐标系以左上角为原点，不过可以通过 [transform the coordinate system](https://gioui.org/op#TransformOp) 进行变换。这意味着 `f32.Point{X:0, Y:0}` 是窗口的左上角。所有绘制操作都使用像素单位，更多信息请参见 [Units](#units) 章节。

例如，下面的代码会在窗口左上角绘制一个 100x100 像素的彩色矩形：

```go
func drawRedRect(ops *op.Ops) {
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
```



> **【英文原文】** 
>
> The `defer`ed line is only deferring the `.Pop()` operation at the end, so we push a rectangular clipping area, set the color to red with `paint.ColorOp`, and then instruct Gio to paint the current area with the current color with `paint.PaintOp`.

**【逐字精准翻译】** 

延迟执行（deferred）的那一行代码仅仅是将末尾的 `.Pop()` 操作推迟执行。我们先压入（push）一个矩形裁剪区域，用 `paint.ColorOp` 把颜色设为红色，然后用 `paint.PaintOp` 指示 Gio 用当前颜色填充当前（裁剪）区域。

- **核心原理（Gio 绘制三步法）：**
  1. **定义形状 (Clip)：** 用 `clip.Rect` 划定绘制范围（画框）。
  2. **设定画刷 (ColorOp)：** 指定画笔的颜色或纹理（画刷）。
  3. **落地绘制 (PaintOp)：** 触发实际绘制命令，将画刷颜色填满裁剪框。

- **专业术语与 Go 技巧剖析：**
  - `defer clip.Rect{...}.Push(ops).Pop()`：**极其巧妙的 Go 语言惯用法！**
    - `Push(ops)` 会在**函数开始时立即执行**，将裁剪区域压入操作栈并返回一个包含 `.Pop()` 方法的 StackToken 结构体；
    - `defer` 捕获这个结构体，并在**函数返回（Exit）时自动调用 `.Pop()`**，从而恢复原本的裁剪状态，优雅避免了内存泄漏和状态污染。
  - `paint.PaintOp{}`：真正触发“着色/填色”动作的操作指令。如果没有 `PaintOp`，之前的 `ColorOp` 就仅仅是准备好了画笔颜色而已。



## Offset（偏移）

> **【英文原文】** 
>
> Operation [`op.TransformOp`](https://gioui.org/op#TransformOp) translates the position of the operations that come after it.
>
> For example, the following function offsets the red rectangle 100 pixels to the right:
>
> Again, note that we are `defer`ing the `.Pop()` of the offset. This means that the offset is applied for the duration of the function, then removed.

**【逐字精准翻译】** 

操作 [op.TransformOp](https://gioui.org/op#TransformOp) 会平移（translate）排在它后面的操作的位置。

例如，下面的函数把红色矩形向右偏移 100 像素：

```go
func drawRedRect10PixelsRight(ops *op.Ops) {
	defer op.Offset(image.Pt(100, 0)).Push(ops).Pop()
	drawRedRect(ops)
}
```

再次注意，我们延迟执行（defer）的是偏移操作的 `.Pop()`。这意味着偏移只在函数执行期间生效，随后（函数退出时）被移除。



## Clipping（裁剪）

> **【英文原文】** 
>
> In some cases we want the drawing to be ==confined to== a non-rectangular shape, for example to avoid overlapping drawings. Package [`gioui.org/op/clip`](https://gioui.org/op/clip) provides exactly that.
>
> [`clip.RRect`](https://gioui.org/op/clip#RRect) clips all subsequent drawing operations to a rectangle with rounded corners. This is useful as a basis for a button background:

**【逐字精准翻译】** 

在某些情况下，我们希望绘制被限制在非矩形形状内，例如为了避免绘制重叠。包 [gioui.org/op/clip](https://gioui.org/op/clip) 正好提供了这个功能。

[clip.RRect](https://gioui.org/op/clip#RRect) 会把后续所有绘制操作裁剪到一个带圆角的矩形。这很适合作为按钮背景的基础：

```go
func redButtonBackground(ops *op.Ops) {
	const r = 10 // 圆角半径
	bounds := image.Rect(0, 0, 100, 100)
	clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Push(ops) // SE=右下, SW=左下, NW=左上, NE=右上
	drawRedRect(ops)
}
```

- **词汇剖析：**
  - `confined to ...`：限制在 / 局限于……。
  - `non-rectangular shape`：非矩形形状。



> **【英文原文】** 
>
> For more complex clipping [`clip.Path`](https://gioui.org/op/clip#Path) can express shapes built from lines and bézier curves. This example draws a triangle with a curved edge:

**【逐字精准翻译】** 

对于更复杂的裁剪，[clip.Path](https://gioui.org/op/clip#Path) 可以用直线和贝塞尔曲线构建形状。下面的例子绘制了一个带曲线边的三角形：

```go
func redTriangle(ops *op.Ops) {
	var path clip.Path
	path.Begin(ops)
	path.Move(f32.Pt(50, 0))                   // 移动起始点
	path.Quad(f32.Pt(0, 90), f32.Pt(50, 100))  // 二次贝塞尔曲线
	path.Line(f32.Pt(-100, 0))                 // 绘制直线
	path.Line(f32.Pt(50, -100))                // 绘制直线
	defer clip.Outline{Path: path.End()}.Op().Push(ops).Pop()
	drawRedRect(ops)
}
```

- **专业术语剖析：**

  - `bézier curves`：**贝塞尔曲线**。`Quad` 表示 Quadratic Bézier Curve（二次贝塞尔曲线），需要一个控制点和一个终点。

  - `clip.Outline`：将封闭路径转化为绘制裁剪边界。



## Lines（线条）

> **【英文原文】** 
>
> To draw lines we can use [`clip.Stroke`](https://gioui.org/op/clip#Stroke) instead of [`clip.Outline`](https://gioui.org/op/clip#Outline). Stroke draws a fixed-width line along a path, whereas Outline simply does not allow drawing outside of the described path area. We can also use the [`paint.FillShape`](https://gioui.org/op/paint#FillShape) helper to avoid managing the clip state or use `ColorOp` or `PaintOp`. `paint.FillShape` lets us specify an `*op.Ops`, a `color.NRGBA`, and a `clip.AreaOp`, and it takes care of filling the clipped area with the color.

**【逐字精准翻译】** 

要绘制线条，我们可以使用 [clip.Stroke](https://gioui.org/op/clip#Stroke) 而不是 [clip.Outline](https://gioui.org/op/clip#Outline)。Stroke 沿着路径绘制固定宽度的线，而 Outline 只是不允许在描述的路径区域外绘制。

我们也可以使用 [paint.FillShape](https://gioui.org/op/paint#FillShape) 辅助函数，避免自己管理裁剪状态，也不用手动使用 `ColorOp` 或 `PaintOp`。`paint.FillShape` 让我们指定一个 `*op.Ops`、一个 `color.NRGBA` 和一个 `clip.AreaOp`，它会自动用颜色填充被裁剪的区域。

- **词汇剖析：**
  - `clip.Stroke`：描边（绘制有宽度的线条轨迹）。
  - `clip.Outline`：轮廓封闭区域（用于整体填充）。
  - `paint.FillShape`：高阶便捷函数（将 Push、ColorOp、PaintOp、Pop 一站式打包）。



> **【英文原文】** 
>
> It’s possible to use the predefined shapes, such as [`clip.RRect`](https://gioui.org/op/clip#RRect):

**【逐字精准翻译】** 

可以使用预定义的形状，例如 [clip.RRect](https://gioui.org/op/clip#RRect)：

```go
func strokeRect(ops *op.Ops) {
	const r = 10
	bounds := image.Rect(20, 20, 80, 80)
	rrect := clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}
	paint.FillShape(ops, red,
		clip.Stroke{
			Path:  rrect.Path(ops), // 转换为路径
			Width: 4,               // 线条宽度为 4 像素
		}.Op(),
	)
}
```

> **【英文原文】** 
>
> Or use a custom shape drawn with [`clip.Path`](https://gioui.org/op/clip#Path):

**【逐字精准翻译】** 

或者使用 [clip.Path](https://gioui.org/op/clip#Path) 绘制自定义形状：

```go
func strokeTriangle(ops *op.Ops) {
	var path clip.Path
	path.Begin(ops)
	path.MoveTo(f32.Pt(30, 30))
	path.LineTo(f32.Pt(70, 30))
	path.LineTo(f32.Pt(50, 70))
	path.Close()

	paint.FillShape(ops, green,
		clip.Stroke{
			Path:  path.End(),
			Width: 4,
		}.Op())
}
```

> **【英文原文】** 
>
> For dashes, stroke end caps and joins, there’s a separate package [gioui.org/x/stroke](https://gioui.org/x/stroke). However, they are not as performant as `clip.Stroke`, as the work to construct the path description must be performed on the CPU.

**【逐字精准翻译】** 

对于虚线、描边端点样式（caps）和连接样式（joins），有一个单独的包 [gioui.org/x/stroke](https://gioui.org/x/stroke)。然而，它们的性能不如 `clip.Stroke`，因为构建路径描述的工作必须在 CPU 上完成。



## Operation Stack（操作栈）

> **【英文原文】** 
>
> Some operations affect all operations that follow them. For example, [`paint.ColorOp`](https://gioui.org/op/paint#ColorOp) sets the “brush” color that is used in subsequent [`paint.PaintOp`](https://gioui.org/op/paint#PaintOp) operations. This drawing context also includes coordinate transformation (set by [`op.TransformOp`](https://gioui.org/op#TransformOp)) and clipping (set by [`clip.Op`](https://gioui.org/op/clip#Op)).
>
> Some operations, such as clips and transformations, allow you to temporarily apply them and later restore the previous state.
>
> For example, the `redButtonBackground` function in the previous section has the unfortunate side-effect of clipping all later operations to the outline of the button background! Let’s make a version of it that doesn’t affect any callers:

**【逐字精准翻译】** 

有些操作会影响后续所有操作。例如，[paint.ColorOp](https://gioui.org/op/paint#ColorOp) 设置了后续 [paint.PaintOp](https://gioui.org/op/paint#PaintOp) 使用的“画刷”颜色。这个绘制上下文（drawing context）还包括坐标变换（由 [op.TransformOp](https://gioui.org/op#TransformOp) 设置）和裁剪（由 [clip.Op](https://gioui.org/op/clip#Op) 设置）。

有些操作（如裁剪和变换）允许你临时应用它们，之后再恢复之前的状态。

例如，上一节的 `redButtonBackground` 函数有一个不幸的副作用：它会把后续所有操作都裁剪到按钮背景的轮廓！让我们做一个不会影响调用者的版本：

```go
func redButtonBackgroundStack(ops *op.Ops) {
	const r = 1 // 圆角半径
	bounds := image.Rect(0, 0, 100, 100)
    // 使用 defer .Pop() 保证出栈恢复状态
	defer clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Push(ops).Pop()
	drawRedRect(ops)
}
```

- **概念透视：** 在上文不带 `defer ... Pop()` 的 `redButtonBackground` 中，`Push(ops)` 把裁剪状态压栈后一直生效，导致后续其他控件都被裁剪了；加入了 `defer ... Pop()` 后，出函数时弹出状态，对外界“零副作用”。



## Drawing Order（绘制顺序）

> **【英文原文】** 
>
> Drawing happens from back to front. Things inserted into the `op.Ops` first are drawn first, and later elements will be drawn on top. In this function the green rectangle is drawn on top of red rectangle:

**【逐字精准翻译】** 

绘制是从后到前（由底层到顶层）进行的。先插入 `op.Ops` 的内容先画，后面的元素会画在顶部（覆盖在上面）。在这个函数中，绿色矩形画在红色矩形之上：

```go
func drawOverlappingRectangles(ops *op.Ops) {
	// 绘制一个红色矩形（先画，在底层）
	cl := clip.Rect{Max: image.Pt(100, 50)}.Push(ops)
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
	cl.Pop()

	// 绘制一个绿色矩形（后画，覆盖在上面）
	cl = clip.Rect{Max: image.Pt(50, 100)}.Push(ops)
	paint.ColorOp{Color: color.NRGBA{G: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
	cl.Pop()
}
```



> **【英文原文】** 
>
> Sometimes you may want to change this order. For example, you may want to delay drawing to apply a transform that is calculated during drawing, or you may want to perform a list of operations several times. For this purpose there is [op.MacroOp](https://gioui.org/op#MacroOp).

**【逐字精准翻译】** 

有时你可能想改变这个顺序。例如，你可能想延迟绘制以应用一个在绘制过程中才计算出来的变换，或者你可能想多次执行一组操作。为此有 [op.MacroOp](https://gioui.org/op#MacroOp)（宏操作）。

```go
func drawFiveRectangles(ops *op.Ops) {
	// 把 drawRedRect 的操作录制进宏。
	macro := op.Record(ops)
	drawRedRect(ops)
	c := macro.Stop() // 停止录制，获取 CallOp 句柄

	// “回放”宏 5 次，每次垂直平移 20px、水平平移 110 像素。
	for i := 0; i < 5; i++ {
		c.Add(ops) // 播放指令
		op.Offset(image.Pt(110, 20)).Add(ops)
	}
}
```

- **专业术语剖析：**
  - `op.Record(ops)`：**指令录制**。它开启一个“宏”，暂时拦截写入 `ops` 的指令，允许你把一段绘制逻辑当成一个“可复用函数对象”重放多次。



## Animation（动画）

> **【英文原文】** 
>
> Gio only issues FrameEvents when the window is resized or the user interacts with the window. However, animation requires continuous redrawing until the animation is completed. For that there is [`op.InvalidateCmd`](https://gioui.org/op#InvalidateCmd).
>
> The following code will animate a red “progress bar” that fills up from left to right over 10 seconds from when the program starts:

**【逐字精准翻译】** 

Gio 只在窗口大小改变或用户与窗口交互时才发出 FrameEvent。然而，动画需要持续进行重新绘制，直到动画完成。为此有 [op.InvalidateCmd](https://gioui.org/op#InvalidateCmd)。

下面的代码会从程序启动开始，在 10 秒内从左到右填充一个红色“进度条”：

```go
var startTime = time.Now()
var duration = 10 * time.Second

func drawProgressBar(ops *op.Ops, source input.Source, now time.Time) {
	// 根据当前时间，计算需要绘制多少进度条
	elapsed := now.Sub(startTime)
	progress := elapsed.Seconds() / duration.Seconds()
	if progress < 1 {
		// 进度条动画尚未结束，向系统发出“请求下一帧重绘”指令！
		source.Execute(op.InvalidateCmd{})
	} else {
		progress = 1
	}

	width := 200 * float32(progress)
	defer clip.Rect{Max: image.Pt(int(width), 20)}.Push(ops).Pop()
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
```

- **架构核心机制剖析：**

  - **按需渲染 (On-demand Rendering)：** 节能是 Gio 的重大优点。为了省电，默认无事件时不绘制。

  - `source.Execute(op.InvalidateCmd{})`：向 Gio 表达“这一帧还没完，请在下一帧刷屏（VSync）时立刻再传一个 `FrameEvent` 给我”，从而驱动连续 60 FPS/120 FPS 的平滑动画！



## Record and replay（录制与回放）

> **【英文原文】** 
>
> While `op.MacroOp` allows you to record and replay operations on a single operation list, [`op.CallOp`](https://gioui.org/op#CallOp) allows for reuse of a separate operation list. This is useful for caching operations that are expensive to re-create, or for animating the disappearance of otherwise removed widgets:

**【逐字精准翻译】** 

虽然 `op.MacroOp` 允许你在同一个操作列表上录制和回放操作，[op.CallOp](https://gioui.org/op#CallOp) 则允许复用一个独立的操作列表。这对于缓存重新创建成本高昂的操作，或者为本已被删除的组件制作消失动画非常有用：

```go
func drawWithCache(ops *op.Ops) {
	// 把操作保存在一个独立的 ops 值中（缓存）。
	cache := new(op.Ops)
	macro := op.Record(cache)

	cl := clip.Rect{Max: image.Pt(100, 100)}.Push(cache)
	paint.ColorOp{Color: color.NRGBA{G: 0x80, A: 0xFF}}.Add(cache)
	paint.PaintOp{}.Add(cache)
	cl.Pop()
	call := macro.Stop()

	// 从缓存中绘制这些操作。
	call.Add(ops)
}
```



> **【英文原文】** 
>
> Note: For this cache to actually save any work across frames, you’ll need to allocate the cache’s `op.Ops` somewhere that persists across frames. Doing it in a local variable like this will mean that the cache is recreated every frame.

**【逐字精准翻译】** 

注意：要让此缓存在跨帧时真正节省工作量，你需要将缓存的 `op.Ops` 分配在能够在帧与帧之间持久留存的位置（例如长久存活的结构体字段中）。像这样在局部变量中创建它将意味着每一帧都会重新创建缓存。



## Images（图片）

> **【英文原文】** 
>
> [`paint.ImageOp`](https://gioui.org/op/paint#ImageOp) is used to draw images. Like [`paint.ColorOp`](https://gioui.org/op/paint#ColorOp), it sets part of the drawing context (the “brush”) that’s used for subsequent [`PaintOp`](https://gioui.org/op/paint#PaintOp). [`ImageOp`](https://gioui.org/op/paint#ImageOp) is used similarly to [`ColorOp`](https://gioui.org/op/paint#ColorOp).
>
> Note that [`image.NRGBA`](https://golang.org/pkg/image#NRGBA) and [`image.Uniform`](https://golang.org/pkg/image#Uniform) images are efficient and treated specially. Other [`Image`](https://golang.org/pkg/image#Image) implementations will undergo a more expensive copy and conversion to the underlying image model.

**【逐字精准翻译】** 

[paint.ImageOp](https://gioui.org/op/paint#ImageOp) 用于绘制图片。和 [paint.ColorOp](https://gioui.org/op/paint#ColorOp) 一样，它设置了绘制上下文的一部分（“画刷”），供后续的 [PaintOp](https://gioui.org/op/paint#PaintOp) 使用。[ImageOp](https://gioui.org/op/paint#ImageOp) 的使用方式与 [ColorOp](https://gioui.org/op/paint#ColorOp) 类似。

注意，[image.NRGBA](https://golang.org/pkg/image#NRGBA) 和 [image.Uniform](https://golang.org/pkg/image#Uniform) 图片是高效的，并得到了特殊处理。其他的 [Image](https://golang.org/pkg/image#Image) 实现将经历成本更高的复制和转换过程，以转为底层图像模型。

```go
func drawImage(ops *op.Ops, img image.Image) {
	imageOp := paint.NewImageOp(img)
	imageOp.Filter = paint.FilterNearest // 最近邻过滤（适合像素风或清晰放大）
	imageOp.Add(ops)
	op.Affine(f32.Affine2D{}.Scale(f32.Pt(0, 0), f32.Pt(4, 4))).Add(ops) // 放大 4 倍
	paint.PaintOp{}.Add(ops)
}
```



> **【英文原文】** 
>
> The image must not be mutated until another [`FrameEvent`](https://gioui.org/io/app#FrameEvent) happens, because the image may be read asynchronously while the frame is being drawn. Additionally, mutations to the image provided to `paint.ImageOp` are not guaranteed to ever be reflected in the drawn content. To update an image on-screen, create a new image.Image and construct a new `paint.ImageOp`.

**【逐字精准翻译】** 

在下一个 [FrameEvent](https://gioui.org/io/app#FrameEvent) 发生之前，切勿修改（mutate）该图像，因为在绘制该帧时，图像可能会被异步读取。此外，对提供给 `paint.ImageOp` 的图片的修改，并不保证会反映到绘制的内容中。要在屏幕上更新一张图片，请创建新的 image.Image 并构造新的 `paint.ImageOp`。

- **性能注意：** 传入 `image.NRGBA` 格式性能最好，GPU 能直接上传像素点。

至此，**Drawing（绘制）** 章节全部解读完成！

随时可以继续推进下一个关键章节：**Input（输入事件处理）**！

---

### 深度解读

**1. 绘制的本质：操作栈 + 当前状态**

Gio 的绘制不是“画一个矩形”这种命令式调用，而是：

1. 设置当前裁剪区域（clip）
2. 设置当前画刷（颜色或图片）
3. 执行 `PaintOp` → 用当前画刷填充当前裁剪区域

这非常像 GPU 的状态机模型。所有复杂图形最终都归结为“路径 + 填充/描边”。

**2. Push / Pop 是最重要的习惯**

几乎所有状态改变操作都支持：

```go
defer something.Push(ops).Pop()
```

这保证了即使函数中途 return 或 panic，状态也能正确恢复。**强烈推荐永远使用 defer 形式**，而不是手动 Push + 手动 Pop。

**3. 裁剪（Clipping）既用于绘制，也用于输入**

`clip.Rect`、`clip.RRect`、`clip.Path` 同时定义了：
- 绘制的可见区域
- 输入事件的命中区域（第3章会详细讲）

这是 Gio 设计的精妙之处：同一套区域描述同时服务渲染和事件路由。

**4. 绘制顺序是严格的“后写覆盖先写”**

先加入 `ops` 的操作先画（在底层），后加入的画在上面。这和画家算法一致。

如果需要改变顺序（比如先画前景再画背景，或者多次复用同一组操作），就必须使用 **Macro**（录制）或 **CallOp**（独立缓存）。

**5. 动画必须主动请求重绘**

Gio 默认是“事件驱动”的：没有用户交互或窗口大小变化，就不会发 FrameEvent。  
动画必须在绘制时检测“我还在动”，然后执行：

```go
source.Execute(op.InvalidateCmd{})
```

这会强制下一帧立刻到来。这是即时模式动画的标准写法。

**6. 图片的隐藏坑**

- 图片数据在帧绘制期间可能被异步读取，所以**绝对不能修改**已经交给 `ImageOp` 的图片。
- 想更新图片 → 必须新建一张 `image.Image` + 新建 `ImageOp`。
- 只有 `image.NRGBA` 和 `image.Uniform` 是高效路径，其他格式会触发 CPU 转换。

**7. 性能关键点**

| 操作                  | 性能 | 说明             |
| --------------------- | ---- | ---------------- |
| `clip.Stroke`         | 高   | GPU 路径         |
| `gioui.org/x/stroke`  | 较低 | CPU 构建复杂描边 |
| `op.Macro` / `CallOp` | 高   | 可跨帧缓存       |
| 每帧重新创建复杂 Path | 低   | 应缓存           |

**8. 与其他章节的关系**

- 所有绘制最终都发生在 `FrameEvent` 的 `draw` 函数里（第1章）。
- 布局系统（第5章）会计算尺寸后，再调用这些绘制操作。
- 输入系统（第3章）复用同样的裁剪区域。

---

**本章核心记忆点**

> **绘制 = 设置裁剪区域 + 设置画刷 + PaintOp**  
> **状态隔离 = 永远用 defer Push().Pop()**  
> **动画 = 主动 InvalidateCmd**  
> **复用 = Macro 或 CallOp**

---

准备好后，回复“继续第3章”，我会立即给出 Input 章节的完整翻译与深度解读。