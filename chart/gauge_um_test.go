package chart_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/chart"
	"github.com/rivo/tview"
	"testing"
)

func TestGaugeUm(t *testing.T) {
	app := tview.NewApplication()
	headerBox := tview.NewBox().SetBorder(true)
	gaugeUm := chart.NewUtilModeGauge()
	screen := tcell.NewSimulationScreen("UTF-8")

	if err := screen.Init(); err != nil {
		panic(err)
	}

	go func() {
		appLayout := tview.NewFlex().SetDirection(tview.FlexRow)
		appLayout.AddItem(headerBox, 1, 0, true)
		appLayout.AddItem(gaugeUm, 50, 0, true)
		err := app.SetScreen(screen).SetRoot(appLayout, true).Run()
		if err != nil {
			panic(err)
		}
	}()

	app.SetFocus(headerBox)
	app.Draw()
	if gaugeUm.HasFocus() {
		t.Errorf("Expected gaugeUm to not have focus, but it does")
	}

	app.SetFocus(gaugeUm)
	app.Draw()
	// gauge will not get focus
	if gaugeUm.HasFocus() {
		t.Errorf("Expected gaugeUm to not have focus, but it does")
	}

	x, y, width, height := gaugeUm.GetRect()
	if x != 0 {
		t.Errorf("Expected x to be 0, but got %v", x)
	}
	if y != 1 {
		t.Errorf("Expected y to be 1, but got %v", y)
	}
	if width != 80 {
		t.Errorf("Expected width to be 80, but got %v", width)
	}
	if height != 50 {
		t.Errorf("Expected height to be 50, but got %v", height)
	}

	app.Stop()
}