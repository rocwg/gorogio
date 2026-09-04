package main

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type HelloGio struct {
	btn     widget.Clickable
	clicked bool
}

func NewHelloGio() *HelloGio {
	return &HelloGio{
		btn:     widget.Clickable{},
		clicked: false,
	}
}

func (h *HelloGio) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	if h.btn.Clicked(gtx) {
		h.clicked = !h.clicked
	}
	egButtonText(th, &h.btn).Layout(gtx)
	txt := "Hello, Gio"
	if h.clicked {
		txt = "Clicked!"
	}
	return egLabelText(th, txt).Layout(gtx)
}

func egButtonText(th *material.Theme, btn *widget.Clickable) material.ButtonStyle {
	return material.Button(th, btn, "Click me")
}

func egLabelText(th *material.Theme, lText string) material.LabelStyle {
	lt := material.H1(th, lText)
	lt.Alignment = text.Middle
	lt.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255}
	return lt
}
