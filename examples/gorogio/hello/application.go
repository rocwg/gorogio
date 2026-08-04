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

// Update
//
// Application 生命周期第一步。
func (a *Application) Update(
	gtx layout.Context,
) {

	a.Page.Update(gtx)

}

// Draw
//
// Application 生命周期第二步。
func (a *Application) Draw(
	gtx layout.Context,
) {
	ui := a.Page.Element(a.Theme)
	ui(gtx) //:最关键
}
