package container

import (
	"gioui.org/layout"

	"github.com/rocwg/gorogio/element"
)

// Stack
//
// 层叠布局。
//
// 对应:
//
// Compose: Box
// SwiftUI: ZStack
// Flutter: Stack
//
// 用于: 背景、浮层、Badge、Floating Button
//

func Stack(
	children ...element.Element,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		items := make(
			[]layout.StackChild,
			0,
			len(children),
		)

		for _, child := range children {
			items = append(
				items,
				layout.Stacked(child),
			)
		}

		return layout.Stack{}.Layout(
			gtx,
			items...,
		)
	}
}
