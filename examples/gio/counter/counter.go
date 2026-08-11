package main

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// Option 定义配置函数类型
type Option func(*Counter)

// -----------------------------------------------------------------------------
// 1. Struct 定义：包含运行期状态（State）与外观配置（Config）
// -----------------------------------------------------------------------------

type Counter struct {
	// 运行期状态（State）
	Value  int              // 当前数值
	decBtn widget.Clickable // 减号按钮交互状态
	incBtn widget.Clickable // 加号按钮交互状态

	// 配置项（Configuration）
	step     int         // 每次加减的步长
	min      int         // 最小值限制
	max      int         // 最大值限制
	btnColor color.NRGBA // 按钮背景色
}

// -----------------------------------------------------------------------------
// 2. Functional Options 配置函数
// -----------------------------------------------------------------------------

// WithStep 设置加减步长
func WithStep(step int) Option {
	return func(c *Counter) {
		if step > 0 {
			c.step = step
		}
	}
}

// WithRange 设置数值上下限
func WithRange(min, max int) Option {
	return func(c *Counter) {
		c.min = min
		c.max = max
	}
}

// WithButtonColor 设置按钮主题颜色
func WithButtonColor(col color.NRGBA) Option {
	return func(c *Counter) {
		c.btnColor = col
	}
}

// -----------------------------------------------------------------------------
// 3. 构造函数 (NewCounter)
// -----------------------------------------------------------------------------

func NewCounter(initialValue int, opts ...Option) *Counter {
	// 默认配置
	c := &Counter{
		Value:    initialValue,
		step:     1,
		min:      0,
		max:      100,
		btnColor: color.NRGBA{R: 0, G: 122, B: 255, A: 255}, // 默认 iOS 蓝
	}

	// 应用自定义配置项
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// -----------------------------------------------------------------------------
// 4. Update 逻辑：处理事件更新状态
// -----------------------------------------------------------------------------

func (c *Counter) update(gtx layout.Context) {
	// 处理减法按钮点击
	if c.decBtn.Clicked(gtx) {
		c.Value -= c.step
		if c.Value < c.min {
			c.Value = c.min
		}
	}

	// 处理加法按钮点击
	if c.incBtn.Clicked(gtx) {
		c.Value += c.step
		if c.Value > c.max {
			c.Value = c.max
		}
	}
}

// -----------------------------------------------------------------------------
// 5. Layout 方法：实现标准的 Gio 布局接口
// -----------------------------------------------------------------------------

// Layout 渲染控件并返回其占用的真实尺寸 (Dimensions)
// 注意：必须使用指针接收者 (c *Counter)
func (c *Counter) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	// 1. 每帧优先处理输入与状态更新
	c.update(gtx)

	// 2. 组合底层布局 (使用 Flex 水平排列：[ - ]  [ 数值 ]  [ + ])
	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
	}.Layout(gtx,
		// 减号按钮
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &c.decBtn, "-")
			btn.Background = c.btnColor
			// 禁用判定
			if c.Value <= c.min {
				btn.Background = color.NRGBA{R: 200, G: 200, B: 200, A: 255}
			}
			return btn.Layout(gtx)
		}),

		// 中间数值显示（带左右外边距）
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Left:  unit.Dp(16),
				Right: unit.Dp(16),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				label := material.H6(th, fmt.Sprintf("%d", c.Value))
				return label.Layout(gtx)
			})
		}),

		// 加号按钮
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &c.incBtn, "+")
			btn.Background = c.btnColor
			// 禁用判定
			if c.Value >= c.max {
				btn.Background = color.NRGBA{R: 200, G: 200, B: 200, A: 255}
			}
			return btn.Layout(gtx)
		}),
	)
}
