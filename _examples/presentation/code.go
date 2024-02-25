package main

import (
	"fmt"

	"github.com/malivvan/mate/view"
)

// The width of the code window.
const codeWidth = 56

// Code returns a primitive which displays the given primitive (with the given
// size) on the left side and its source code on the right side.
func Code(p view.Primitive, width, height int, code string) view.Primitive {
	// Set up code view.
	codeView := view.NewTextView()
	codeView.SetWrap(false)
	codeView.SetDynamicColors(true)
	codeView.SetPadding(1, 1, 2, 0)
	fmt.Fprint(codeView, code)

	f := view.NewFlex()
	f.AddItem(Center(width, height, p), 0, 1, true)
	f.AddItem(codeView, codeWidth, 1, false)
	return f
}
