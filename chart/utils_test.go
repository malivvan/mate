package chart_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/chart"
	"github.com/malivvan/mate/view"
	"testing"
)

func TestGetColorName(t *testing.T) {
	tests := []struct {
		color     tcell.Color
		colorName string
	}{
		{color: tcell.ColorWhite, colorName: "white"},
		{color: tcell.ColorBlack, colorName: "black"},
		{color: tcell.NewRGBColor(0, 1, 2), colorName: ""},
	}

	for _, test := range tests {
		if got := chart.GetColorName(test.color); got != test.colorName {
			t.Errorf("Expected color name to be %v, but got %v", test.colorName, got)
		}
	}
}

func TestGetMessageWidth(t *testing.T) {
	tests := []struct {
		msg   string
		width int
	}{
		{msg: "test", width: 4},
		{msg: "test01\ntest001", width: 7},
		{msg: "", width: 0},
	}

	for _, test := range tests {
		if got := chart.GetMessageWidth(test.msg); got != test.width {
			t.Errorf("Expected width to be %v, but got %v", test.width, got)
		}
	}
}

func TestGetMaxFloat64From2dSlice(t *testing.T) {
	tests := []struct {
		have  [][]float64
		wants float64
	}{
		{have: [][]float64{}, wants: 0},
		{have: [][]float64{
			{5, -1, 0, -10, 12},
			{15, -11, 0, -110, 22},
		}, wants: 22},
		{have: [][]float64{
			{-5, -1, -2, -10, -12},
			{-15, -11, -1, -110, -22},
		}, wants: -1},
	}

	for _, test := range tests {
		if got := chart.GetMaxFloat64From2dSlice(test.have); got != test.wants {
			t.Errorf("Expected max value to be %v, but got %v", test.wants, got)
		}
	}
}

func TestGetMaxFloat64FromSlice(t *testing.T) {
	tests := []struct {
		have  []float64
		wants float64
	}{
		{have: []float64{}, wants: 0},
		{have: []float64{5, -1, 0, -10, 12}, wants: 12},
		{have: []float64{-10, -20, -9, -1}, wants: -1},
	}

	for _, test := range tests {
		if got := chart.GetMaxFloat64FromSlice(test.have); got != test.wants {
			t.Errorf("Expected max value to be %v, but got %v", test.wants, got)
		}
	}
}

func TestAbsInt(t *testing.T) {
	tests := []struct {
		have  int
		wants int
	}{
		{have: 2, wants: 2},
		{have: -2, wants: 2},
		{have: 0, wants: 0},
	}

	for _, test := range tests {
		if got := chart.AbsInt(test.have); got != test.wants {
			t.Errorf("Expected absolute value to be %v, but got %v", test.wants, got)
		}
	}
}

func TestDrawLine(t *testing.T) {
	screen := tcell.NewSimulationScreen("UTF-8")
	screenWidth := 70
	screenHeight := 30
	lineStartX := 0
	lineStartY := 0
	lineLength := 20
	screen.SetSize(screenWidth, screenHeight)
	screen.Init()
	screen.Clear()

	// draw and test horizontal line
	chart.DrawLine(screen, lineStartX, lineStartY, lineLength, 0, tcell.ColorDefault)
	screen.Show()

	cellRune, _, _, _ := screen.GetContent(lineStartX, lineStartY)
	if cellRune != view.BoxDrawingsLightTripleDashHorizontal {
		t.Errorf("Expected rune to be %v, but got %v", view.BoxDrawingsLightTripleDashHorizontal, cellRune)
	}

	// draw and test vertical line
	screen.Clear()
	chart.DrawLine(screen, lineStartX, lineStartY, lineLength, 1, tcell.ColorDefault)
	screen.Show()

	cellRune, _, _, _ = screen.GetContent(lineStartX, lineStartY)
	if cellRune != view.BoxDrawingsLightTripleDashVertical {
		t.Errorf("Expected rune to be %v, but got %v", view.BoxDrawingsLightTripleDashVertical, cellRune)
	}
}
