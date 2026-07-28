package view

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type EditDialog struct {
	Name     widget.Editor
	Type     widget.Editor
	Status   widget.Editor
	SaveBtn  widget.Clickable
	CloseBtn widget.Clickable
}

func NewEditDialog() *EditDialog {
	d := &EditDialog{}
	d.Name.SingleLine = true
	d.Type.SingleLine = true
	d.Status.SingleLine = true
	return d
}

func (d *EditDialog) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Inset{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return material.H6(th, "Edit Resource").Layout(gtx)
	})
}
