package main

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// MainPage 业务编排层：组合原子组件，处理上下文业务
type MainPage struct {
	counter  *Counter
	spawnBtn widget.Clickable

	// 业务动作：向外暴露“请求创建子窗口”的事件
	OnRequestSpawnWindow func()
}

func NewMainPage() *MainPage {
	return &MainPage{
		counter: NewCounter(10),
	}
}

func (p *MainPage) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	// 处理点击事件
	if p.spawnBtn.Clicked(gtx) {
		if p.OnRequestSpawnWindow != nil {
			p.OnRequestSpawnWindow()
		}
	}

	// 编排布局：上方是计数器，下方是弹窗按钮
	return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.counter.Layout(gtx, th)
		}),
		layout.Rigid(layout.Spacer{Height: 20}.Layout),
		layout.Rigid(material.Button(th, &p.spawnBtn, "派生平行子窗口").Layout),
	)
}
