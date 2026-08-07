package primitive

import (
	"gioui.org/layout"
	"gioui.org/widget"

	"github.com/rocwg/gorogio/element"
	"github.com/rocwg/gorogio/style"
)

func Icon(
	th *style.Theme,
	data []byte,
) element.Element {

	return func(
		gtx layout.Context,
	) layout.Dimensions {

		icon, _ := widget.NewIcon(data)
		return icon.Layout(gtx, th.Material.Bg)
	}
}
