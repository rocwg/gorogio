package view

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Button struct {
	click widget.Clickable
}

func (b *Button) Layout(
	th *material.Theme,
	gtx layout.Context,
) layout.Dimensions {

	return material.Button(
		th,
		&b.click,
		"Save",
	).Layout(gtx)

}
