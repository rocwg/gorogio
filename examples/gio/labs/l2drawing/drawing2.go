package main

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

// (3) Lines
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
