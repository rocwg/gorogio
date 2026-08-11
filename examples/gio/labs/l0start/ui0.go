package main

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func myUI(gtx layout.Context, th *material.Theme) layout.Dimensions {

	// Define a large label with an appropriate text:
	title := material.H1(th, "Hello, Gio")

	// Change the color of the label.
	maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
	title.Color = maroon

	// Change the position of the label.
	title.Alignment = text.Middle

	// Draw the label to the graphics context.
	return title.Layout(gtx)
}
