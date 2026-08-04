package main

import (
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
			view.Text(th, "Hello Gio"),
			view.Spacer(20),
			view.Text(th, "Goro UI"),
		),
	)
}
