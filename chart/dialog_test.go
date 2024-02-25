package chart_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/chart"
	"github.com/malivvan/mate/view"
	"testing"
)

func TestDialog(t *testing.T) {
	app := view.NewApplication()
	headerBox := view.NewBox()
	headerBox.SetBorder(true)
	msgDialog := chart.NewMessageDialog(chart.InfoDialog)
	screen := tcell.NewSimulationScreen("UTF-8")

	if err := screen.Init(); err != nil {
		panic(err)
	}

	go func() {
		appLayout := view.NewFlex()
		appLayout.SetDirection(view.FlexRow)
		appLayout.AddItem(headerBox, 0, 1, true)
		appLayout.AddItem(msgDialog, 0, 1, true)

		app.SetScreen(screen)
		app.SetRoot(appLayout, true)
		err := app.Run()
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
	}()

	tests := []struct {
		msgType int
		bgColor tcell.Color
	}{
		{msgType: chart.InfoDialog, bgColor: tcell.ColorSteelBlue},
		{msgType: chart.ErrorDailog, bgColor: tcell.ColorOrangeRed},
	}

	for _, test := range tests {
		msgDialog.SetType(test.msgType)
		app.Draw()
		if msgDialog.GetBackgroundColor() != test.bgColor {
			t.Errorf("Expected background color %v, but got %v", test.bgColor, msgDialog.GetBackgroundColor())
		}
	}

	app.SetFocus(headerBox)
	app.Draw()
	if msgDialog.HasFocus() {
		t.Errorf("Expected msgDialog to not have focus, but it does")
	}

	app.SetFocus(msgDialog)
	app.Draw()
	if !msgDialog.HasFocus() {
		t.Errorf("Expected msgDialog to have focus, but it does not")
	}

	app.Stop()
}
