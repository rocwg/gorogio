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
		err := draw1(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
