package main

import (
	"image"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
)

func draw6(window *app.Window) error {
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
			spacer(gtx)

			// 更新显示。
			e.Frame(&ops)
		}
	}
}

func spacer(gtx layout.Context) layout.Dimensions {
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return colorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Rigid(layout.Spacer{Width: 20}.Layout), // 插入 20px 间距
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return colorBox(gtx, gtx.Constraints.Min, blue)
		}),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return colorBox(gtx, image.Pt(100, 100), red)
		}),
		layout.Rigid(layout.Spacer{Width: 20}.Layout),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return colorBox(gtx, gtx.Constraints.Min, green)
		}),
	)
}
