package element

import (
	"gioui.org/layout"
)

// Element
//
// 核心 UI 元素协议。
//
// 类似：
//
// Compose: @Composable
// SwiftUI: View
// Flutter: Widget
//
// 但是这里保持 Gio 的简单模型。
//
// 第一阶段：直接复用 Gio Widget。
//
// 这样可以保持零成本封装。
type Element = layout.Widget
