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

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		flex := layout.Flex{
			Axis:      layout.Horizontal,
			Alignment: options.Alignment,
		}

		items := make(
			[]layout.FlexChild,
			0,
			len(children)*2,
		)

		for index, child := range children {

			if index > 0 &&
				options.Spacing > 0 {

				items = append(
					items,

					layout.Rigid(
						func(
							gtx layout.Context,
						) layout.Dimensions {

							return layout.Spacer{
								Width: options.Spacing,
							}.Layout(gtx)
						},
					),
				)
			}

			items = append(
				items,
				layout.Rigid(child),
			)
		}

		return flex.Layout(
			gtx,
			items...,
		)
	}
}
