package primitive

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/rocwg/gorogio/style"

	"github.com/rocwg/gorogio/element"
)

// Text 最基础文本元素

func H3(
	th *style.Theme,
	value string,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		return material.H3(
			th.Material,
			value,
		).Layout(gtx)
	}
}

func Body(
	th *style.Theme,
	value string,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		return material.Body1(
			th.Material,
			value,
		).Layout(gtx)
	}
}
