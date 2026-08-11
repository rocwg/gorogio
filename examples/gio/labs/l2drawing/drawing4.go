package main

import (
	"image"
	"image/color"

	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

// (5) Drawing Order
func drawOverlappingRectangles(ops *op.Ops) {
	// Draw a red rectangle.
	cl := clip.Rect{Max: image.Pt(100, 50)}.Push(ops)
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
	cl.Pop()

	// Draw a green rectangle.
	cl = clip.Rect{Max: image.Pt(50, 100)}.Push(ops)
	paint.ColorOp{Color: color.NRGBA{G: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
	cl.Pop()
}

func drawFiveRectangles(ops *op.Ops) {
	// Record drawRedRect operations into the macro.
	macro := op.Record(ops)
	drawRedRect(ops)
	c := macro.Stop()

	// “Play back” the macro 5 times, each time
	// translated vertically 20px and horizontally 110 pixels.
	for range 5 {
		c.Add(ops)
		op.Offset(image.Pt(110, 20)).Add(ops)
	}
}
