package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget"
	"github.com/rocwg/gorogio/internal"
	"github.com/rocwg/gorogio/labs/l8text"
)

/*
Gio is a library for implementing immediate mode user interfaces.
This approach can be implemented in multiple ways, however the overarching similarity is that the program:

 1. Listens for events such as mouse or keyboard input.
 2. Updates its internal state based on the event.
 3. Runs code that lays out and redraws the user interface state.

A minimal immediate mode command-line UI in pseudocode:

Gio 是一个用于实现「立即模式（immediate mode）」用户界面的库。
这种方法可以有多种具体实现方式，但它们在整体思想上有一个共同点：程序会执行以下流程：

	1. 监听事件，例如鼠标输入或键盘输入。
	2. 根据事件更新程序的内部状态（internal state）。
	3. 运行一段代码，根据当前的用户界面状态进行布局计算，并重新绘制界面。

下面是一个最简的、立即模式的命令行 UI 的伪代码示例：
*/

func main() {
	// 第 0 幕：线程与窗口的出生
	// goroutine → 你的 UI 逻辑跑在另一个 goroutine 里
	go func() {
		// 创建窗口
		window := new(app.Window)
		window.Option(app.Title("Test Window"))

		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	// 运行 → 把当前线程交给平台事件循环（Windows / macOS / Wayland）
	app.Main()
}

/*
当前 state 清单，恰到好处：
 - ops     → 指令缓冲区容器；可以复用，Reset 清空指令，不需要 new
 - th      → Theme / Shaper state
 - btn     → 交互 state
 - clicked → 业务 state

其他的都属于 frame 临时量，每帧构造：
 - gtx   → 每帧新建，绑定当前事件
 - text
 - ButtonStyle / LabelStyle  → 每帧构造，轻量，绑定当前 state

完全符合 IMUI 原则。
*/

func run(window *app.Window) error {

	// ========== STATE 区 ==========
	var ops op.Ops // ops 是“指令缓冲区容器”，不是“绘制结果”

	th := internal.NewTheme()
	btn := new(widget.Clickable) // 交互 state（非常关键）
	clicked := false             // 业务 state

	// 进入“永动机”主循环
	for {
		//监听事件
		switch e := window.Event().(type) {
		case app.DestroyEvent: // 窗口关闭，世界结束。
			return e.Err
		case app.FrameEvent: // 系统请求你：现在请给我一帧画面。

			// 上一帧作废
			ops.Reset()

			// ========== FRAME 区 ==========
			gtx := app.NewContext(&ops, e) // 本帧上下文
			if btn.Clicked(gtx) {
				clicked = !clicked
			}
			l8text.EgButtonText(th.Material, btn).Layout(gtx)
			text := "Hello, Gio"
			if clicked {
				text = "Clicked!"
			}
			l8text.EgLabelText(th.Material, text).Layout(gtx) // UI = f(state)
			// ========== FRAME 区 ==========

			// 提交这一帧
			e.Frame(&ops)
		}
	}
}
