package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/chart"
	"github.com/malivvan/mate/view"
)

func main() {
	app := view.NewApplication()
	barGraph := chart.NewBarChart()
	barGraph.SetRect(4, 2, 50, 20)
	barGraph.SetBorder(true)
	barGraph.SetTitle("System Resource Usage")

	barGraph.AddBar("cpu", 80, tcell.ColorBlue)
	barGraph.AddBar("mem", 20, tcell.ColorRed)
	barGraph.AddBar("swap", 40, tcell.ColorGreen)
	barGraph.AddBar("disk", 40, tcell.ColorOrange)
	barGraph.SetMaxValue(100)
	barGraph.SetAxesColor(tcell.ColorAntiqueWhite)
	barGraph.SetAxesLabelColor(tcell.ColorAntiqueWhite)

	app.SetRoot(barGraph, true)
	app.EnableMouse(true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
