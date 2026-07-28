package view

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ResourceTable struct {
	RowClicks []widget.Clickable
}

func NewResourceTable(n int) *ResourceTable {
	return &ResourceTable{
		RowClicks: make([]widget.Clickable, n),
	}
}

func (t *ResourceTable) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H6(th, "Resources").Layout(gtx)
		}),
	)
}
