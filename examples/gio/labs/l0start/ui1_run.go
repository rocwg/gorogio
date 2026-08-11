package main

import (
	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func run1(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops

	// 初始化应用状态与 UI 结构体
	ui := NewMainUI(theme)

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// 渲染整个 UI 树并获取最终尺寸
			ui.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}
