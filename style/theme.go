package style

import (
	"image/color"

	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type Spacing struct {
	Small  unit.Dp
	Medium unit.Dp
	Large  unit.Dp
}

// Theme Render Environment（渲染环境）。
type Theme struct {

	// 内部复用 Gio material.Theme
	Material *material.Theme

	// 扩展：自定义尺寸
	Spacing Spacing

	// Divider 默认颜色
	DividerColor color.NRGBA
}

// NewTheme 创建默认主题
func NewTheme() *Theme {

	// 初始化你自己的主题
	th := &Theme{}

	// 默认 Material 主题
	th.Material = material.NewTheme()

	// ?
	th.Material.Shaper =
		text.NewShaper(
			text.WithCollection(
				gofont.Collection(),
			),
		)

	// 自定义尺寸
	th.Spacing.Small = unit.Dp(4)
	th.Spacing.Medium = unit.Dp(8)
	th.Spacing.Large = unit.Dp(16)

	th.DividerColor =
		color.NRGBA{
			R: 200,
			G: 200,
			B: 200,
			A: 255,
		}
	return th
}
