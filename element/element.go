package element

import (
	"gioui.org/layout"
)

// Element
//
// goro-ui 的核心 UI 元素协议。
//
// 一个 UI 元素，只需要具备：
//
// 1. 接收布局上下文
// 2. 完成自身布局
// 3. 返回占用尺寸
//
// 类似：
//
// Compose: @Composable
// SwiftUI: View
// Flutter: Widget
//
// 但是这里保持 Gio 的简单模型。
//
//
// 第一阶段：
// 直接复用 Gio Widget。
//
// 这样可以保持零成本封装。
//
// 后续如果需要多 backend，
// 再引入 adapter。
//
//type Element func(
//	layout.Context,
//) layout.Dimensions
//

type Element = layout.Widget
