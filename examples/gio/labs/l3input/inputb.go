package main

import (
	"image"
	"sync"
	"time"

	"gioui.org/app"
	"gioui.org/op"
)

func draw1(window *app.Window) error {
	// 按钮状态（包含互斥锁）
	var button struct {
		lock   sync.Mutex
		offset int
	}

	updateOffset := func(v int) {
		button.lock.Lock()
		defer button.lock.Unlock()
		button.offset = v
	}
	readOffset := func() int {
		button.lock.Lock()
		defer button.lock.Unlock()
		return button.offset
	}

	// 后台 Goroutine：每秒触发一次位置更新
	go func() {
		changes := time.NewTicker(time.Second)
		defer changes.Stop()
		for t := range changes.C {
			updateOffset(int((t.Second() % 3) * 100))
			window.Invalidate() // 核心：唤醒 GUI 渲染循环
		}
	}()

	ops := new(op.Ops)
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ops.Reset()

			// 根据状态偏移按钮。
			op.Offset(image.Pt(readOffset(), 0)).Add(ops)

			// 处理按钮输入并绘制。
			doButton(ops, e.Source)

			// 更新显示。
			e.Frame(ops)
		}
	}
}
