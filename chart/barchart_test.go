package chart_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/chart"
	"github.com/rivo/tview"
	"testing"
)

func TestBarChart(t *testing.T) {
	app := tview.NewApplication()
	headerBox := tview.NewBox().SetBorder(true)
	barchart := chart.NewBarChart()
	screen := tcell.NewSimulationScreen("UTF-8")

	if err := screen.Init(); err != nil {
		panic(err)
	}

	go func() {
		appLayout := tview.NewFlex().SetDirection(tview.FlexRow)
		appLayout.AddItem(headerBox, 1, 0, true)
		appLayout.AddItem(barchart, 50, 0, true)
		err := app.SetScreen(screen).SetRoot(appLayout, true).Run()
		if err != nil {
			panic(err)
		}
	}()

	app.SetFocus(headerBox)
	app.Draw()
	if barchart.HasFocus() {
		t.Errorf("Expected barchart to not have focus, but it does")
	}

	app.SetFocus(barchart)
	app.Draw()
	if !barchart.HasFocus() {
		t.Errorf("Expected barchart to have focus, but it does not")
	}

	app.SetFocus(headerBox)
	app.Draw()
	if barchart.HasFocus() {
		t.Errorf("Expected barchart to not have focus, but it does")
	}

	app.SetFocus(barchart)
	app.Draw()
	if !barchart.HasFocus() {
		t.Errorf("Expected barchart to have focus, but it does not")
	}

	app.Stop()
}
