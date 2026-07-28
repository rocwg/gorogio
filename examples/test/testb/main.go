package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/rocwg/gorogio/style"
	"github.com/rocwg/gorogio/view"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("My App"))
		if err := run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	// 1. 初始化主题与状态
	th := style.NewTheme()
	submitBtn := &view.PrimaryButton{Text: "提交数据"}

	var ops op.Ops

	for {
		e := w.Event()
		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// 2. 响应业务逻辑
			if submitBtn.Clicked(gtx) {
				log.Println("业务点击：数据提交中...")
			}

			// 3. 页面渲染
			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// 调用封装好的组件进行布局
				return submitBtn.Layout(gtx, th)
			})

			e.Frame(gtx.Ops)
		}
	}
}
