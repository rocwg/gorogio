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
	theme *style.Theme
	click widget.Clickable
	text  string
}

// NewButton 创建按钮
func NewButton(
	theme *style.Theme,
	text string,
) *Button {

	return &Button{
		theme: theme,
		text:  text,
	}
}

// Element ？
func (b *Button) Element() element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		return material.Button(
			b.theme.Material,
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
