package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/view"
)

// End shows the final slide.
func End(nextSlide func()) (title string, info string, content view.Primitive) {
	textView := view.NewTextView()
	textView.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://code.rocketnine.space/tslocum/view"
	fmt.Fprint(textView, url)
	return "End", "", Center(len(url), 1, textView)
}
