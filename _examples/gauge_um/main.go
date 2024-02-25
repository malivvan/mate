package main

import (
	"github.com/malivvan/mate/chart"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/view"
)

func main() {
	app := view.NewApplication()
	gauge := chart.NewUtilModeGauge()
	gauge.SetLabel("cpu usage:")
	gauge.SetLabelColor(tcell.ColorLightSkyBlue)
	gauge.SetRect(10, 4, 50, 3)
	gauge.SetWarnPercentage(65)
	gauge.SetCritPercentage(80)
	gauge.SetBorder(true)

	update := func() {
		tick := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-tick.C:
				randNum := float64(rand.Float64() * 100)
				gauge.SetValue(randNum)
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
