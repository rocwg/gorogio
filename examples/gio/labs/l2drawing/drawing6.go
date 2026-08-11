package main

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

//(7) Record and replay

func drawWithCache(ops *op.Ops) {
	// Save the operations in an independent ops value (the cache).
	cache := new(op.Ops)
	macro := op.Record(cache)

	cl := clip.Rect{Max: image.Pt(100, 100)}.Push(cache)
	paint.ColorOp{Color: color.NRGBA{G: 0x80, A: 0xFF}}.Add(cache)
	paint.PaintOp{}.Add(cache)
	cl.Pop()
	call := macro.Stop()

	// Draw the operations from the cache.
	call.Add(ops)
}

//(8) Images

func drawImage(ops *op.Ops, img image.Image) {
	imageOp := paint.NewImageOp(img)
	imageOp.Filter = paint.FilterNearest
	imageOp.Add(ops)
	op.Affine(f32.Affine2D{}.Scale(f32.Pt(0, 0), f32.Pt(4, 4))).Add(ops)
	paint.PaintOp{}.Add(ops)
}
