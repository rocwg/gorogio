package main

import (
	"sync"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// 1. 全局窗口管理器
type WindowManager struct {
	mu      sync.Mutex
	theme   *material.Theme
	windows map[*app.Window]struct{}
}

func NewWindowManager() *WindowManager {
	return &WindowManager{
		theme:   material.NewTheme(),
		windows: make(map[*app.Window]struct{}),
	}
}

// 打开一个新窗口
func (wm *WindowManager) OpenWindow(title string) {
	go func() {
		w := new(app.Window)
		w.Option(app.Title(title))

		wm.mu.Lock()
		wm.windows[w] = struct{}{}
		wm.mu.Unlock()

		defer func() {
			wm.mu.Lock()
			delete(wm.windows, w)
			wm.mu.Unlock()
		}()

		var ops op.Ops
		var openBtn widget.Clickable

		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				return
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				// 响应点击：弹出新窗口
				if openBtn.Clicked(gtx) {
					wm.OpenWindow("子窗口")
				}

				// 布局UI
				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.Button(wm.theme, &openBtn, "打开新窗口").Layout(gtx)
				})

				e.Frame(gtx.Ops)
			}
		}
	}()
}
