package main

import (
	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/rocwg/gorogio/app"
	"github.com/rocwg/gorogio/component"
	"github.com/rocwg/gorogio/container"
	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/modifier"
	"github.com/rocwg/gorogio/primitive"
	"github.com/rocwg/gorogio/style"
)

// compile-time interface check.
//
// 如果 HelloScreen 不满足 app.Scene
// 编译直接失败。
var _ app.Scene = (*HelloScreen)(nil)

// HelloPage
//
// 页面对象。
//
// 负责:
//
// 1. 页面状态
// 2. 页面事件
// 3. 页面 UI

type HelloScreen struct {
	Counter *component.Counter
}

// NewHelloPage
//
// 创建页面。
func NewHelloPage() *HelloScreen {

	return &HelloScreen{
		Counter: component.NewCounter(),
	}
}

// Update
//
// 生命周期: Application -> Page.Update()
//
// 处理输入事件。
func (h *HelloScreen) Update(
	gtx layout.Context,
) {
	h.Counter.Update(gtx)
}

// build
//
// 构建 UI Tree。
func (h *HelloScreen) build(
	th *style.Theme,
) element.Element {

	return container.Center(

		modifier.Padding(

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

				container.VerticalSpace(
					unit.Dp(20),
				),

				h.Counter.Element(th),
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
func (h *HelloScreen) Element(
	th *style.Theme,
) element.Element {

	return h.build(th)
}

func (h *HelloScreen) Name() string {
	return "hello"
}
