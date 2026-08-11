package main

import (
	"image"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
)

func draw2(window *app.Window) error {
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
			stacked(gtx)

			// 更新显示。
			e.Frame(&ops)
		}
	}
}

func stacked(gtx layout.Context) layout.Dimensions {
	return layout.Stack{}.Layout(gtx,
		// 强制控件与第二个相同大小。
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			// 这将有 100x100 的最小约束。
			return colorBox(gtx, gtx.Constraints.Min, red)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return colorBox(gtx, image.Pt(100, 30), green)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return colorBox(gtx, image.Pt(30, 100), blue)
		}),
	)
}
