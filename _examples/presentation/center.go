package main

import "github.com/malivvan/mate/view"

// Center returns a new primitive which shows the provided primitive in its
// center, given the provided primitive's size.
func Center(width, height int, p view.Primitive) view.Primitive {
	subFlex := view.NewFlex()
	subFlex.SetDirection(view.FlexRow)
	subFlex.AddItem(view.NewBox(), 0, 1, false)
	subFlex.AddItem(p, height, 1, true)
	subFlex.AddItem(view.NewBox(), 0, 1, false)

	flex := view.NewFlex()
	flex.AddItem(view.NewBox(), 0, 1, false)
	flex.AddItem(subFlex, width, 1, true)
	flex.AddItem(view.NewBox(), 0, 1, false)

	return flex
}
