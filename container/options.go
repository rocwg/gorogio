package container

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

// Options Container 通用布局配置。
//
// 对应:
//
// Compose:
// Column(
//
//	verticalArrangement,
//	horizontalAlignment
//
// )
//
// SwiftUI:
// VStack(
//
//	alignment,
//	spacing
//
// )
//
// Flutter:
// MainAxisAlignment
// CrossAxisAlignment
type Options struct {

	// 子元素之间距离
	Spacing unit.Dp

	// Flex 交叉轴方向对齐
	//
	// Vertical: 控制水平位置
	// Horizontal: 控制垂直位置
	Alignment layout.Alignment
}
