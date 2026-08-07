package app

import (
	"gioui.org/layout"

	"github.com/rocwg/gorogio/style"
)

type Application struct {
	Theme  *style.Theme
	Screen Screen
}

func New(
	page Screen,
) *Application {

	return &Application{
		Theme:  style.NewTheme(),
		Screen: page,
	}
}

// Update
//
// Application 生命周期第一步。
func (a *Application) Update(
	gtx layout.Context,
) {
	a.Screen.Update(gtx)
}

// Draw
//
// Application 生命周期第二步。
func (a *Application) Draw(
	gtx layout.Context,
) {
	//Render
	root := a.Screen.Element(a.Theme)
	root(gtx) //最关键
}
