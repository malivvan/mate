package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/editor"
	"github.com/rivo/tview"
	"log"
	"os"
	"path/filepath"
)

func saveBuffer(b *editor.Buffer, path string) error {
	return os.WriteFile(path, []byte(b.String()), 0600)
}

func main() {

	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		workdir, _ := os.Getwd()
		path = filepath.Join(workdir, "_examples/editor/editor.go")
	}
	content, _ := os.ReadFile(path)

	var colorscheme editor.Colorscheme
	if monokai := editor.Assets.FindFile(editor.RTColorscheme, "monokai"); monokai != nil {
		if data, err := monokai.Data(); err == nil {
			colorscheme = editor.ParseColorscheme(string(data))
		}
	}

	app := tview.NewApplication()
	buffer := editor.NewBufferFromString(string(content), path)
	root := editor.NewView(buffer)
	root.SetRuntimeFiles(editor.Assets)
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
