package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
)

func draw1(window *app.Window) error {
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		// 窗口已被关闭。
		case app.DestroyEvent:
			return e.Err
		// 绘制窗口状态的请求。
		case app.FrameEvent:
			// 为新帧重置 layout.Context。
			gtx := app.NewContext(&ops, e)

			// 根据 e.Queue 中的事件，把状态绘制进 ops。
			b := Button{
				pressed: true,
			}
			b.Layout(gtx)

			// 更新显示。
			e.Frame(&ops)
		}
	}
}

type Button struct {
	pressed bool
}

func (b *Button) Layout(gtx layout.Context) layout.Dimensions {
	// 限制指针事件响应的区域为 100x100。
	area := clip.Rect(image.Rect(0, 0, 100, 100)).Push(gtx.Ops)

	// 将指针事件路由目标绑定到按钮指针 b 上。
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
		default:
			panic("unhandled default case")
		}
	}

	area.Pop()

	// 绘制按钮。
	col := color.NRGBA{R: 0x80, A: 0xFF}
	if b.pressed {
		col = color.NRGBA{G: 0x80, A: 0xFF} // 按下变绿，未按下为红
	}
	return drawSquare(gtx.Ops, col)
}
