// Demo code for the InputField primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/view"
)

func main() {
	app := view.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	inputField := view.NewInputField()
	inputField.SetLabel("Enter a number: ")
	inputField.SetPlaceholder("E.g. 1234")
	inputField.SetFieldWidth(10)
	inputField.SetAcceptanceFunc(view.InputFieldInteger)
	inputField.SetDoneFunc(func(key tcell.Key) {
		app.Stop()
	})

	app.SetRoot(inputField, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
