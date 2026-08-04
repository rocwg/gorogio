package app

import (
	"gioui.org/layout"

	"github.com/rocwg/gorogio/page"
	"github.com/rocwg/gorogio/style"
)

type Application struct {
	Theme *style.Theme
	Page  page.Page
}

func New(
	page page.Page,
) *Application {

	return &Application{
		Theme: style.NewTheme(),
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

	// 2. Render
	root := a.Page.Element(a.Theme)

	root(gtx) //:最关键
}
