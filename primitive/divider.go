package primitive

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/style"
)

// Divider
//
// 一个基础视觉元素。
//
// 类似：
// Compose: Divider()
// SwiftUI: Divider()
// Flutter: Divider()
//
// 无状态。
func Divider(
	th *style.Theme,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		height := gtx.Dp(unit.Dp(1))

		size :=
			image.Point{
				X: gtx.Constraints.Max.X,
				Y: height,
			}

		defer clip.Rect{
			Max: size,
		}.Push(gtx.Ops).Pop()

		paint.Fill(
			gtx.Ops,
			th.DividerColor,
		)

		return layout.Dimensions{
			Size: size,
		}
	}
}
