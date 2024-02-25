package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/view"

	"github.com/malivvan/mate/image"
)

// generateModal makes a centered object
func generateModal(p view.Primitive, width, height int) view.Primitive {
	content := view.NewFlex()
	content.SetDirection(view.FlexRow)
	content.AddItem(nil, 0, 1, false)
	content.AddItem(p, height, 1, true)
	content.AddItem(nil, 0, 1, false)

	m := view.NewFlex()
	m.AddItem(nil, 0, 1, false)
	m.AddItem(content, width, 1, true)
	m.AddItem(nil, 0, 1, false)
	return m
}

func main() {
	// Create the application.
	a := view.NewApplication()

	// Create our starfield GIF.
	bg, err := image.GifFromImagePath("starfield.gif")
	if err != nil {
		panic(fmt.Errorf("Unable to load gif: %v", err))
	}
	go image.Animate(a)

	// Create our Hello World text.
	txt := view.NewTextView()
	txt.SetText("Hello, World")
	txt.SetTextAlign(view.AlignCenter)
	txt.SetDoneFunc(func(e tcell.Key) {
		a.Stop()
	})
	txt.SetBorder(true)

	// Create a layered page view with a modal
	panels := view.NewPanels()
	panels.AddPanel("bg", bg, true, true)
	panels.AddPanel("txt", generateModal(txt, 24, 3), true, true)

	a.SetRoot(panels, true)

	// Start the application.
	if err := a.Run(); err != nil {
		panic(err)
	}
}
