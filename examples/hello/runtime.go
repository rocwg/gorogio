package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
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
	//放在哪里？
	var application = NewApplication()

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
			application.Draw(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
