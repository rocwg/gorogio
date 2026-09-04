package main

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Option func(*Counter)

type Counter struct {
	Value  int
	decBtn widget.Clickable
	incBtn widget.Clickable

	step     int
	min      int
	max      int
	btnColor color.NRGBA
}

func WithStep(step int) Option {
	return func(c *Counter) {
		if step > 0 {
			c.step = step
		}
	}
}

func WithRange(min, max int) Option {
	return func(c *Counter) {
		c.min = min
		c.max = max
	}
}

func WithButtonColor(col color.NRGBA) Option {
	return func(c *Counter) {
		c.btnColor = col
	}
}

func NewCounter(initialValue int, opts ...Option) *Counter {
	c := &Counter{
		Value:    initialValue,
		step:     1,
		min:      0,
		max:      100,
		btnColor: color.NRGBA{R: 0, G: 122, B: 255, A: 255},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Counter) update(gtx layout.Context) {
	if c.decBtn.Clicked(gtx) {
		c.Value -= c.step
		if c.Value < c.min {
			c.Value = c.min
		}
	}
	if c.incBtn.Clicked(gtx) {
		c.Value += c.step
		if c.Value > c.max {
			c.Value = c.max
		}
	}
}

// Layout 实现 Component 接口
func (c *Counter) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	c.update(gtx)

	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &c.decBtn, "-")
			btn.Background = c.btnColor
			if c.Value <= c.min {
				btn.Background = color.NRGBA{R: 200, G: 200, B: 200, A: 255}
			}
			return btn.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Left:  unit.Dp(16),
				Right: unit.Dp(16),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.H6(th, fmt.Sprintf("%d", c.Value)).Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &c.incBtn, "+")
			btn.Background = c.btnColor
			if c.Value >= c.max {
				btn.Background = color.NRGBA{R: 200, G: 200, B: 200, A: 255}
			}
			return btn.Layout(gtx)
		}),
	)
}
