package main

import (
	"image"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
)

// (2) Clipping
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
