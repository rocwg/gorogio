package container

import (
	"gioui.org/layout"

	"github.com/rocwg/gorogio/element"
)

// Row
//
// 水平布局容器。
//
// 对应:
//
// Compose: Row
// SwiftUI: HStack
// Flutter: Row
//

func Row(
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
			Axis: layout.Horizontal,
		}.Layout(
			gtx,
			items...,
		)
	}
}
