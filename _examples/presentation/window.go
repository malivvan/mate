package main

import (
	"github.com/malivvan/mate/view"
)

const loremIpsumText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

// Window returns the window page.
func Window(nextSlide func()) (title string, info string, content view.Primitive) {
	wm := view.NewWindowManager()

	list := view.NewList()
	list.ShowSecondaryText(false)
	list.AddItem(view.NewListItem("Item #1"))
	list.AddItem(view.NewListItem("Item #2"))
	list.AddItem(view.NewListItem("Item #3"))
	list.AddItem(view.NewListItem("Item #4"))
	list.AddItem(view.NewListItem("Item #5"))
	list.AddItem(view.NewListItem("Item #6"))
	list.AddItem(view.NewListItem("Item #7"))

	loremIpsum := view.NewTextView()
	loremIpsum.SetText(loremIpsumText)

	w1 := view.NewWindow(list)
	w1.SetRect(2, 2, 10, 7)

	w2 := view.NewWindow(loremIpsum)
	w2.SetRect(7, 4, 12, 12)

	w1.SetTitle("List")
	w2.SetTitle("Lorem Ipsum")

	wm.Add(w1, w2)

	return "Window", windowInfo, wm
}
