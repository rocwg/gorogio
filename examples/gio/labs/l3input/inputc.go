package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/input"
	"gioui.org/io/pointer"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func draw3(window *app.Window) error {
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			// 窗口已被关闭。
			return e.Err
		case app.FrameEvent:
			// 绘制窗口状态的请求。

			// 将操作列表重置清零。
			ops.Reset()

			// 根据 e.Source 中的事件，将状态绘制到 ops 中。
			doPointerTree(&ops, e.Source)

			// 更新显示。
			e.Frame(&ops)
		}
	}
}

var (
	// 声明多个变量，同时用作状态和输入 tag。
	root, child1, child2 bool
)

// displayForTag 将对 press 和 release 事件感兴趣的 pointer.InputOp 添加到给定的 op.Ops 中，
// 使用指定的 tag。它还会根据 tag 的当前值向当前 clip 区域填充颜色。
func displayForTag(ops *op.Ops, tag *bool, rect clip.Rect) {
	event.Op(ops, tag)

	// 根据 tag 是否被按下选择颜色。
	c := color.NRGBA{B: 0xFF, A: 0xFF}
	if *tag {
		c = color.NRGBA{R: 0xFF, A: 0xFF}
	}
	// 用半透明颜色填充当前裁剪区域。
	translucent := c
	translucent.A = 0x44
	paint.ColorOp{Color: translucent}.Add(ops)
	paint.PaintOp{}.Add(ops)

	// 将裁剪区域缩小为矩形轮廓，然后绘制该轮廓。
	// 这有助于更容易地看出重叠的矩形。
	defer clip.Stroke{
		Path:  rect.Path(),
		Width: 5,
	}.Op().Push(ops).Pop()
	paint.ColorOp{Color: c}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func doPointerTree(ops *op.Ops, q input.Source) {
	// 为每个 tag 处理在上一次绘制帧和当前帧之间到达的事件。
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
			default:
				panic("unhandled default case")
			}
		}
	}

	// 将 rootArea 的兴趣区域限制为 200x200 的矩形。
	rootRect := clip.Rect(image.Rect(0, 0, 200, 200))
	rootArea := rootRect.Push(ops)
	displayForTag(ops, &root, rootRect)

	// 在 Pop-ing 根区域之前添加的任何 clip 区域都被视为它的子节点。
	child1Rect := clip.Rect(image.Rect(25, 25, 175, 100))
	child1Area := child1Rect.Push(ops)
	displayForTag(ops, &child1, child1Rect)
	child1Area.Pop()

	child2Rect := clip.Rect(image.Rect(100, 25, 175, 175))
	child2Area := child2Rect.Push(ops)
	displayForTag(ops, &child2, child2Rect)
	child2Area.Pop()

	rootArea.Pop()
	// 现在我们添加的任何东西都“不是” rootArea 的子节点。
}
