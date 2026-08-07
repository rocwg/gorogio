package container

import (
	"gioui.org/layout"

	"github.com/rocwg/gorogio/element"
)

// Center
//
// 让子元素水平垂直居中。
//
// Compose:
// Box(contentAlignment = Center)
//
// SwiftUI:
// frame + alignment
//
// Gio:
// layout.Center
func Center(
	child element.Element,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		return layout.Center.Layout(
			gtx,
			child,
		)
	}
}
