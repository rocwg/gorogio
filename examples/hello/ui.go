package main

import (
	"gioui.org/layout"
)

var application = NewApplication()

func DrawUI(gtx layout.Context) {
	application.Draw(gtx)
}
