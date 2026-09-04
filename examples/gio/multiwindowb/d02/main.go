package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// 1. 中央窗口管理器
type WindowManager struct {
	mu          sync.Mutex
	activeCount int32 // 当前活跃的窗口数量
}

func NewWindowManager() *WindowManager {
	return &WindowManager{}
}

// OpenWindow 动态打开一个新窗口
func (wm *WindowManager) OpenWindow(title string) {
	// 增加活跃窗口计数
	atomic.AddInt32(&wm.activeCount, 1)

	// 每个窗口在一个全新的 Goroutine 中运行
	go func() {
		defer func() {
			// 窗口关闭时减少计数，若无活跃窗口则完全退出程序
			if atomic.AddInt32(&wm.activeCount, -1) == 0 {
				fmt.Println("所有窗口已关闭，程序退出。")
				os.Exit(0)
			}
		}()

		w := new(app.Window)
		w.Option(
			app.Title(title),
			app.Size(unit.Dp(450), unit.Dp(350)),
		)

		// 必须每个窗口初始化自己的 Theme 和 Ops
		th := material.NewTheme()
		var ops op.Ops
		var openBtn widget.Clickable
		childIndex := 1

		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				fmt.Printf("窗口 [%s] 销毁\n", title)
				return

			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				// 响应按钮点击：动态弹出子窗口
				if openBtn.Clicked(gtx) {
					subTitle := fmt.Sprintf("%s - 子弹窗 %d", title, childIndex)
					childIndex++
					wm.OpenWindow(subTitle)
				}

				// 渲染界面
				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis:      layout.Vertical,
						Alignment: layout.Middle,
					}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.H6(th, fmt.Sprintf("当前窗口: %s", title)).Layout(gtx)
						}),
						layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &openBtn, "点击弹出新窗口")
							return btn.Layout(gtx)
						}),
					)
				})

				e.Frame(gtx.Ops)
			}
		}
	}()
}

func main() {
	wm := NewWindowManager()

	go func() {
		// 启动最初的主窗口
		wm.OpenWindow("主控窗口")
	}()

	app.Main()
}
