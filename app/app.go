package app

import (
	"gioui.org/layout"

	"github.com/rocwg/gorogio/style"
)

type Application struct {
	Theme *style.Theme
	Scene Scene
}

func New(
	scene Scene,
) *Application {

	return &Application{
		Theme: style.NewTheme(),
		Scene: scene,
	}
}

// Update
//
// Application 生命周期第一步。
func (a *Application) Update(
	gtx layout.Context,
) {
	a.Scene.Update(gtx)
}

// Draw
//
// Application 生命周期第二步。
func (a *Application) Draw(
	gtx layout.Context,
) {
	//Render
	root := a.Scene.Element(a.Theme)
	root(gtx) //最关键
}
