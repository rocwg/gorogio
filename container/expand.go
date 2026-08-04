package container

import (
	"gioui.org/layout"

	"github.com/rocwg/gorogio/element"
)

// Expand 让子元素占据剩余空间。
//
// 对应:
//
// Compose: Modifier.weight(1)
// Flutter: Expanded
// SwiftUI: Spacer
// Gio: layout.Flexed
func Expand(
	child element.Element,
) layout.FlexChild {

	return layout.Flexed(
		1,
		child,
	)
}
