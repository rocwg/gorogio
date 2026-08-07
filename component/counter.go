package component

import (
	"strconv"

	"gioui.org/layout"

	"github.com/rocwg/gorogio/container"
	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/primitive"
	"github.com/rocwg/gorogio/style"
)

// component=有交互、有组合、有业务语义的 UI
// 1. Component Ownership：Page 不应该知道 Counter 内部如何工作。
// 2. Counter 应该自己拥有状态。

type Counter struct {
	count int

	increment *Button
	reset     *Button
}

func NewCounter() *Counter {

	counter := &Counter{}

	counter.increment =
		NewButton("+").
			OnClick(
				func() {
					counter.Increment()
				},
			)

	counter.reset =
		NewButton("Reset").
			OnClick(
				func() {
					counter.Reset()
				},
			)

	return counter
}

func (c *Counter) Update(
	gtx layout.Context,
) {
	c.increment.Update(gtx)
	c.reset.Update(gtx)
}

func (c *Counter) Value() int {
	return c.count
}

func (c *Counter) Element(
	th *style.Theme,
) element.Element {

	return container.Column(
		container.Options{
			Spacing:   th.Spacing.Large,
			Alignment: layout.Middle,
		},

		primitive.Body(
			th,
			"Count: "+strconv.Itoa(c.count),
		),

		container.Row(
			container.Options{
				Spacing: th.Spacing.Large,
			},
			c.increment.Element(th),
			c.reset.Element(th),
		),
	)
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Reset() {
	c.count = 0
}
