package main

import (
	"fmt"

	"gioui.org/layout"

	"github.com/rocwg/gorogio/app"
	"github.com/rocwg/gorogio/component"
	"github.com/rocwg/gorogio/container"
	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/modifier"
	"github.com/rocwg/gorogio/primitive"
	"github.com/rocwg/gorogio/style"
)

var _ app.Screen = (*SettingsScreen)(nil)

type SettingsScreen struct {
	DarkMode *component.Checkbox

	Save *component.Button
}

func NewSettingsScreen() *SettingsScreen {
	screen := &SettingsScreen{
		DarkMode: component.NewCheckbox("Dark Mode"),
	}
	screen.Save =
		component.NewButton("Save").
			OnClick(
				func() {
					fmt.Println(
						"DarkMode:",
						screen.DarkMode.Checked(),
					)
				},
			)
	return screen
}

func (s *SettingsScreen) Update(
	gtx layout.Context,
) {
	s.DarkMode.Update(gtx)
	s.Save.Update(gtx)
}

func (s *SettingsScreen) Element(
	th *style.Theme,
) element.Element {

	return modifier.Padding(

		40,

		container.Column(
			container.Options{
				Spacing:   th.Spacing.Large,
				Alignment: layout.Middle,
			},
			primitive.H3(th, "Settings"),

			s.DarkMode.Element(th),
			s.Save.Element(th),
		),
	)
}
