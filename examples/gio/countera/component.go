package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

// Component 模式 C 的原子渲染单元
type Component interface {
	Layout(gtx layout.Context, th *material.Theme) layout.Dimensions
}

// WindowNode 窗口树的抽象节点
type WindowNode struct {
	Title    string
	Width    int
	Height   int
	Content  Component
	Children []*WindowNode // 子节点（模式 B：主从关系）
}

func NewWindowNode(title string, w, h int, content Component) *WindowNode {
	return &WindowNode{
		Title:   title,
		Width:   w,
		Height:  h,
		Content: content,
	}
}

func (n *WindowNode) AddChild(child *WindowNode) {
	n.Children = append(n.Children, child)
}
