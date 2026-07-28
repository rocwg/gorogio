package style

import (
	"image/color"

	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// Theme 是你应用的全局主题配置，充当设计规范字典
type Theme struct {
	Material *material.Theme

	// 自定义设计系统扩展（可根据需要增减）
	Spacing struct {
		Small  unit.Dp
		Medium unit.Dp
		Large  unit.Dp
	}
	Colors struct {
		Background color.NRGBA
		CardBg     color.NRGBA
		Primary    color.NRGBA
		Text       color.NRGBA
	}
}

// NewTheme 初始化你自己的主题
func NewTheme() *Theme {
	th := &Theme{}
	// 基于 Gio 的默认 Material 主题
	th.Material = material.NewTheme()
	th.Material.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	// 定义你的业务/应用专属尺寸与配色
	th.Spacing.Small = unit.Dp(4)
	th.Spacing.Medium = unit.Dp(8)
	th.Spacing.Large = unit.Dp(16)

	th.Colors.Background = color.NRGBA{R: 245, G: 245, B: 245, A: 255}
	th.Colors.CardBg = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	th.Colors.Primary = color.NRGBA{R: 30, G: 144, B: 255, A: 255} // 例如 DodgerBlue
	th.Colors.Text = color.NRGBA{R: 33, G: 33, B: 33, A: 255}

	return th
}
