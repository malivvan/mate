package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/editor"
	"github.com/malivvan/mate/view"
	"log"
	"os"
)

func saveBuffer(b *editor.Buffer, path string) error {
	return os.WriteFile(path, []byte(b.String()), 0600)
}

func main() {

	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path = "/home/malivvan/Videos/mate/_examples/editor/main.go"
	}
	content, _ := os.ReadFile(path)

	var colorscheme editor.Colorscheme
	if monokai := editor.ColorschemeAssets().Get("monokai"); monokai != nil {
		colorscheme = editor.ParseColorscheme(string(monokai.Data))
	}

	app := view.NewApplication()
	buffer := editor.NewBufferFromString(string(content), path)
	root := editor.NewView(buffer)
	root.SetColorscheme(colorscheme)
	root.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlS:
			saveBuffer(buffer, path)
			return nil
		case tcell.KeyCtrlQ:
			app.Stop()
			return nil
		}
		return event
	})
	app.SetRoot(root, true)

	if err := app.Run(); err != nil {
		log.Fatalf("%v", err)
	}
}
