package component

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/style"
)

// Button
//
// 一个有状态 UI Component。
//
// state: Clickable
// event: onClick
// render: Element()
//
// 对应:
//
// Compose: Button()
// SwiftUI: Button
// Flutter: ElevatedButton
type Button struct {
	click widget.Clickable

	text string

	onClick func()
}

// NewButton
//
// 创建 Button。
func NewButton(
	text string,
) *Button {

	return &Button{
		text: text,
	}
}

// OnClick
//
// 注册点击行为。
func (b *Button) OnClick(
	handler func(),
) *Button {

	b.onClick = handler
	return b
}

// Update
//
// 处理输入事件。
//
// 每一帧调用。
func (b *Button) Update(
	gtx layout.Context,
) {

	for b.click.Clicked(gtx) {

		if b.onClick != nil {
			b.onClick()
		}
	}
}

// Element
//
// 构建 UI Element。
func (b *Button) Element(
	th *style.Theme,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		return material.Button(
			th.Material,
			&b.click,
			b.text,
		).Layout(gtx)
	}
}
