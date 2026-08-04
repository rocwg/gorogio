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
// Gio:	Flex + Vertical
func Column(
	children ...element.Element,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		items := make(
			[]layout.FlexChild,
			0,
			len(children),
		)

		for _, child := range children {
			items = append(
				items,
				layout.Rigid(child),
			)
		}

		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(
			gtx,
			items...,
		)
	}
}
