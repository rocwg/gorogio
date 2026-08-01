package main

import (
	"image/color"
	"strconv"

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

func NewHelloPage(state *CounterState) *HelloPage {
	return &HelloPage{
		State: state,
	}
}

//Layout：
//Column
//    Header
//    Spacer
//    Counter
//    Spacer
//    Actions(Row)

func (p *HelloPage) Layout(
	gtx layout.Context,
	theme *material.Theme,
) layout.Dimensions {

	// 处理事件
	for p.Increment.Clicked(gtx) {
		p.State.Increment()
	}

	for p.Reset.Clicked(gtx) {
		p.State.Reset()
	}

	//第一：Column（纵向布局）

	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Middle,
	}.Layout(

		gtx,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.layoutHeader(gtx, theme)
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Height: 20}.Layout(gtx)
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.layoutCounter(gtx, theme)
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Height: 20}.Layout(gtx)
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.layoutActions(gtx, theme)
		}),
	)

}

func (p *HelloPage) layoutHeader(
	gtx layout.Context,
	theme *material.Theme,
) layout.Dimensions {

	// Define a large label with an appropriate text:
	title := material.H3(theme, "Hello, Gio")

	// Change the color of the label.
	title.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255}

	// Change the position of the label.
	title.Alignment = text.Middle

	// Draw the label to the graphics context.
	return title.Layout(gtx)
}

func (p *HelloPage) layoutCounter(
	gtx layout.Context,
	theme *material.Theme,
) layout.Dimensions {

	txt := "Count : " + strconv.Itoa(p.State.Count)
	label := material.Body1(theme, txt)

	label.Alignment = text.Middle

	return label.Layout(gtx)
}

func (p *HelloPage) layoutActions(
	gtx layout.Context,
	theme *material.Theme,
) layout.Dimensions {

	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
	}.Layout(

		gtx,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Button(theme, &p.Increment, "+").Layout(gtx)
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: 16}.Layout(gtx)
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Button(theme, &p.Reset, "Reset").Layout(gtx)
		}),
	)
}
