package main

// A simple Gio program. See https://gioui.org for more information.

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(
			app.Decorated(false), // 关键点：开启自定义边框/无边框模式！
		)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	// 定义主窗口的按钮状态
	var openBtn widget.Clickable
	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// 1. 监听按钮点击事件
			if openBtn.Clicked(gtx) {
				// 触发点击后，在一个独立的 goroutine 里创建全新的 OS 子窗口
				openSubWindow(th)
			}

			// 2. 绘制主窗口界面 (垂直居中布局：标题 + 按钮)
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						l := material.H4(th, "Hello, Gio")
						l.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255}
						l.Alignment = text.Middle
						return l.Layout(gtx)
					}),
					// 增加一点内边距 (Padding)
					layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						// 渲染 Material 风格按钮
						btn := material.Button(th, &openBtn, "打开独立子窗口")
						return btn.Layout(gtx)
					}),
				)
			})

			e.Frame(gtx.Ops)
		}
	}
}
