package main

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ButtonVariant uint8

const (
	// ButtonPrimary 使用 Gio Material 主题的默认按钮颜色。
	ButtonPrimary ButtonVariant = iota
	// ButtonSecondary 用于弱化的次要操作。
	ButtonSecondary
)

// ButtonProps 描述按钮的视觉属性，而不保存按钮本身的交互状态。
type ButtonProps struct {
	Text    string
	Variant ButtonVariant
}

// Button 是一个薄的、Compose 风格的按钮封装。
// Clickable 由调用者长期持有；否则每一帧创建新值会丢失点击状态。
func Button(gtx layout.Context, th *material.Theme, click *widget.Clickable, props ButtonProps) layout.Dimensions {
	button := material.Button(th, click, props.Text)
	if props.Variant == ButtonSecondary {
		button.Background = color.NRGBA{R: 92, G: 92, B: 92, A: 255}
	}
	return button.Layout(gtx)
}
