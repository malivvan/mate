package main

import (
	"github.com/malivvan/mate/menu"
	"log"

	"github.com/rivo/tview"
)

func clickedMessageFn(msg string) func(*menu.MenuItem) {
	return func(*menu.MenuItem) { log.Printf("%v clicked\n", msg) }
}

func main() {
	app := tview.NewApplication()

	fileMenu := menu.NewMenuItem("File")
	fileMenu.AddItem(menu.NewMenuItem("New File").SetOnClick(clickedMessageFn("New File")))
	fileMenu.AddItem(menu.NewMenuItem("Open File").SetOnClick(clickedMessageFn("Open File")))

	saveSubForReal := menu.NewMenuItem("Save For Real").
		AddItem(menu.NewMenuItem("For really real").SetOnClick(clickedMessageFn("For really real"))).
		AddItem(menu.NewMenuItem("For really fake").SetOnClick(clickedMessageFn("For really fake")))
	saveSubForFake := menu.NewMenuItem("Save For Fake").SetOnClick(clickedMessageFn("Safe for fake"))

	fileMenu.AddItem(menu.NewMenuItem("Save File").
		// Add submenu items to save
		AddItem(saveSubForReal).
		AddItem(saveSubForFake).SetOnClick(clickedMessageFn("Save File")))

	fileMenu.AddItem(menu.NewMenuItem("Close File").SetOnClick(clickedMessageFn("Close File")))
	fileMenu.AddItem(menu.NewMenuItem("Exit").SetOnClick(func(*menu.MenuItem) { app.Stop() }))
	editMenu := menu.NewMenuItem("Edit")
	editMenu.AddItem(menu.NewMenuItem("Copy").SetOnClick(clickedMessageFn("Copy")))
	editMenu.AddItem(menu.NewMenuItem("Cut").SetOnClick(clickedMessageFn("Cut")))
	editMenu.AddItem(menu.NewMenuItem("Paste").SetOnClick(clickedMessageFn("Paste")))

	menuBar := menu.NewMenuBar().
		AddItem(fileMenu).
		AddItem(editMenu)

	menuBar.SetRect(0, 0, 100, 15)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(menuBar, 1, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Hello, world!"), 0, 1, true)

	app.EnableMouse(true).SetRoot(flex, true).SetFocus(flex).SetAfterDrawFunc(menuBar.AfterDraw())

	if err := app.Run(); err != nil {
		panic(err)
	}
}
