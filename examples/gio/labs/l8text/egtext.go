package main

import (
	"image/color"

	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

/*
这是最 Gio、最干净的写法：直接返回值（推荐）

然后在 STATE 区：
	btn := new(widget.Clickable) //必须保存
	th := internal.NewTheme()

在 FRAME 区：
	EgButtonText(th, btn).Layout(gtx)  // 不需要保存
*/

// EgLabelText widget 文本标签
func egLabelText(th *material.Theme, lText string) material.LabelStyle {
	// 构造
	lt := material.H1(th, lText)

	// 更改
	lt.Alignment = text.Middle
	lt.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255}

	return lt
}

func egButtonText(th *material.Theme, btn *widget.Clickable) material.ButtonStyle {
	return material.Button(th, btn, "Click me")
}
