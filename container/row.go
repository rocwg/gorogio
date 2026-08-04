package container

import (
	"gioui.org/layout"

	"github.com/rocwg/gorogio/element"
)

// Row 水平布局容器。
//
// 对应:
//
// Compose: Row
// SwiftUI: HStack
// Flutter: Row
// Gio: layout.Flex + Horizontal
func Row(
	options Options,
	children ...element.Element,
) element.Element {

	return flexLayout(
		layout.Horizontal,
		options,
		children,
	)
}
