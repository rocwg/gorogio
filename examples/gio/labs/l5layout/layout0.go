package main

import (
	"image"

	"gioui.org/app"
	"gioui.org/op"
)

func draw0(window *app.Window) error {
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
			colorBox(gtx, image.Point{X: 100, Y: 100}, green)

			// 更新显示。
			e.Frame(&ops)
		}
	}
}
