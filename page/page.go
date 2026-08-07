package page

import (
	"gioui.org/layout"

	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/style"
)

// Page 是 Application 管理的页面生命周期对象。
//
// 页面生命周期协议。
//
// 负责:
//  1. Update 处理输入事件、修改状态。
//  2. Element 构建 UI Tree。
//
// Application 不关心具体页面实现。
type Page interface {

	// Update 每一帧调用。
	//
	// 负责:
	// - Button Click
	// - Input Event
	// - State Change
	Update(
		gtx layout.Context,
	)

	// Element 返回 UI Element Tree。
	//
	// 一个页面的最终呈现
	Element(
		th *style.Theme,
	) element.Element
}
