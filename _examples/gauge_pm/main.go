package main

import (
	"github.com/malivvan/mate/chart"
	"time"

	"github.com/malivvan/mate/view"
)

func main() {
	app := view.NewApplication()
	gauge := chart.NewPercentageModeGauge()
	gauge.SetTitle("percentage mode gauge")
	gauge.SetRect(10, 4, 50, 3)
	gauge.SetBorder(true)

	value := 0
	gauge.SetMaxValue(50)

	update := func() {
		tick := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-tick.C:
				if value > gauge.GetMaxValue() {
					value = 0
				} else {
					value = value + 1
				}
				gauge.SetValue(value)
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
