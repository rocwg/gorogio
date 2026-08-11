package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
)

// main 只负责创建窗口并交给 run 处理 UI 生命周期。
func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("Egg timer"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		if err := run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
