package component

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/style"
)

// Button 第一个 UI 单元
//
// 对应:
//
// Compose: Button()
// SwiftUI: Button
// Flutter: ElevatedButton
type Button struct {
	click widget.Clickable
	text  string
}

// NewButton 创建按钮
func NewButton(
	text string,
) *Button {

	return &Button{
		text: text,
	}
}

// Element 根据 Theme 渲染 Button
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

// Clicked 查询按钮是否点击
func (b *Button) Clicked(
	gtx layout.Context,
) bool {

	return b.click.Clicked(gtx)
}
