package modifier

import (
	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/rocwg/gorogio/element"
)

// Padding
//
// 给元素增加内边距。
//
// 对应:
//
// Compose: Modifier.padding()
// SwiftUI: padding()
// Flutter: Padding()
//
// Gio:
//
//	layout.Inset
//

func Padding(
	value unit.Dp,
	child element.Element,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		return layout.Inset{
			Top:    value,
			Bottom: value,
			Left:   value,
			Right:  value,
		}.Layout(gtx, child)
	}
}
