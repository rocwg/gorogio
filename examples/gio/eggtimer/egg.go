package main

import (
	"image"
	"image/color"
	"math"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

// EggProps 是鸡蛋组件的输入。组件本身不保存业务状态，
// 每一帧只根据传入的 Progress 生成当前的绘制操作。
type EggProps struct {
	Progress float32 // 取值范围为 0 到 1。
}

// Egg 绘制随进度改变颜色的鸡蛋曲线。
func Egg(gtx layout.Context, props EggProps) layout.Dimensions {
	var eggPath clip.Path
	op.Offset(image.Pt(gtx.Dp(200), gtx.Dp(150))).Add(gtx.Ops)
	eggPath.Begin(gtx.Ops)
	for deg := 0.0; deg <= 360; deg++ {
		rad := deg * math.Pi / 180
		cosT, sinT := math.Cos(rad), math.Sin(rad)
		const a, b, d = 110.0, 150.0, 20.0
		x := a * cosT
		y := -(math.Sqrt(b*b-d*d*cosT*cosT) + d*sinT) * sinT
		eggPath.LineTo(f32.Pt(float32(x), float32(y)))
	}
	eggPath.Close()

	eggArea := clip.Outline{Path: eggPath.End()}.Op()
	eggColor := color.NRGBA{
		R: 255,
		G: uint8(239 * (1 - props.Progress)),
		B: uint8(174 * (1 - props.Progress)),
		A: 255,
	}
	paint.FillShape(gtx.Ops, eggColor, eggArea)
	return layout.Dimensions{Size: image.Point{Y: 335}}
}
