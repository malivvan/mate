package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/view"
)

func main() {
	app := view.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	box := view.NewBox()
	box.SetBackgroundColor(tcell.ColorBlue.TrueColor())

	frame := view.NewFrame(box)
	frame.SetBorders(2, 2, 2, 2, 4, 4)
	frame.AddText("Header left", true, view.AlignLeft, tcell.ColorWhite.TrueColor())
	frame.AddText("Header middle", true, view.AlignCenter, tcell.ColorWhite.TrueColor())
	frame.AddText("Header right", true, view.AlignRight, tcell.ColorWhite.TrueColor())
	frame.AddText("Header second middle", true, view.AlignCenter, tcell.ColorRed.TrueColor())
	frame.AddText("Footer middle", false, view.AlignCenter, tcell.ColorGreen.TrueColor())
	frame.AddText("Footer second middle", false, view.AlignCenter, tcell.ColorGreen.TrueColor())

	app.SetRoot(frame, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
