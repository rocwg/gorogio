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

// HelloPage
//
// 页面对象。
//
// 负责:
//
// 1. 页面状态
// 2. 页面事件
// 3. 页面 UI

type HelloPage struct {
	State *CounterState

	Increment *component.Button
	Reset     *component.Button
}

// NewHelloPage
//
// 创建页面。
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

// Update
//
// 生命周期:
//
// Application
//
//	|
//	v
//
// Page.Update()
//
// 处理输入事件。
func (p *HelloPage) Update(
	gtx layout.Context,
) {

	p.Increment.Update(gtx)

	p.Reset.Update(gtx)

}

// build
//
// 构建 UI Tree。
func (p *HelloPage) build(
	th *style.Theme,
) element.Element {

	return modifier.Padding(

		40,

		container.Column(

			container.Options{
				Spacing:   th.Spacing.Large,
				Alignment: layout.Middle,
			},

			view.H3(
				th,
				"Hello Gio",
			),

			view.Body(
				th,
				"Count : "+strconv.Itoa(p.State.Count),
			),

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
// Page 对外暴露的 UI Element。
//
// 注意:
//
// 不负责 Update。
// 只负责 UI。
func (p *HelloPage) Element(
	th *style.Theme,
) element.Element {

	return p.build(th)
}
