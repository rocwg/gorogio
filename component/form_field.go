package component

import (
	"gioui.org/layout"
	"github.com/rocwg/gorogio/container"
	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/primitive"
	"github.com/rocwg/gorogio/style"
)

type FormField struct {
	label string
	input *Input
}

func NewFormField(label string, hint string) *FormField {
	return &FormField{
		label: label,
		input: NewInput(hint),
	}
}

func (f *FormField) Value() string {
	return f.input.Value()
}

func (f *FormField) SetValue(
	value string,
) {
	f.input.SetValue(value)
}

func (f *FormField) Update(
	gtx layout.Context,
) {
	f.input.Update(gtx)
}

func (f *FormField) Element(
	th *style.Theme,
) element.Element {

	return container.Row(
		container.Options{
			Spacing: th.Spacing.Large,
		},
		primitive.Body(th, f.label),
		f.input.Element(th),
	)
}
