package main

import (
	"image"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
)

func main() {
	go func() {
		// 1. Window 建议用 new()，直接拿到 *app.Window 指针
		window := new(app.Window)
		window.Option(app.Title("Gio App"))

		err := run(window)
		//err := run0(window)
		//err := run1(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	// 2. Ops 建议用 var 声明值类型，零分配、语义清爽
	var ops op.Ops

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// 取 &ops 传入 Context，内部自动 Reset() 缓冲区
			gtx := app.NewContext(&ops, e)

			// 执行你的 UI 逻辑
			colorBox(gtx.Ops, image.Point{X: 100, Y: 100}, red)

			// 提交渲染指令
			e.Frame(gtx.Ops)
		}
	}
}
