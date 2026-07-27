package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
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
		// 关键配置：app.Decorated(false) 开启无边框沉浸模式！
		w.Option(
			app.Title("Gio IDE Style Window"),
			app.Size(unit.Dp(900), unit.Dp(600)),
			app.Decorated(false), // 隐藏系统默认标题栏！
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

	// 自定义高级暗黑 Theme (Catppuccin 风格)
	th.Palette = material.Palette{
		Bg:         color.NRGBA{R: 0x1E, G: 0x1E, B: 0x2E, A: 0xFF}, // 主背景漆黑
		Fg:         color.NRGBA{R: 0xCD, G: 0xD6, B: 0xF4, A: 0xFF}, // 主文字白色
		ContrastBg: color.NRGBA{R: 0x89, G: 0xB4, B: 0xFA, A: 0xFF}, // 高亮蓝色
	}

	var closeBtn widget.Clickable
	var minBtn widget.Clickable
	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// 响应自制标题栏中的“关闭”与“最小化”按钮点击
			if closeBtn.Clicked(gtx) {
				//w.Perform(app.DestroyCmd{})
				w.Perform(system.ActionClose)
			}
			if minBtn.Clicked(gtx) {
				//w.Perform(app.MinimizeCmd{})
				w.Perform(system.ActionMinimize)
			}

			// 全屏 Flex 垂直布局：[ 顶部自制标题栏 ] + [ 中间工作区 ] + [ 底部状态栏 ]
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				// 1. 顶部标题栏 (跟随 Theme 背景)
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return drawCustomTitleBar(gtx, th, w, &minBtn, &closeBtn)
				}),
				// 2. 主工作区 (填满剩余空间)
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						l := material.H5(th, "Zed / GoLand 风格沉浸式工作区")
						return l.Layout(gtx)
					})
				}),
				// 3. 底部状态栏 (稍微深一点的背景)
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return drawStatusBar(gtx, th)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}

// 绘制顶部自定义标题栏
func drawCustomTitleBar(gtx layout.Context, th *material.Theme, w *app.Window, minBtn, closeBtn *widget.Clickable) layout.Dimensions {
	//height := gtx.Dp(unit.Dp(38))

	// 1. 填充标题栏背景色（深一点的灰黑色 #181825）
	titleBg := color.NRGBA{R: 0x18, G: 0x18, B: 0x25, A: 0xFF}
	paint.FillShape(gtx.Ops, titleBg, clip.Rect{Max: gtx.Constraints.Max}.Op())

	// 2. 实现窗口拖拽：在整个标题栏区域叠加 Perform(app.DragCmd{})
	// 这样鼠标在标题栏空白处按住拖动，就能像系统原生标题栏一样移动整个窗口！
	/* Note: 可以在标题栏空白处绑定 pointer 输入并调用 w.Perform(app.DragCmd{}) */

	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		// 左侧图标与标题文字
		layout.Rigid(layout.Spacer{Width: unit.Dp(12)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			title := material.Body2(th, "My GoLand-like IDE")
			return title.Layout(gtx)
		}),
		// 中间弹性占位符 (把右侧控制按钮推到最右边)
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{Size: gtx.Constraints.Min}
		}),
		// 右侧最小化按钮
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, minBtn, " — ")
			btn.Background = color.NRGBA{A: 0} // 透明背景
			btn.Color = th.Fg
			return btn.Layout(gtx)
		}),
		// 右侧关闭按钮
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, closeBtn, " ✕ ")
			btn.Background = color.NRGBA{A: 0}                          // 透明背景
			btn.Color = color.NRGBA{R: 0xF3, G: 0x8B, B: 0x88, A: 0xFF} // 红色
			return btn.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
	)
}

// 绘制底部状态栏
func drawStatusBar(gtx layout.Context, th *material.Theme) layout.Dimensions {
	// 填充底部栏背景色 (#11111B)
	statusBg := color.NRGBA{R: 0x11, G: 0x11, B: 0x1B, A: 0xFF}
	paint.FillShape(gtx.Ops, statusBg, clip.Rect{Max: gtx.Constraints.Max}.Op())

	return layout.Inset{
		Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(12), Right: unit.Dp(12),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				txt := material.Caption(th, "Git: main*")
				return txt.Layout(gtx)
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return layout.Dimensions{Size: gtx.Constraints.Min}
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				txt := material.Caption(th, "UTF-8 | Go 1.27")
				return txt.Layout(gtx)
			}),
		)
	})
}
