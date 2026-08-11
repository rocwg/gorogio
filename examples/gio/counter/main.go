package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("Custom Widget Demo"), app.Size(unit.Dp(400), unit.Dp(300)))
		if err := run(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops

	// 1. 在事件循环外部实例化有状态组件（持久化状态）
	// 使用 Option 自定义参数：初始值 10，步长 5，范围 0-50，绿色按钮
	customCounter := NewCounter(10,
		WithStep(5),
		WithRange(0, 50),
		WithButtonColor(color.NRGBA{R: 46, G: 139, B: 87, A: 255}),
	)

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// 居中绘制我们的自定义控件
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// 2. 调用自定义组件的 Layout 方法
				return customCounter.Layout(gtx, theme)
			})

			e.Frame(gtx.Ops)
		}
	}
}
