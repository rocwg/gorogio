package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
)

func draw4(window *app.Window) error {
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
			listing(gtx)

			// 更新显示。
			e.Frame(&ops)
		}
	}
}

// 必须在函数体外定义，保持滚动状态持久
var list = layout.List{}

func listing(gtx layout.Context) layout.Dimensions {
	// 参数分别表示：gtx，元素总数，子项生成闭包
	return list.Layout(gtx, 100, func(gtx layout.Context, i int) layout.Dimensions {
		col := color.NRGBA{R: byte(i * 20), G: 0x20, B: 0x20, A: 0xFF}
		return colorBox(gtx, image.Pt(20, 100), col)
	})
}
