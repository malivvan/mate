package main

import (
	"github.com/malivvan/mate/view"
)

func main() {
	app := view.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	modal := view.NewModal()
	modal.SetText("Do you want to quit the application?")
	modal.AddButtons([]string{"Quit", "Cancel"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Quit" {
			app.Stop()
		}
	})

	app.SetRoot(modal, false)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
