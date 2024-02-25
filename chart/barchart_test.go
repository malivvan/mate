package chart_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/chart"
	"github.com/malivvan/mate/view"
	"testing"
)

func TestBarChart(t *testing.T) {
	app := view.NewApplication()
	headerBox := view.NewBox()
	headerBox.SetBorder(true)
	barchart := chart.NewBarChart()
	screen := tcell.NewSimulationScreen("UTF-8")

	if err := screen.Init(); err != nil {
		panic(err)
	}

	go func() {
		appLayout := view.NewFlex()
		appLayout.SetDirection(view.FlexRow)
		appLayout.AddItem(headerBox, 1, 0, true)
		appLayout.AddItem(barchart, 50, 0, true)

		app.SetScreen(screen)
		app.SetRoot(appLayout, true)
		err := app.Run()
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
