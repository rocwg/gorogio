package container

import (
	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/rocwg/gorogio/element"
)

// VerticalSpace
//
// 垂直方向占用空间。
//
// 类似:
// Compose: Spacer(height)
// SwiftUI: Spacer()
func VerticalSpace(
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

// HorizontalSpace
//
// 水平方向占据固定空间。
func HorizontalSpace(
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
