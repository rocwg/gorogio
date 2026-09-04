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

type WindowApp struct {
	title  string
	window *app.Window
	ops    op.Ops          // 窗口私有 Ops
	theme  *material.Theme // 窗口私有 Theme (线程安全的关键!)
}

func NewWindowApp(title string) *WindowApp {
	w := new(app.Window)
	w.Option(
		app.Title(title),
		app.Size(unit.Dp(400), unit.Dp(300)),
	)
	return &WindowApp{
		title:  title,
		window: w,
		// 每个窗口在创建时生成属于自己的 Theme 实例，避免并发读写字体 Shaper
		theme: material.NewTheme(),
	}
}

func (wa *WindowApp) Run() error {
	for {
		switch e := wa.window.Event().(type) {
		case app.DestroyEvent:
			fmt.Printf("窗口 [%s] 已关闭\n", wa.title)
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&wa.ops, e)

			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.H5(wa.theme, fmt.Sprintf("这是: %s", wa.title)).Layout(gtx)
			})

			e.Frame(gtx.Ops)
		}
	}
}

func main() {
	go func() {
		// 启动第一个窗口
		go func() {
			win1 := NewWindowApp("主窗口 (Window 1)")
			if err := win1.Run(); err != nil {
				log.Println(err)
			}
		}()

		// 启动第二个窗口
		go func() {
			win2 := NewWindowApp("子窗口 (Window 2)")
			if err := win2.Run(); err != nil {
				log.Println(err)
			}
		}()
	}()

	app.Main()
}
