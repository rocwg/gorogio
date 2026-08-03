package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system" // 窗口动作包
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(
			app.Title("Gio Modern Window"),
			app.Size(unit.Dp(900), unit.Dp(600)),
			app.Decorated(false), // 沉浸式无边框
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

	var closeBtn widget.Clickable
	var minBtn widget.Clickable
	var maxBtn widget.Clickable

	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// 1. 响应窗口控制按钮动作
			if closeBtn.Clicked(gtx) {
				w.Perform(system.ActionClose)
			}
			if minBtn.Clicked(gtx) {
				w.Perform(system.ActionMinimize)
			}
			if maxBtn.Clicked(gtx) {
				w.Perform(system.ActionMaximize)
			}

			// 2. 界面根布局：顶部标题栏 + 中央内容区
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				// 顶部标题栏
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return drawModernTitleBar(gtx, th, &minBtn, &maxBtn, &closeBtn)
				}),

				// 中央工作区
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.H5(th, "适配最新 Gio API 的沉浸式窗口").Layout(gtx)
					})
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}

// 绘制适配最新 Gio API 的自定义标题栏
func drawModernTitleBar(gtx layout.Context, th *material.Theme, minBtn, maxBtn, closeBtn *widget.Clickable) layout.Dimensions {
	titleHeight := gtx.Dp(unit.Dp(38))
	gtx.Constraints.Min.Y = titleHeight
	gtx.Constraints.Max.Y = titleHeight

	// 1. 绘制标题栏背景色 (#181825)
	titleBg := color.NRGBA{R: 0x18, G: 0x18, B: 0x25, A: 0xFF}
	paint.FillShape(gtx.Ops, titleBg, clip.Rect{Max: gtx.Constraints.Max}.Op())

	// 2. 核心修正：强转为 system.ActionInputOp，标记此 clip 区域为 OS 拖拽区
	defer clip.Rect{Max: gtx.Constraints.Max}.Op().Push(gtx.Ops).Pop()
	system.ActionInputOp(system.ActionMove).Add(gtx.Ops) // 👈 这里加上类型强转即可！

	// 3. 绘制标题文本与控制按钮
	return layout.Inset{
		Left: unit.Dp(12), Right: unit.Dp(8),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Body2(th, "My App").Layout(gtx)
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return layout.Dimensions{Size: gtx.Constraints.Min}
			}),
			// 最小化按钮
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, minBtn, " — ")
				btn.Background = color.NRGBA{A: 0}
				return btn.Layout(gtx)
			}),
			// 最大化按钮
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, maxBtn, " □ ")
				btn.Background = color.NRGBA{A: 0}
				return btn.Layout(gtx)
			}),
			// 关闭按钮
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, closeBtn, " ✕ ")
				btn.Background = color.NRGBA{A: 0}
				btn.Color = color.NRGBA{R: 0xFF, G: 0x55, B: 0x55, A: 0xFF}
				return btn.Layout(gtx)
			}),
		)
	})
}
