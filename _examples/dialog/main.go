// Demo code for the bar chart primitive.
package main

import (
	"github.com/malivvan/mate/chart"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	dialog := chart.NewMessageDialog(chart.ErrorDailog)
	dialog.SetTitle("error dialog")
	dialog.SetMessage("This is first line of error\nThis is second line of the error message")
	dialog.SetDoneFunc(func() {
		app.Stop()
	})

	if err := app.SetRoot(dialog, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
