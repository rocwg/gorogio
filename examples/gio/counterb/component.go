package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// Component 模式 C 的原子渲染单元
type Component interface {
	Layout(gtx layout.Context, th *material.Theme) layout.Dimensions
}

// WindowNode 抽象窗口树节点
type WindowNode struct {
	Title    string
	Width    int
	Height   int
	Content  Component
	Children []*WindowNode
}

func NewWindowNode(title string, w, h int, content Component) *WindowNode {
	return &WindowNode{
		Title:   title,
		Width:   w,
		Height:  h,
		Content: content,
	}
}

// SpawnChild 模式 B：在独立的 Goroutine 中动态弹出 OS 子窗口
func SpawnChild(parent *WindowNode, child *WindowNode) {
	parent.Children = append(parent.Children, child)
	go func() {
		_ = ServeWindowNode(child)
	}()
}

// ServeWindowNode 统一的 OS 窗口渲染管线
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

			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return node.Content.Layout(gtx, theme)
			})

			e.Frame(gtx.Ops)
		}
	}
}
