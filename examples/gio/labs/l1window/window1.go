package main

import (
	"image"

	"gioui.org/app"
	"gioui.org/op"
)

func run1(window *app.Window) error {
	// 1. 在栈/堆上分配了一个 Ops 结构体
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:

			// 2. 把 ops 的指针 (&ops) 传给了 NewContext
			// app.NewContext 内部会自动执行 ops.Reset()
			gtx := app.NewContext(&ops, e)

			// ========== FRAME 区 ==========
			// 优先将 gtx 传递给 UI 布局或绘制函数
			colorBox(gtx.Ops, image.Point{X: 100, Y: 100}, red)
			// ========== FRAME 区 ==========

			// 提交绘制指令
			e.Frame(gtx.Ops) // 或者 e.Frame(&ops)，两者完全等价，但推荐写 gtx.Ops 语义更统一
		}
	}
}
