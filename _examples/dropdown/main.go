package main

import "github.com/malivvan/mate/view"

func main() {
	app := view.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	dropdown := view.NewDropDown()
	dropdown.SetLabel("Select an option (hit Enter): ")
	dropdown.SetOptions(nil,
		view.NewDropDownOption("First"),
		view.NewDropDownOption("Second"),
		view.NewDropDownOption("Third"),
		view.NewDropDownOption("Fourth"),
		view.NewDropDownOption("Fifth"))

	app.SetRoot(dropdown, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
