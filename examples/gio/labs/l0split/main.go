package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	//创建窗口
	go func() {
		window := new(app.Window)

		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme() //创建主题
	var ops op.Ops

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			// do
			//exampleSplitVisual(gtx, theme)
			//exampleSplitRatio(gtx, theme)
			exampleSplit(gtx, theme)

			// GPU
			e.Frame(gtx.Ops)
		}
	}
}
