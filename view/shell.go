package view

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/rocwg/gorogio/state"
)

type Shell struct {
	Sidebar *Sidebar
	TopBar  *TopBar
	Table   *ResourceTable
	Dialog  *EditDialog
}

func NewShell(state *state.AppState) *Shell {
	items := []string{"resource", "users", "settings"}
	return &Shell{
		Sidebar: NewSidebar(items, state.Nav),
		TopBar:  NewTopBar(),
		Table:   NewResourceTable(len(state.Resources)),
		Dialog:  NewEditDialog(),
	}
}

func (s *Shell) Layout(gtx layout.Context, th *material.Theme, state *state.AppState) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return s.Sidebar.Layout(gtx, th)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return s.TopBar.Layout(gtx, th)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return s.Table.Layout(gtx, th)
				}),
			)
		}),
	)
}
