package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

// Component 模式 C 的原子渲染单元
// 它是纯粹的声明式渲染视图，只负责在指定的 Context 和 Theme 下绘制自己
type Component interface {
	Layout(gtx layout.Context, th *material.Theme) layout.Dimensions
}
