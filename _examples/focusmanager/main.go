package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/keybind"
	"github.com/malivvan/mate/view"
)

func wrap(f func()) func(ev *tcell.EventKey) *tcell.EventKey {
	return func(ev *tcell.EventKey) *tcell.EventKey {
		f()
		return nil
	}
}

func main() {
	app := view.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	input1 := view.NewInputField()
	input1.SetLabel("InputField 1")

	input2 := view.NewInputField()
	input2.SetLabel("InputField 2")

	input3 := view.NewInputField()
	input3.SetLabel("InputField 3")

	input4 := view.NewInputField()
	input4.SetLabel("InputField 4")

	grid := view.NewGrid()
	grid.SetBorder(true)
	grid.SetTitle(" Press Tab to advance focus ")
	grid.AddItem(input1, 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(input2, 0, 1, 1, 1, 0, 0, false)
	grid.AddItem(input3, 1, 1, 1, 1, 0, 0, false)
	grid.AddItem(input4, 1, 0, 1, 1, 0, 0, false)

	focusManager := view.NewFocusManager(app.SetFocus)
	focusManager.SetWrapAround(true)
	focusManager.Add(input1, input2, input3, input4)

	inputHandler := keybind.NewConfiguration()
	for _, key := range view.Keys.MovePreviousField {
		err := inputHandler.Set(key, wrap(focusManager.FocusPrevious))
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, key := range view.Keys.MoveNextField {
		err := inputHandler.Set(key, wrap(focusManager.FocusNext))
		if err != nil {
			log.Fatal(err)
		}
	}
	app.SetInputCapture(inputHandler.Capture)

	app.SetRoot(grid, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
