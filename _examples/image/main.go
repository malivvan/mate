package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/malivvan/mate/image"
)

// generateModal makes a centered object
func generateModal(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

func main() {
	// Create the application.
	a := tview.NewApplication()

	// Create our starfield GIF.
	bg, err := image.GifFromImagePath("starfield.gif")
	if err != nil {
		panic(fmt.Errorf("Unable to load gif: %v", err))
	}
	go image.Animate(a)

	// Create our Hello World text.
	txt := tview.NewTextView()
	txt.SetText("Hello, World").
		SetTextAlign(tview.AlignCenter).
		SetDoneFunc(func(e tcell.Key) {
			a.Stop()
		}).
		SetBorder(true)

	// Create a layered page view with a modal
	a.SetRoot(tview.NewPages().
		AddPage("bg", bg, true, true).
		AddPage("txt", generateModal(txt, 24, 3), true, true), true)

	// Start the application.
	if err := a.Run(); err != nil {
		panic(err)
	}
}
