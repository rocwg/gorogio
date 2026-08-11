package main

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// main_ui.go (可独立为单独的 Go 文件)

type MainUI struct {
	theme  *material.Theme
	button widget.Clickable // 保存点击状态
	count  int
}

func NewMainUI(th *material.Theme) *MainUI {
	return &MainUI{theme: th}
}

// 遵循 Gio 标准的 Layout 签名

func (ui *MainUI) Layout(gtx layout.Context) layout.Dimensions {
	// 1. 处理交互逻辑
	if ui.button.Clicked(gtx) {
		ui.count++
	}

	// 2. 处理布局渲染
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			title := material.H1(ui.theme, fmt.Sprintf("Clicked: %d", ui.count))
			return title.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(ui.theme, &ui.button, "Click Me")
			return btn.Layout(gtx)
		}),
	)
}
