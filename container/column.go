package container

import (
	"gioui.org/layout"

	"github.com/rocwg/gorogio/element"
)

// Column 垂直布局容器。
//
// 对应:
//
// Compose: Column
// SwiftUI: VStack
// Flutter: Column
// Gio: layout.Flex + Vertical
func Column(
	options Options,
	children ...element.Element,
) element.Element {

	return flexLayout(
		layout.Vertical,
		options,
		children,
	)
}
