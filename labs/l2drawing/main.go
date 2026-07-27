package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/input"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
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

			r := new(input.Router)
			drawProgressBar(&ops, r.Source(), time.Now())

			e.Frame(&ops)
		}
	}
}

func drawRedRect(ops *op.Ops) {
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func drawRedRect10PixelsRight(ops *op.Ops) {
	defer op.Offset(image.Pt(100, 0)).Push(ops).Pop()
	drawRedRect(ops)
}

func redButtonBackground(ops *op.Ops) {
	const r = 10 // roundness
	bounds := image.Rect(0, 0, 100, 100)
	clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Push(ops)
	drawRedRect(ops)
}

func redTriangle(ops *op.Ops) {
	var path clip.Path
	path.Begin(ops)
	path.Move(f32.Pt(50, 0))
	path.Quad(f32.Pt(0, 90), f32.Pt(50, 100))
	path.Line(f32.Pt(-100, 0))
	path.Line(f32.Pt(50, -100))
	defer clip.Outline{Path: path.End()}.Op().Push(ops).Pop()
	drawRedRect(ops)
}

func strokeRect(ops *op.Ops) {
	const r = 10
	bounds := image.Rect(20, 20, 80, 80)
	rrect := clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}

	red := color.NRGBA{R: 0x80, A: 0xFF}
	paint.FillShape(ops, red,
		clip.Stroke{
			Path:  rrect.Path(ops),
			Width: 4,
		}.Op(),
	)
}

func strokeTriangle(ops *op.Ops) {
	var path clip.Path
	path.Begin(ops)
	path.MoveTo(f32.Pt(30, 30))
	path.LineTo(f32.Pt(70, 30))
	path.LineTo(f32.Pt(50, 70))
	path.Close()

	green := color.NRGBA{R: 0x80, A: 0xFF}
	paint.FillShape(ops, green,
		clip.Stroke{
			Path:  path.End(),
			Width: 4,
		}.Op())
}

var startTime = time.Now()
var duration = 10 * time.Second

func drawProgressBar(ops *op.Ops, source input.Source, now time.Time) {
	// Calculate how much of the progress bar to draw,
	// based on the current time.
	elapsed := now.Sub(startTime)
	progress := elapsed.Seconds() / duration.Seconds()
	if progress < 1 {
		// The progress bar hasn’t yet finished animating.
		source.Execute(op.InvalidateCmd{})
	} else {
		progress = 1
	}

	width := 200 * float32(progress)
	defer clip.Rect{Max: image.Pt(int(width), 20)}.Push(ops).Pop()
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
