package main

import (
	"github.com/malivvan/mate/chart"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/view"
)

func main() {
	app := view.NewApplication()
	gauge := chart.NewActivityModeGauge()
	gauge.SetTitle("activity mode gauge")
	gauge.SetPgBgColor(tcell.ColorOrange)
	gauge.SetRect(10, 4, 50, 3)
	gauge.SetBorder(true)

	update := func() {
		tick := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-tick.C:
				gauge.Pulse()
				app.Draw()
			}
		}
	}
	go update()

	app.SetRoot(gauge, true)
	app.EnableMouse(true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
