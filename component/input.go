package component

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/style"
)

// Input
//
// 一个简单输入组件。
//
// 拥有：
// 1. Gio Editor 状态
// 2. label
//
// 对外：
// 1. Update 生命周期
// 2. Element 渲染
// 3. Value 获取内容
type Input struct {
	// Gio 输入状态
	editor widget.Editor

	// 显示提示
	label string
}

// NewInput 创建输入框
func NewInput(
	label string,
) *Input {
	return &Input{
		label:  label,
		editor: widget.Editor{},
	}
}

// Update 事件处理
func (i *Input) Update(
	gtx layout.Context,
) {
	// 当前版本无需额外事件处理。
	// Gio Editor 在 Element 渲染阶段处理输入事件。
}

// Value 获取输入内容
func (i *Input) Value() string {
	return i.editor.Text()
}

func (i *Input) SetValue(value string) {
	i.editor.SetText(value)
}

// Element
//
// 负责生成 UI 元素。
func (i *Input) Element(
	th *style.Theme,
) element.Element {

	return func(gtx layout.Context) layout.Dimensions {
		return material.Editor(
			th.Material,
			&i.editor,
			i.label,
		).Layout(gtx)
	}
}
