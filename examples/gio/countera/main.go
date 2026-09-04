package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	// 1. 构建业务组件（模式 C）
	customCounter := NewCounter(10,
		WithStep(5),
		WithRange(0, 50),
		WithButtonColor(color.NRGBA{R: 46, G: 139, B: 87, A: 255}),
	)

	// 2. 构建窗口树（Root Node）
	rootNode := NewWindowNode("Custom Widget Demo", 400, 300, customCounter)

	// 3. 挂载运行逻辑
	go func() {
		if err := ServeWindowNode(rootNode); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}

// ServeWindowNode 统一的树节点到 OS Window 的映射桥梁
func ServeWindowNode(node *WindowNode) error {
	w := new(app.Window)
	w.Option(app.Title(node.Title), app.Size(unit.Dp(float32(node.Width)), unit.Dp(float32(node.Height))))

	theme := material.NewTheme()
	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// 居中渲染节点的 Component 内容
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return node.Content.Layout(gtx, theme)
			})

			e.Frame(gtx.Ops)
		}
	}
}
