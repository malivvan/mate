package chart_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/chart"
	"github.com/malivvan/mate/view"
	"testing"
)

func TestSparkline(t *testing.T) {
	app := view.NewApplication()
	headerBox := view.NewBox()
	headerBox.SetBorder(true)
	sparkline := chart.NewSparkline()
	screen := tcell.NewSimulationScreen("UTF-8")

	if err := screen.Init(); err != nil {
		panic(err)
	}

	go func() {
		appLayout := view.NewFlex()
		appLayout.SetDirection(view.FlexRow)
		appLayout.AddItem(headerBox, 1, 0, true)
		appLayout.AddItem(sparkline, 50, 0, true)

		app.SetScreen(screen)
		app.SetRoot(appLayout, true)
		err := app.Run()
		if err != nil {
			panic(err)
		}
	}()

	app.SetFocus(headerBox)
	app.Draw()
	if sparkline.HasFocus() {
		t.Errorf("Expected sparkline to not have focus, but it does")
	}

	app.SetFocus(sparkline)
	app.Draw()
	if !sparkline.HasFocus() {
		t.Errorf("Expected sparkline to have focus, but it does not")
	}

	x, y, width, height := sparkline.GetRect()
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

	tests := []struct {
		title string
		color tcell.Color
	}{
		{title: "test01", color: tcell.ColorDarkOrange},
		{title: "abc123", color: tcell.ColorBlue},
	}

	for _, test := range tests {
		sparkline.SetDataTitle(test.title)
		sparkline.SetDataTitleColor(test.color)
		app.Draw()

		for x := 0; x < len(test.title); x++ {
			rune, _, style, _ := screen.GetContent(x, 1)
			fg, _, _ := style.Decompose()

			if fg != test.color {
				t.Errorf("Expected color to be %v, but got %v", test.color, fg)
			}
			if string(rune) != string(test.title[x]) {
				t.Errorf("Expected rune to be %v, but got %v", string(test.title[x]), string(rune))
			}
		}
	}

	app.Stop()
}
