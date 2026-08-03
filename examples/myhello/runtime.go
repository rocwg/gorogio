package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
)

func Run() {
	go func() {
		w := new(app.Window)
		w.Option(
			app.Title("Hello goro-ui"),
		)
		err := runWindow(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func runWindow(
	w *app.Window,
) error {
	application := NewApplication()
	var ops op.Ops

	for {
		switch e := w.Event().(type) {

		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			application.Draw(gtx)

			e.Frame(gtx.Ops)
		}
	}
}
