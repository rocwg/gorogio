package view

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Sidebar struct {
	Items    []string
	Selected string
	Clicks   []widget.Clickable
}

func NewSidebar(items []string, selected string) *Sidebar {
	clicks := make([]widget.Clickable, len(items))
	return &Sidebar{Items: items, Selected: selected, Clicks: clicks}
}

func (s *Sidebar) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H6(th, "Admin").Layout(gtx)
		}),
	)
}
