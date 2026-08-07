package component

import (
	"gioui.org/layout"

	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/style"
)

// Component 是 UI Tree 中可复用节点。
// Screen 也是一种特殊的 Composite Component。
// 但是由于承担 Application 生命周期入口，
// 独立定义 Screen 协议。

// Component
//
// gorogio UI 组件生命周期协议。
//
// 一个 Component:
//
// 1. 更新内部交互状态
// 2. 生成 UI 单元
type Component interface {

	// Update component internal state
	//
	// 每帧调用。
	//
	// 负责：
	// - click
	// - input
	// - gesture
	Update(
		gtx layout.Context,
	)

	// Element Build UI element
	//
	// 生成 UI。
	Element(
		th *style.Theme,
	) element.Element
}
