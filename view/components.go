package view

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/rocwg/gorogio/style"
)

// PrimaryButton 业务级主要按钮
type PrimaryButton struct {
	Clickable widget.Clickable
	Text      string
}

// Layout 负责组件的渲染与事件响应
// 注意：保持即时模式的 signature (gtx layout.Context) -> layout.Dimensions
func (b *PrimaryButton) Layout(gtx layout.Context, th *style.Theme) layout.Dimensions {
	// 使用 Theme 中定义的样式覆盖 Gio 默认样式
	btn := material.Button(th.Material, &b.Clickable, b.Text)
	btn.Background = th.Colors.Primary
	btn.Inset = layout.UniformInset(th.Spacing.Medium)

	return btn.Layout(gtx)
}

// Clicked 暴露一个简单的点击判断方法，避免上层代码去查 Clickable 细节
func (b *PrimaryButton) Clicked(gtx layout.Context) bool {
	return b.Clickable.Clicked(gtx)
}
