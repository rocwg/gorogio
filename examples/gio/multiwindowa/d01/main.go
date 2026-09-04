package main

import (
	"fmt"
	"log"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// 1. 定义单个窗口的上下文结构体
type WindowApp struct {
	title  string
	window *app.Window
	ops    op.Ops // 必须是窗口私有的 Ops
	theme  *material.Theme
}

func NewWindowApp(title string, theme *material.Theme) *WindowApp {
	w := new(app.Window)
	w.Option(
		app.Title(title),
		app.Size(unit.Dp(400), unit.Dp(300)),
	)
	return &WindowApp{
		title:  title,
		window: w,
		theme:  theme,
	}
}

// 2. 每一个窗口独立的事件循环逻辑
func (wa *WindowApp) Run() error {
	for {
		switch e := wa.window.Event().(type) {
		case app.DestroyEvent:
			// 当窗口关闭时，只退出当前窗口的 Loop
			fmt.Printf("窗口 [%s] 已关闭\n", wa.title)
			return e.Err

		case app.FrameEvent:
			// 使用窗口私有的 &wa.ops
			gtx := app.NewContext(&wa.ops, e)

			// 绘制当前窗口特有的 UI
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.H5(wa.theme, fmt.Sprintf("这是: %s", wa.title)).Layout(gtx)
			})

			e.Frame(gtx.Ops)
		}
	}
}

func main() {
	// 主协程启动 OS UI 循环
	go func() {
		theme := material.NewTheme()

		// 启动第一个窗口 (Goroutine 1)
		go func() {
			win1 := NewWindowApp("主窗口 (Window 1)", theme)
			if err := win1.Run(); err != nil {
				log.Println(err)
			}
		}()

		// 启动第二个窗口 (Goroutine 2)
		go func() {
			win2 := NewWindowApp("子窗口 (Window 2)", theme)
			if err := win2.Run(); err != nil {
				log.Println(err)
			}
		}()
	}()

	app.Main()
}
