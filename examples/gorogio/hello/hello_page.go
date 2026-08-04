package main

import (
	"strconv"

	"gioui.org/layout"

	"github.com/rocwg/gorogio/component"
	"github.com/rocwg/gorogio/container"
	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/modifier"
	"github.com/rocwg/gorogio/page"
	"github.com/rocwg/gorogio/primitive"
	"github.com/rocwg/gorogio/style"
)

// compile-time interface check.
//
// 如果 HelloPage 不满足 page.Page
// 编译直接失败。
var _ page.Page = (*HelloPage)(nil)

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

	hello := &HelloPage{
		State: state,
	}

	hello.Increment =
		component.NewButton("+").
			OnClick(
				func() {
					hello.State.Increment()
				},
			)

	hello.Reset =
		component.NewButton("Reset").
			OnClick(
				func() {
					hello.State.Reset()
				},
			)

	return hello
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

			primitive.H3(
				th,
				"Hello Gio",
			),

			primitive.Body(
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
