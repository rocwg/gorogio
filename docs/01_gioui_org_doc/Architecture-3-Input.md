**第3章：Input（输入）**

**原文地址**：https://gioui.org/doc/architecture/input

---

### 完整中文翻译

# Input

# Input 对鼠标和键盘做出反应

输入通过 [app.FrameEvent](https://gioui.org/app#FrameEvent) 的 [Queue](https://gioui.org/app#input.Source) 字段传递给控件。

`input.Source` 中最常见的一些事件有：

* [key.Event](https://gioui.org/io/key#Event)、[key.FocusEvent](https://gioui.org/io/key#FocusEvent) —— 键盘输入。
* [key.EditEvent](https://gioui.org/io/key#EditEvent) —— 文本编辑。
* [pointer.Event](https://gioui.org/io/pointer#Event) —— 鼠标和触摸输入。

程序可以按自己喜欢的方式响应这些事件——例如，更新本地数据结构或执行用户触发的动作。[FrameEvent](https://gioui.org/app#FrameEvent) 是特殊的——当程序收到 [FrameEvent](https://gioui.org/app#FrameEvent) 时，它有责任通过调用 [e.Frame](https://gioui.org/app#FrameEvent.Frame) 函数，并传入代表新状态的操作列表，来更新显示。这些操作是在响应 [FrameEvent](https://gioui.org/app#FrameEvent) 时立即生成的，这是 Gio 被称为“即时模式”GUI 的主要原因。

事件处理器（例如 [gioui.org/gesture](https://gioui.org/gesture) 包中的 [Click](https://gioui.org/gesture#Click) 和 [Scroll](https://gioui.org/gesture#Scroll)）可以从单个点击事件中检测出更高层次的动作。

为了在多个不同控件之间分发输入，Gio 需要知道事件处理器及其配置。然而，由于 Gio 框架是无状态的，程序没有直接的方式来指定这些信息。

相反，一些操作会把输入事件类型（例如键盘按键）与程序自己选择的任意 [tags](https://gioui.org/io/event#Tag)（interface{} 值）关联起来。程序在处理 [FrameEvent](https://gioui.org/app#FrameEvent) 时创建这些操作——输入操作和其他操作一样。作为回报，[input.Source](https://gioui.org/io/input#Source) 会提供自上一帧以来到达的事件，并按 tag 分类。

你可以把 tag 想象成某个输入区域的唯一键。Gio 事件路由会把该区域上的输入事件与为该区域提供的 tag 关联起来。然后你可以通过向 `input.Source` 提供相同的 tag，在下一帧获取这些事件。控件通常会封装这个事件逻辑，把指向自己持久状态的指针作为其输入区域的 tag。

下面的例子演示了指针输入处理：

```go
var tag = new(bool) // 我们也可以用 &pressed 代替。
var pressed = false

func doButton(ops *op.Ops, q input.Source) {
	// 把感兴趣的区域限制在一个 100x100 的矩形内。
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()

	// 声明 `tag` 是目标之一。
	event.Op(ops, tag)

	// 处理上一帧到这一帧之间到达的事件。
	for {
		ev, ok := q.Event(pointer.Filter{
			Target: tag,
			Kinds:  pointer.Press | pointer.Release,
		})
		if !ok {
			break
		}

		if x, ok := ev.(pointer.Event); ok {
			switch x.Kind {
			case pointer.Press:
				pressed = true
			case pointer.Release:
				pressed = false
			}
		}
	}

	// 绘制按钮。
	var c color.NRGBA
	if pressed {
		c = color.NRGBA{R: 0xFF, A: 0xFF}
	} else {
		c = color.NRGBA{G: 0xFF, A: 0xFF}
	}
	paint.ColorOp{Color: c}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
```

使用 Go 指针值作为输入 tag 很方便，因为把指针转换成 interface{} 成本很低，而且很容易让这个值专属于某个本地数据结构，从而避免 tag 冲突的风险。

更多细节请查看 [gioui.org/io/pointer](https://gioui.org/io/pointer)（指针/鼠标事件）和 [gioui.org/io/key](https://gioui.org/io/key)（键盘事件）。

## External input（外部输入）

一个完整的帧包括：获取输入、注册输入兴趣、绘制新状态：

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
		// 请求绘制窗口状态。

		// 把操作列表重置回零。
		ops.Reset()
		// 根据 e.Queue 中的事件，把状态绘制进 ops。
		draw(&ops, e.Source)
		// 更新显示。
		e.Frame(&ops)
	}
}
```

让我们让按钮每秒改变一次位置。我们用 [Ticker](https://golang.org/pkg/time#Ticker) 作为外部变化的例子。我们使用锁来保护状态，一旦修改了状态，就需要通过 [window.Invalidate()](https://gioui.org/app#Window.Invalidate) 通知窗口重新触发渲染。

```go
var window app.Window
window.Option(app.Title(title))

var button struct {
	lock   sync.Mutex
	offset int
}

updateOffset := func(v int) {
	button.lock.Lock()
	defer button.lock.Unlock()
	button.offset = v
}
readOffset := func() int {
	button.lock.Lock()
	defer button.lock.Unlock()
	return button.offset
}

go func() {
	changes := time.NewTicker(time.Second)
	defer changes.Stop()
	for t := range changes.C {
		updateOffset(int((t.Second() % 3) * 100))
		window.Invalidate()
	}
}()

ops := new(op.Ops)
for {
	switch e := window.Event().(type) {
	case app.DestroyEvent:
		return e.Err
	case app.FrameEvent:
		ops.Reset()

		// 根据状态偏移按钮。
		op.Offset(image.Pt(readOffset(), 0)).Add(ops)

		// 处理按钮输入并绘制。
		doButton(ops, e.Source)

		// 更新显示。
		e.Frame(ops)
	}
}
```

用这些概念写程序可能会变得非常冗长，这就是为什么 Gio 为常见的外观和行为提供了标准控件。大多数程序最终主要使用控件，很少直接使用底层操作。

## Advanced Input Topics（高级输入主题）

此标题下的内容探索了 Gio 输入操作的更高级用法。这些内容主要对编写自定义控件的人有用，对于使用 Gio 高层控件和布局 API 来说并不是严格必需的。

### Input Tree（输入树）

你可能已经注意到，前面的例子使用了 `clip.AreaOp`（通过 `clip.Rect` 构造）来描述它想要指针输入的位置。这是因为 Gio 使用 `clip.AreaOp` 同时描述绘制区域和输入区域。正如你上面看到的，你经常既想在某个区域内绘制，又想在该区域内接受输入，所以这种复用很方便。

`clip.AreaOp` 形成了一棵隐式的输入区域树，每个区域都可能对指针输入、键盘输入或两者都感兴趣。

下面是一个探索指针事件如何与这棵树结构交互的例子。

```go
var (
	// 声明一些变量，同时用作状态和输入 tag。
	root, child1, child2 bool
)

// displayForTag 使用给定的 tag，向给定的 op.Ops 添加一个对按压和释放事件感兴趣的 pointer.InputOp。
// 它还会根据 tag 的当前值，用颜色填充当前裁剪区域。
func displayForTag(ops *op.Ops, tag *bool, rect clip.Rect) {
	event.Op(ops, tag)

	// 根据 tag 是否被按压选择颜色。
	c := color.NRGBA{B: 0xFF, A: 0xFF}
	if *tag {
		c = color.NRGBA{R: 0xFF, A: 0xFF}
	}
	// 用半透明颜色填充当前裁剪区域。
	translucent := c
	translucent.A = 0x44
	paint.ColorOp{Color: translucent}.Add(ops)
	paint.PaintOp{}.Add(ops)

	// 把裁剪区域缩小到矩形的轮廓，然后绘制该轮廓。这应该更容易看清重叠的矩形。
	defer clip.Stroke{
		Path:  rect.Path(),
		Width: 5,
	}.Op().Push(ops).Pop()
	paint.ColorOp{Color: c}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func doPointerTree(ops *op.Ops, q input.Source) {
	// 为每个 tag 处理上一帧到这一帧之间到达的事件。
	for _, tag := range []*bool{&root, &child1, &child2} {
		for {
			ev, ok := q.Event(pointer.Filter{
				Target: tag,
				Kinds:  pointer.Press | pointer.Release,
			})
			if !ok {
				break
			}

			x, ok := ev.(pointer.Event)
			if !ok {
				continue
			}

			switch x.Kind {
			case pointer.Press:
				*tag = true
			case pointer.Release:
				*tag = false
			}
		}
	}

	// 把根区域的兴趣范围限制在一个 200x200 的矩形内。
	rootRect := clip.Rect(image.Rect(0, 0, 200, 200))
	rootArea := rootRect.Push(ops)
	displayForTag(ops, &root, rootRect)

	// 在 Pop 根区域之前添加的任何裁剪区域都被视为它的子区域。
	child1Rect := clip.Rect(image.Rect(25, 25, 175, 100))
	child1Area := child1Rect.Push(ops)
	displayForTag(ops, &child1, child1Rect)
	child1Area.Pop()

	child2Rect := clip.Rect(image.Rect(100, 25, 175, 175))
	child2Area := child2Rect.Push(ops)
	displayForTag(ops, &child2, child2Rect)
	child2Area.Pop()

	rootArea.Pop()
	// 现在我们添加的任何东西都_不是_ rootArea 的子区域。
}
```

试着点击三个蓝色矩形中的每一个。你应该会看到：点击最大的矩形只会让它自己变红，而点击它内部的两个矩形中的任何一个，都会让你点击的那个矩形**以及**最外层的矩形一起变红。

这是因为指针输入事件会沿着 `clip.AreaOp` 树向上传播，寻找对该类事件感兴趣的 `pointer.Filter`。它们不会在第一个感兴趣的 `pointer.Filter` 处停止，而是会一直向上传到树的根。这意味着我们点击的矩形和包含它的矩形都会收到来自点击嵌套矩形的 `pointer.Press` 和 `pointer.Release`。

还要注意，如果你点击两个子矩形重叠的区域，只有最上层的（最后绘制的）矩形会收到点击。默认情况下，Gio 在路由指针事件时只考虑最前面的区域及其祖先。如果你想改变这一点，可以使用 `pointer.PassOp` 让指针事件穿过一个输入区域到达它下面的区域。这对于布局覆盖层和类似元素很有用。关于这个操作的细节，请参见 [pointer 包的文档](https://pkg.go.dev/gioui.org/io/pointer#hdr-Pass_through)。

---

### 深度解读

**1. 即时模式输入的核心哲学**

传统 GUI（Qt、WinForms、Flutter）有控件树，框架自动把事件路由到正确的控件。  
Gio **没有**这种树。它完全是无状态的。

每一帧你必须：
1. 用 `event.Op(ops, tag)` **主动注册**“我对这个区域的这些事件感兴趣”
2. 在**下一帧**用 `q.Event(Filter{Target: tag})` 把事件取回来

这就是为什么它叫“即时模式”：输入兴趣和绘制都是每帧重新声明的。

**2. Tag 是关键抽象**

Tag 本质上就是一个 `interface{}`，但实践中几乎永远用**指针**：

```go
type MyButton struct {
    pressed bool
}

// 在 Layout 里
event.Op(ops, b)          // b 是 *MyButton
// 取事件
q.Event(pointer.Filter{Target: b, ...})
```

用指针的好处：
- 转换成本几乎为零
- 天然唯一（每个实例地址不同）
- 可以直接关联到控件状态

**3. 输入树 = 裁剪树**

这是 Gio 设计最精妙的地方之一：

- 绘制用的 `clip.Rect / RRect / Path` 
- 输入命中区域用的也是同一套东西

所以你写：

```go
defer clip.Rect{...}.Push(ops).Pop()
event.Op(ops, tag)
// 绘制
```

就同时完成了“这个区域既画东西，也接收输入”。

事件会**向上冒泡**到所有祖先区域。这就是为什么点击子区域时，父区域也能收到事件。

**4. 默认命中规则与 PassOp**

默认行为：
- 只命中最上层（最后绘制）的区域 + 它的所有祖先
- 下面的兄弟区域收不到事件

想让事件穿透（比如半透明遮罩、工具提示），就用：

```go
pointer.PassOp{}.Add(ops)
```

**5. 外部输入必须主动 Invalidate**

Gio 的主循环只在“系统认为需要重绘”时才给你 FrameEvent。  
定时器、网络、文件变化等外部事件，**必须**自己调用：

```go
window.Invalidate()
```

否则界面不会更新。这是新手最容易踩的坑之一。

**6. 实际开发建议**

| 场景           | 推荐做法                                              |
| -------------- | ----------------------------------------------------- |
| 简单按钮       | 直接用 `event.Op` + 自己处理 Press/Release            |
| 复杂手势       | 使用 `gioui.org/gesture` 包（Click、Scroll、Drag 等） |
| 自定义控件     | 把 `*自己` 作为 tag，封装在 Layout 方法里             |
| 多层叠加       | 注意输入树冒泡 + 必要时用 PassOp                      |
| 跨线程修改状态 | 加锁 + Invalidate                                     |

**7. 与第2章的紧密联系**

绘制章节的 `clip` 操作，在本章被复用为输入区域。  
这就是为什么 Gio 的绘制和输入感觉“长在一起”——它们共享同一套区域描述。

---

**本章核心记忆点**

> **输入 = 每帧主动注册兴趣区域（event.Op） + 下一帧按 tag 取事件**  
> **裁剪区域同时定义绘制和命中**  
> **事件向上冒泡**  
> **外部变化必须 Invalidate**

---

准备好后，回复“继续第4章”，我会继续给出 Widget 章节的完整翻译与深度解读。