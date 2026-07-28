package view

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type TopBar struct {
	Search    widget.Editor
	SearchBtn widget.Clickable
	AddBtn    widget.Clickable
}

func NewTopBar() *TopBar {
	t := &TopBar{}
	t.Search.SingleLine = true
	t.Search.SetText("")
	return t
}

func (t *TopBar) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx)
}
