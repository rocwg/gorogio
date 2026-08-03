package component

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/rocwg/gorogio/element"
)

// Button
//
// goro-ui 第一个组件。
//
// 负责:
//
// 1. 点击状态
// 2. Material Button 渲染
//
// 对应:
//
// Compose: Button()
// SwiftUI: Button
// Flutter: ElevatedButton
//

type Button struct {
	click widget.Clickable

	text string

	theme *material.Theme
}

// NewButton
//
// 创建按钮
//

func NewButton(
	theme *material.Theme,
	text string,
) *Button {

	return &Button{
		theme: theme,
		text:  text,
	}
}

// Element
//
// 返回 goro-ui Element
//

func (b *Button) Element() element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		return material.Button(
			b.theme,
			&b.click,
			b.text,
		).Layout(gtx)
	}
}

// Clicked
//
// 查询按钮是否点击
//

func (b *Button) Clicked(
	gtx layout.Context,
) bool {

	return b.click.Clicked(gtx)
}
