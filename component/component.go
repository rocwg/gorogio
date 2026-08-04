package component

import (
	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/style"
)

// Component
//
// Component 是可复用的交互式 UI 单元。
//
// Component 不负责页面生命周期。
// 是否处理输入事件，由具体组件自行决定。
type Component interface {

	// Build 构建组件 UI。
	Build(
		th *style.Theme,
	) element.Element
}
