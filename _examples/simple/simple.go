package main

import (
	"github.com/malivvan/mate/view"
)

func main() {
	box := view.NewBox()
	box.SetBorder(true)
	box.SetTitle("Hello, world!")
	app := view.NewApplication()
	app.SetRoot(box, true)
	app.EnableMouse(true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
