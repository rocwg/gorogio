package main

import (
	"image"

	"gioui.org/app"
	"gioui.org/op"
)

func run0(window *app.Window) error {
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// 上一帧作废
			ops.Reset()

			// ========== FRAME 区 ==========
			colorBox(&ops, image.Point{X: 100, Y: 100}, green)
			// ========== FRAME 区 ==========

			// 提交这一帧
			e.Frame(&ops)
		}
	}
}
