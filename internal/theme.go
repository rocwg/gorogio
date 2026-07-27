// Package internal/theme.go
// 目标：封装 gorogio 实验场的主题、颜色、字体
// 说明：
//   - 统一 material.Theme
//   - 定义主要颜色（Primary / Accent / Background）
//   - 封装字体和通用样式
//   - 可扩展工具函数，例如 FillBackground
//
// TODO: 根据实验场需求完善颜色、字体和样式
package internal

import (
	"image/color"

	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/widget/material"
)

// Theme 封装 Gio 的 material theme 和一些通用颜色
type Theme struct {
	Material *material.Theme // Material 主题对象
	Primary  color.NRGBA     // 主色
	Accent   color.NRGBA     // 强调色
}

// NewTheme 初始化 Theme 并返回默认主题
func NewTheme() *Theme {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	return &Theme{
		Material: th,                                         //
		Primary:  color.NRGBA{R: 33, G: 150, B: 243, A: 255}, // 蓝色
		Accent:   color.NRGBA{R: 255, G: 193, B: 7, A: 255},  // 黄色
	}
}
