package container

import (
	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/rocwg/gorogio/element"
)

// flexLayout 内部 Flex 引擎。
//
// 负责:
//
// 1. Element 转换为 Gio FlexChild
// 2. spacing 注入
// 3. Flex 执行
//
// 外部不直接看到 Gio layout.Flex。
//

func flexLayout(
	axis layout.Axis,
	options Options,
	children []element.Element,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

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
					spacingChild(axis, options.Spacing),
				)
			}

			items = append(
				items,
				layout.Rigid(child),
			)
		}

		return layout.Flex{
			Axis:      axis,
			Alignment: options.Alignment,
		}.Layout(
			gtx,
			items...,
		)
	}
}

// spacingChild
//
// 根据方向创建间距。
func spacingChild(
	axis layout.Axis,
	value unit.Dp,
) layout.FlexChild {

	return layout.Rigid(
		func(
			gtx layout.Context,
		) layout.Dimensions {

			if axis == layout.Vertical {

				return layout.Spacer{
					Height: value,
				}.Layout(gtx)
			}

			return layout.Spacer{
				Width: value,
			}.Layout(gtx)

		},
	)
}
