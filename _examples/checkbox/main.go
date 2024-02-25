package main

import (
	"github.com/malivvan/mate/view"
)

func main() {
	app := view.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	checkbox := view.NewCheckBox()
	checkbox.SetLabel("Hit Enter to check box: ")

	app.SetRoot(checkbox, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
