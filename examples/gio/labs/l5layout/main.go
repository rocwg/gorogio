package main

import (
	"log"
	"os"

	"gioui.org/app"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("Test Window"))

		//err := draw0(window)
		//err := draw1(window)
		//err := draw2(window)
		//err := draw3(window)
		//err := draw4(window)
		//err := draw5(window)
		err := draw6(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
