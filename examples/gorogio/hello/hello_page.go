package main

import (
	"gioui.org/layout"
	"github.com/rocwg/gorogio/container"
	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/modifier"
	"github.com/rocwg/gorogio/style"
	"github.com/rocwg/gorogio/view"
)

type HelloPage struct {
	State *CounterState
}

func NewHelloPage(
	state *CounterState,
) *HelloPage {

	return &HelloPage{
		State: state,
	}
}

func (p *HelloPage) Element(
	th *style.Theme,
) element.Element {

	return modifier.Padding(
		40,
		container.Column(
			container.Options{
				Spacing:   50,
				Alignment: layout.Middle,
			},
			view.Text(th, "Hello Gio"),
			view.Text(th, "Goro UI"),
		),
	)
}
