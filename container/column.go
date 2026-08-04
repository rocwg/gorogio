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

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		flex := layout.Flex{
			Axis:      layout.Vertical,
			Alignment: options.Alignment,
		}

		items := make(
			[]layout.FlexChild,
			0,
			len(children)*2,
		)

		for index, child := range children {

			// 第二个元素开始增加间距
			if index > 0 &&
				options.Spacing > 0 {

				items = append(
					items,

					layout.Rigid(
						func(
							gtx layout.Context,
						) layout.Dimensions {

							return layout.Spacer{
								Height: options.Spacing,
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
