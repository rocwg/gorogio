package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

//Gio Window + Event Loop

func Run() {
	go func() {
		window := new(app.Window)
		err := runWindow(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func runWindow(window *app.Window) error {
	//定义操作
	var ops op.Ops

	//监听事件
	for {
		switch e := window.Event().(type) {

		//app.DestroyEvent表示用户按下了关闭按钮。
		case app.DestroyEvent:
			return e.Err

		//app.FrameEvent这意味着程序应该处理输入并渲染新帧。
		//绘制文本
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			//UI 描述（未来 view 层的雏形）
			DrawUI(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}

//Theme 应该属于 Application，而不是 UI。
//State
//Component

type Application struct {
	Theme   *material.Theme
	Counter CounterState
	Hello   *HelloPage
}

func NewApplication() *Application {
	//创建主题
	theme := material.NewTheme()
	//状态
	state := CounterState{}
	//
	component := NewHelloComponent(&state)

	return &Application{
		Theme:   theme,
		Counter: state,
		Hello:   component,
	}
}

func (a *Application) Draw(
	gtx layout.Context,
) {
	//
	a.Hello.Layout(gtx, a.Theme)
}
