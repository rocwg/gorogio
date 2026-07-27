package main

import (
	"image/color"
	"log"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// openSubWindow 创建并启动一个真正的 OS 级独立子窗口
func openSubWindow(th *material.Theme) {
	go func() {
		// 1. 向操作系统申请独立的窗口句柄
		w := new(app.Window)
		w.Option(
			app.Title("独立 Hello 弹窗"),
			app.Size(unit.Dp(350), unit.Dp(200)),
		)

		// 2. 属于这个子窗口独立的事件循环 (Event Loop)
		if err := loopSub(w); err != nil {
			log.Fatal(err)
		}
	}()
}

func loopSub(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			// 子窗口被关闭，直接退出当前 goroutine，主窗口完全不受影响
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// 绘制子窗口内容
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				l := material.H5(th, "Hello from Sub-Window!")
				l.Color = color.NRGBA{R: 0, G: 127, B: 70, A: 255}
				l.Alignment = text.Middle
				return l.Layout(gtx)
			})
			e.Frame(gtx.Ops)
		}
	}
}
