package main

import (
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
	Counter *component.Counter
}

// NewHelloPage
//
// 创建页面。
func NewHelloPage() *HelloPage {

	return &HelloPage{
		Counter: component.NewCounter(),
	}
}

// Update
//
// 生命周期: Application -> Page.Update()
//
// 处理输入事件。
func (p *HelloPage) Update(
	gtx layout.Context,
) {
	p.Counter.Update(gtx)
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

			p.Counter.Element(th),
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
