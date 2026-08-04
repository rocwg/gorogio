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

	page := &HelloPage{
		State: state,
	}

	page.Increment =
		component.NewButton("+").
			OnClick(
				func() {
					page.State.Increment()
				},
			)

	page.Reset =
		component.NewButton("Reset").
			OnClick(
				func() {
					page.State.Reset()
				},
			)

	return page
}

// Event 层

// Update
//
// 处理页面输入事件。
func (p *HelloPage) Update(
	gtx layout.Context,
) {
	p.Increment.Update(gtx)
	p.Reset.Update(gtx)
}

// View 层
func (p *HelloPage) View(
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

// Element
//
// 构建页面 Element Tree。
func (p *HelloPage) Element(
	th *style.Theme,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		// 先处理事件，再渲染 UI。
		p.Update(gtx)

		return p.View(th)(gtx)
	}
}
