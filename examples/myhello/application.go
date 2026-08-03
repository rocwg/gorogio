package main

import (
	"gioui.org/layout"

	"github.com/rocwg/gorogio/style"
)

type Application struct {
	Theme *style.Theme

	Page *HelloPage
}

func NewApplication() *Application {
	theme := style.NewTheme()
	state := &CounterState{}
	page := NewHelloPage(state)

	return &Application{
		Theme: theme,
		Page:  page,
	}
}

func (a *Application) Draw(
	gtx layout.Context,
) {
	ui := a.Page.Element(a.Theme)
	ui(gtx) //:最关键
}
