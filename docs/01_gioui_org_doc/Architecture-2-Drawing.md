**第2章：Drawing（绘制）**

**原文地址**：https://gioui.org/doc/architecture/drawing

---

### 完整中文翻译

# Drawing

# Drawing 在屏幕上显示东西

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

`defer` 那一行实际上只是在函数结束时调用 `.Pop()`。我们先压入一个矩形裁剪区域，用 `paint.ColorOp` 把颜色设为红色，然后用 `paint.PaintOp` 指示 Gio 用当前颜色填充当前区域。

## Offset（偏移）

操作 [op.TransformOp](https://gioui.org/op#TransformOp) 会平移它之后所有操作的位置。

例如，下面的函数把红色矩形向右偏移 100 像素：

```go
func drawRedRect10PixelsRight(ops *op.Ops) {
	defer op.Offset(image.Pt(100, 0)).Push(ops).Pop()
	drawRedRect(ops)
}
```

同样，我们 `defer` 了偏移的 `.Pop()`。这意味着偏移只在函数执行期间生效，之后会被移除。

## Clipping（裁剪）

有时我们希望绘制被限制在非矩形形状内，例如避免重叠绘制。包 [gioui.org/op/clip](https://gioui.org/op/clip) 正好提供了这个功能。

[clip.RRect](https://gioui.org/op/clip#RRect) 会把后续所有绘制操作裁剪到一个带圆角的矩形。这很适合作为按钮背景的基础：

```go
func redButtonBackground(ops *op.Ops) {
	const r = 10 // 圆角半径
	bounds := image.Rect(0, 0, 100, 100)
	clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Push(ops)
	drawRedRect(ops)
}
```

对于更复杂的裁剪，[clip.Path](https://gioui.org/op/clip#Path) 可以用直线和贝塞尔曲线构建形状。下面的例子绘制了一个带曲线边的三角形：

```go
func redTriangle(ops *op.Ops) {
	var path clip.Path
	path.Begin(ops)
	path.Move(f32.Pt(50, 0))
	path.Quad(f32.Pt(0, 90), f32.Pt(50, 100))
	path.Line(f32.Pt(-100, 0))
	path.Line(f32.Pt(50, -100))
	defer clip.Outline{Path: path.End()}.Op().Push(ops).Pop()
	drawRedRect(ops)
}
```

## Lines（线条）

要绘制线条，我们可以使用 [clip.Stroke](https://gioui.org/op/clip#Stroke) 而不是 [clip.Outline](https://gioui.org/op/clip#Outline)。Stroke 沿着路径绘制固定宽度的线，而 Outline 只是不允许在描述的路径区域外绘制。

我们也可以使用 [paint.FillShape](https://gioui.org/op/paint#FillShape) 辅助函数，避免自己管理裁剪状态，也不用手动使用 `ColorOp` 或 `PaintOp`。`paint.FillShape` 让我们指定一个 `*op.Ops`、一个 `color.NRGBA` 和一个 `clip.AreaOp`，它会自动用颜色填充被裁剪的区域。

可以使用预定义的形状，例如 [clip.RRect](https://gioui.org/op/clip#RRect)：

```go
func strokeRect(ops *op.Ops) {
	const r = 10
	bounds := image.Rect(20, 20, 80, 80)
	rrect := clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}
	paint.FillShape(ops, red,
		clip.Stroke{
			Path:  rrect.Path(ops),
			Width: 4,
		}.Op(),
	)
}
```

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

对于虚线、描边端点样式和连接样式，有一个单独的包 [gioui.org/x/stroke](https://gioui.org/x/stroke)。不过它们不如 `clip.Stroke` 高效，因为构建路径描述的工作必须在 CPU 上完成。

## Operation Stack（操作栈）

有些操作会影响后续所有操作。例如，[paint.ColorOp](https://gioui.org/op/paint#ColorOp) 设置了后续 [paint.PaintOp](https://gioui.org/op/paint#PaintOp) 使用的“画刷”颜色。这个绘制上下文还包括坐标变换（由 [op.TransformOp](https://gioui.org/op#TransformOp) 设置）和裁剪（由 [clip.Op](https://gioui.org/op/clip#Op) 设置）。

有些操作（如裁剪和变换）允许你临时应用它们，之后再恢复之前的状态。

例如，上一节的 `redButtonBackground` 函数有一个不幸的副作用：它会把后续所有操作都裁剪到按钮背景的轮廓！让我们做一个不会影响调用者的版本：

```go
func redButtonBackgroundStack(ops *op.Ops) {
	const r = 1 // 圆角半径
	bounds := image.Rect(0, 0, 100, 100)
	defer clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Push(ops).Pop()
	drawRedRect(ops)
}
```

## Drawing Order（绘制顺序）

绘制是从后往前进行的。先插入 `op.Ops` 的内容先画，后面的元素会画在上面。在这个函数中，绿色矩形画在红色矩形之上：

```go
func drawOverlappingRectangles(ops *op.Ops) {
	// 画红色矩形。
	cl := clip.Rect{Max: image.Pt(100, 50)}.Push(ops)
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
	cl.Pop()

	// 画绿色矩形。
	cl = clip.Rect{Max: image.Pt(50, 100)}.Push(ops)
	paint.ColorOp{Color: color.NRGBA{G: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
	cl.Pop()
}
```

有时你可能想改变这个顺序。例如，你可能想延迟绘制以应用一个在绘制过程中才计算出来的变换，或者你可能想多次执行一组操作。为此有 [op.MacroOp](https://gioui.org/op#MacroOp)。

```go
func drawFiveRectangles(ops *op.Ops) {
	// 把 drawRedRect 的操作录制进宏。
	macro := op.Record(ops)
	drawRedRect(ops)
	c := macro.Stop()

	// “回放”宏 5 次，每次垂直平移 20px、水平平移 110 像素。
	for i := 0; i < 5; i++ {
		c.Add(ops)
		op.Offset(image.Pt(110, 20)).Add(ops)
	}
}
```

## Animation（动画）

Gio 只在窗口大小改变或用户与窗口交互时才发出 FrameEvent。然而，动画需要持续重绘直到动画完成。为此有 [op.InvalidateCmd](https://gioui.org/op#InvalidateCmd)。

下面的代码会从程序启动开始，在 10 秒内从左到右填充一个红色“进度条”：

```go
var startTime = time.Now()
var duration = 10 * time.Second

func drawProgressBar(ops *op.Ops, source input.Source, now time.Time) {
	// 根据当前时间计算进度条要画多少。
	elapsed := now.Sub(startTime)
	progress := elapsed.Seconds() / duration.Seconds()
	if progress < 1 {
		// 进度条还没动画完。
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

## Record and replay（录制与回放）

虽然 `op.MacroOp` 允许你在同一个操作列表上录制和回放操作，[op.CallOp](https://gioui.org/op#CallOp) 则允许复用一个独立的操作列表。这对于缓存那些重新创建代价较高的操作，或者用于动画那些已经移除的控件消失过程很有用：

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

注意：为了让这个缓存真正在跨帧之间节省工作，你需要把缓存的 `op.Ops` 分配在能跨帧持久存在的地方。像上面这样在局部变量里做，意味着缓存每帧都会被重新创建。

## Images（图片）

[paint.ImageOp](https://gioui.org/op/paint#ImageOp) 用于绘制图片。和 [paint.ColorOp](https://gioui.org/op/paint#ColorOp) 一样，它设置了绘制上下文的一部分（“画刷”），供后续的 [PaintOp](https://gioui.org/op/paint#PaintOp) 使用。[ImageOp](https://gioui.org/op/paint#ImageOp) 的使用方式与 [ColorOp](https://gioui.org/op/paint#ColorOp) 类似。

注意，[image.NRGBA](https://golang.org/pkg/image#NRGBA) 和 [image.Uniform](https://golang.org/pkg/image#Uniform) 图片是高效的，会被特殊处理。其他 [Image](https://golang.org/pkg/image#Image) 实现会经历更昂贵的复制和转换到底层图像模型。

```go
func drawImage(ops *op.Ops, img image.Image) {
	imageOp := paint.NewImageOp(img)
	imageOp.Filter = paint.FilterNearest
	imageOp.Add(ops)
	op.Affine(f32.Affine2D{}.Scale(f32.Pt(0, 0), f32.Pt(4, 4))).Add(ops)
	paint.PaintOp{}.Add(ops)
}
```

图片在另一个 [FrameEvent](https://gioui.org/io/app#FrameEvent) 发生之前不能被修改，因为图片可能在帧绘制过程中被异步读取。此外，对提供给 `paint.ImageOp` 的图片的修改，并不保证会反映到绘制的内容中。要在屏幕上更新一张图片，请创建新的 image.Image 并构造新的 `paint.ImageOp`。

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