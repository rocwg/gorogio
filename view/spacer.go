package view

import (
	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/rocwg/gorogio/element"
)

func Spacer(
	height unit.Dp,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		return layout.Spacer{
			Height: height,
		}.Layout(gtx)
	}
}

func SpacerWidth(
	width unit.Dp,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		return layout.Spacer{
			Width: width,
		}.Layout(gtx)
	}
}
