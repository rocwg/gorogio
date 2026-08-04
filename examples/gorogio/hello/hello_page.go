package main

import (
	"strconv"

	"gioui.org/layout"

	"github.com/rocwg/gorogio/component"
	"github.com/rocwg/gorogio/container"
	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/modifier"
	"github.com/rocwg/gorogio/style"
	"github.com/rocwg/gorogio/view"
)

type HelloPage struct {
	State *CounterState

	Increment *component.Button
	Reset     *component.Button
}

func NewHelloPage(
	state *CounterState,
) *HelloPage {

	return &HelloPage{
		State: state,

		Increment: component.NewButton("+"),
		Reset:     component.NewButton("Reset"),
	}
}

// update
//
// 处理页面事件。
func (p *HelloPage) update(
	gtx layout.Context,
) {

	if p.Increment.Clicked(gtx) {
		p.State.Increment()
	}
	if p.Reset.Clicked(gtx) {
		p.State.Reset()
	}
}

// Element
//
// 构建页面 Element Tree。
func (p *HelloPage) Element(
	th *style.Theme,
) element.Element {

	return modifier.Padding(
		40,
		container.Column(
			container.Options{
				Spacing:   th.Spacing.Large,
				Alignment: layout.Middle,
			},
			view.H3(th, "Hello Gio"),
			view.Body(th, "Count : "+strconv.Itoa(p.State.Count)),
			container.Row(
				container.Options{
					Spacing: th.Spacing.Large,
				},
				p.Increment.Element(th),
				p.Reset.Element(th),
			),
		),
	)
}
