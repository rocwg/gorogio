package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func draw0(window *app.Window) error {
	// 1. 初始化全局 Material 主题
	th := material.NewTheme()
	// 2. 配置字体整形器（Shaper），这里加载了标准的 Go 字体库集合（Go Font Collection）
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		// 窗口已被关闭。
		case app.DestroyEvent:
			return e.Err

		// 绘制窗口状态的请求。
		case app.FrameEvent:
			// 为新帧重置 layout.Context。
			gtx := app.NewContext(&ops, e)

			// 根据 e.Queue 中的事件，把状态绘制进 ops。
			themedApplication(gtx, th)

			// 更新显示。
			e.Frame(gtx.Ops)
		}
	}
}

// 1. 状态对象（必须在帧间持久化）
var isChecked widget.Bool

func themedApplication(gtx layout.Context, th *material.Theme) layout.Dimensions {
	var checkboxLabel string

	// 2. 根据用户输入事件更新状态（如点击切换 True/False）
	isChecked.Update(gtx)
	if isChecked.Value {
		checkboxLabel = "checked"
	} else {
		checkboxLabel = "not-checked"
	}

	// 3. 声明式构建 UI 布局
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		// 使用主题创建 H3 标题组件样式并布局
		layout.Rigid(material.H3(th, "Hello, World!").Layout),
		// 使用主题创建 CheckBox 复选框组件样式并布局
		layout.Rigid(material.CheckBox(th, &isChecked, checkboxLabel).Layout),
	)
}
