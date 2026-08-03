package style

import (
	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// Theme
//
// goro-ui 视觉环境。
//
// 第一阶段:
//
// 内部复用 Gio material.Theme。
//
// 后续:
//
// 可以加入:
//
// Color
// Typography
// Spacing
//

type Theme struct {
	Material *material.Theme

	// 扩展：自定义尺寸
	Spacing struct {
		Small  unit.Dp
		Medium unit.Dp
		Large  unit.Dp
	}
}

// NewTheme
//
// 创建默认主题
//

func NewTheme() *Theme {

	// 初始化你自己的主题
	th := &Theme{}

	// 基于 Gio 的默认 Material 主题
	th.Material = material.NewTheme()
	th.Material.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	// 自定义尺寸
	th.Spacing.Small = unit.Dp(4)
	th.Spacing.Medium = unit.Dp(8)
	th.Spacing.Large = unit.Dp(16)

	return th
}
