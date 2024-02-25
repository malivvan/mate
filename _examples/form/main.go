package main

import (
	"github.com/malivvan/mate/view"
)

func main() {
	app := view.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	form := view.NewForm()
	form.AddDropDownSimple("Title", 0, nil, "Mr.", "Ms.", "Mrs.", "Dr.", "Prof.")
	form.AddInputField("First name", "", 20, nil, nil)
	form.AddInputField("Last name", "", 20, nil, nil)
	addressField := view.NewInputField()
	addressField.SetLabel("Address")
	addressField.SetFieldWidth(30)
	addressField.SetFieldNote("Your complete address")
	form.AddFormItem(addressField)
	form.AddPasswordField("Password", "", 10, '*', nil)
	form.AddCheckBox("", "Age 18+", false, nil)
	form.AddButton("Save", nil)
	form.AddButton("Quit", func() {
		app.Stop()
	})
	form.SetBorder(true)
	form.SetTitle("Enter some data")
	form.SetTitleAlign(view.AlignLeft)

	app.SetRoot(form, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
