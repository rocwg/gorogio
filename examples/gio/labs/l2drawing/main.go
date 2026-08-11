package main

import (
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/op"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("Test Window"))

		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	var ops op.Ops
	//th := internal.NewTheme()

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ops.Reset()

			//gtx := app.NewContext(&ops, e)
			//drawRedRect(&ops)
			//drawRedRect10PixelsRight(&ops)
			//redButtonBackground(&ops)
			//redTriangle(&ops)
			//strokeRect(&ops)
			//strokeTriangle(&ops)
			//redButtonBackgroundStack(&ops)
			//drawOverlappingRectangles(&ops)
			//drawFiveRectangles(&ops)

			drawProgressBar(&ops, e.Source, time.Now())

			e.Frame(&ops)
		}
	}
}

//drawWithCache(&ops)
//drawImage(&ops, image.NewRGBA(image.Rect(0, 0, 800, 600)))
