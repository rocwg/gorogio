package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

//Theme 应该属于 Application，而不是 UI。
//Page

type Application struct {
	Theme *material.Theme
	Hello *HelloPage
}

func NewApplication() *Application {
	//创建主题
	theme := material.NewTheme()
	//状态
	state := &CounterState{}
	//
	page := NewHelloPage(state)

	return &Application{
		Theme: theme,
		Hello: page,
	}
}

func (a *Application) Draw(
	gtx layout.Context,
) {
	//
	a.Hello.Layout(gtx, a.Theme)
}
