package component

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/style"
)

// Checkbox
//
// 一个简单复选框组件。
//
// 自己拥有：
// checked 状态
type Checkbox struct {
	click widget.Bool

	label string
}

// NewCheckbox
func NewCheckbox(
	label string,
) *Checkbox {

	return &Checkbox{
		label: label,
	}
}

// Update
//
// 每一帧处理点击事件。
func (c *Checkbox) Update(
	gtx layout.Context,
) {
	c.click.Update(gtx)
}

// Checked
//
// 获取当前状态。
func (c *Checkbox) Checked() bool {
	return c.click.Value
}

// Toggle
//
// 修改状态。
func (c *Checkbox) Toggle() {
	c.click.Value = !c.click.Value
}

// Element
func (c *Checkbox) Element(
	th *style.Theme,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		return material.CheckBox(
			th.Material,
			&c.click,
			c.label,
		).Layout(gtx)
	}
}
