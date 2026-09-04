package main

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// Counter 纯粹的原子组件，不包含任何窗口派生业务
type Counter struct {
	value    int
	step     int
	min, max int

	incBtn widget.Clickable
	decBtn widget.Clickable

	btnColor color.NRGBA
}

func NewCounter(initValue int) *Counter {
	return &Counter{
		value:    initValue,
		step:     1,
		min:      0,
		max:      100,
		btnColor: color.NRGBA{R: 46, G: 139, B: 87, A: 255},
	}
}

func (c *Counter) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	if c.incBtn.Clicked(gtx) && c.value+c.step <= c.max {
		c.value += c.step
	}
	if c.decBtn.Clicked(gtx) && c.value-c.step >= c.min {
		c.value -= c.step
	}

	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(material.Button(th, &c.decBtn, "-").Layout),
		layout.Rigid(layout.Spacer{Width: 10}.Layout),
		layout.Rigid(material.Body1(th, fmt.Sprintf("%d", c.value)).Layout),
		layout.Rigid(layout.Spacer{Width: 10}.Layout),
		layout.Rigid(material.Button(th, &c.incBtn, "+").Layout),
	)
}
