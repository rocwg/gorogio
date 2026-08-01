package main

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

//其实：HelloComponent，已经不是一个 Component。
//而是：一个 Page。
//Application
//        │
//        ▼
//    HelloPage
//        │
//        ├── Title()
//        ├── Button()
//        └── Counter()

type HelloPage struct {
	Increment widget.Clickable
	Reset     widget.Clickable

	State *CounterState
}

func NewHelloComponent(
	state *CounterState,
) *HelloPage {

	return &HelloPage{
		State: state,
	}
}

func (c *HelloPage) Layout(
	gtx layout.Context,
	theme *material.Theme,
) layout.Dimensions {

	for c.Increment.Clicked(gtx) {
		c.State.Increment()
	}

	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
	}.Layout(

		gtx,

		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {

				return c.Title(
					gtx,
					theme,
				)

			},
		),

		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {

				return layout.Spacer{
					Height: 20,
				}.Layout(gtx)

			},
		),

		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {

				return c.ButtonView(
					gtx,
					theme,
				)

			},
		),

		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {

				return c.Counter(
					gtx,
					theme,
				)

			},
		),
	)

}

func (c *HelloPage) Title(
	gtx layout.Context,
	theme *material.Theme,
) layout.Dimensions {

	// Define a large label with an appropriate text:
	title := material.H1(theme, "Hello, Gio")

	// Change the color of the label.
	title.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255}

	// Change the position of the label.
	title.Alignment = text.Middle

	// Draw the label to the graphics context.
	return title.Layout(gtx)
}

func (c *HelloPage) ButtonView(
	gtx layout.Context,
	theme *material.Theme,
) layout.Dimensions {

	button := material.Button(theme, &c.Increment, "Click Me")
	return button.Layout(gtx)
}

func (c *HelloPage) Counter(
	gtx layout.Context,
	theme *material.Theme,
) layout.Dimensions {

	label := material.Body1(
		theme,
		"Count: "+itoa(
			c.State.Count,
		),
	)
	return label.Layout(gtx)
}

func itoa(v int) string {

	if v == 0 {
		return "0"
	}

	buf := ""
	for v > 0 {
		buf = string(rune('0'+v%10)) + buf
		v /= 10
	}
	return buf
}
