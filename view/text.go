package view

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	"github.com/rocwg/gorogio/element"
)

// Text
//
// 最基础文本元素。
//
// 使用: view.Text("Hello")
// 返回: element.Element
//

func Text(
	value string,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		theme := material.NewTheme()

		return material.Body1(
			theme,
			value,
		).Layout(gtx)
	}
}
