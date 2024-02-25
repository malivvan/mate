package chart_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/chart"
	"github.com/rivo/tview"
	"testing"
)

func TestSpinner(t *testing.T) {
	app := tview.NewApplication()
	headerBox := tview.NewBox().SetBorder(true)
	spinner := chart.NewSpinner()
	screen := tcell.NewSimulationScreen("UTF-8")

	if err := screen.Init(); err != nil {
		panic(err)
	}

	go func() {
		appLayout := tview.NewFlex().SetDirection(tview.FlexRow)
		appLayout.AddItem(headerBox, 1, 0, true)
		appLayout.AddItem(spinner, 50, 0, true)
		err := app.SetScreen(screen).SetRoot(appLayout, true).Run()
		if err != nil {
			panic(err)
		}
	}()

	app.SetFocus(headerBox)
	app.Draw()
	if spinner.HasFocus() {
		t.Errorf("Expected spinner to not have focus, but it does")
	}

	app.SetFocus(spinner)
	app.Draw()
	if !spinner.HasFocus() {
		t.Errorf("Expected spinner to have focus, but it does not")
	}

	x, y, width, height := spinner.GetRect()
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

	spinner.SetStyle(chart.SpinnerGrowHorizontal)
	spinner.Reset()
	app.Draw()

	r, _, _, _ := screen.GetContent(0, 1)
	if r != '▉' {
		t.Errorf("Expected rune to be '▉', but got %v", r)
	}

	spinner.Pulse()
	app.Draw()
	r, _, _, _ = screen.GetContent(0, 1)
	if r != '▊' {
		t.Errorf("Expected rune to be '▊', but got %v", r)
	}

	customStyle := []rune{'\u2705', '\u274C'}
	spinner.SetCustomStyle(customStyle)
	spinner.Reset()

	app.Draw()
	r, _, _, _ = screen.GetContent(0, 1)
	if r != customStyle[0] {
		t.Errorf("Expected rune to be %v, but got %v", customStyle[0], r)
	}

	spinner.Pulse()
	app.Draw()
	r, _, _, _ = screen.GetContent(0, 1)
	if r != customStyle[1] {
		t.Errorf("Expected rune to be %v, but got %v", customStyle[1], r)
	}

	spinner.Pulse()
	app.Draw()
	r, _, _, _ = screen.GetContent(0, 1)
	if r != customStyle[0] {
		t.Errorf("Expected rune to be %v, but got %v", customStyle[0], r)
	}

	app.Stop()
}
