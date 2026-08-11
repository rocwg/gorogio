package main

import (
	"image"

	"gioui.org/op"
	"gioui.org/op/clip"
)

// (4) Operation Stack
func redButtonBackgroundStack(ops *op.Ops) {
	const r = 1 // roundness
	bounds := image.Rect(0, 0, 100, 100)
	defer clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Push(ops).Pop()
	drawRedRect(ops)
}
