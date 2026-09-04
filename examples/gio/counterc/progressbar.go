package main

import (
	"image"
	"image/color"
	"time"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type ProgressBar struct {
	startTime time.Time
	duration  time.Duration
}

func NewProgressBar(duration time.Duration) *ProgressBar {
	return &ProgressBar{
		startTime: time.Now(),
		duration:  duration,
	}
}

func (p *ProgressBar) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	now := gtx.Now
	elapsed := now.Sub(p.startTime)
	progress := elapsed.Seconds() / p.duration.Seconds()

	if progress < 1 {
		gtx.Execute(op.InvalidateCmd{})
	} else {
		progress = 1
	}

	width := float32(gtx.Dp(unit.Dp(200))) * float32(progress)
	height := gtx.Dp(unit.Dp(20))

	defer clip.Rect{Max: image.Pt(int(width), height)}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color.NRGBA{R: 0x2E, G: 0x8B, B: 0x57, A: 0xFF}}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return layout.Dimensions{Size: image.Pt(gtx.Dp(unit.Dp(200)), height)}
}
