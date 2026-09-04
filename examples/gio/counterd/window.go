package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// WindowNode 抽象 OS 窗口节点（Layer 1 主窗口 或 Layer 2 平行子窗口）
type WindowNode struct {
	Title    string
	Width    int
	Height   int
	Content  Component     // 承载的模式 C 业务组件
	Children []*WindowNode // 维护的 Layer 2 子窗口（可选）
}

func NewWindowNode(title string, w, h int, content Component) *WindowNode {
	return &WindowNode{
		Title:   title,
		Width:   w,
		Height:  h,
		Content: content,
	}
}

// SpawnChild 派生平行子窗口 (Layer 2)
func SpawnChild(parent *WindowNode, child *WindowNode) {
	if parent != nil {
		parent.Children = append(parent.Children, child)
	}
	go func() {
		_ = ServeWindowNode(child)
	}()
}

// ServeWindowNode 驱动 Gio 窗口事件循环
func ServeWindowNode(node *WindowNode) error {
	w := new(app.Window)
	w.Option(
		app.Title(node.Title),
		app.Size(unit.Dp(float32(node.Width)), unit.Dp(float32(node.Height))),
	)

	theme := material.NewTheme()
	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// 渲染挂载的 Component 内容
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if node.Content != nil {
					return node.Content.Layout(gtx, theme)
				}
				return layout.Dimensions{}
			})

			e.Frame(gtx.Ops)
		}
	}
}
