package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"

	"github.com/rocwg/gorogio/container"
	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/modifier"
	"github.com/rocwg/gorogio/style"
	"github.com/rocwg/gorogio/view"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("My App"))
		if err := run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	th := style.NewTheme()
	var ops op.Ops

	for {
		e := w.Event()
		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			ui := page(th)
			ui(gtx)

			e.Frame(gtx.Ops)
		}
	}
}

func page(
	th *style.Theme,
) element.Element {

	return container.Column(

		container.Options{
			Spacing:   16,
			Alignment: layout.Middle,
		},

		view.Text(th, "Hello Gio"),
		view.Text(th, "Goro UI"),

		modifier.Padding(
			30,
			view.Text(th, "Goro UI"),
		),
	)
}
