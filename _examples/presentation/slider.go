package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/view"
)

const sliderCode = `[green]package[white] main

[green]import[white] (
    [red]"fmt"[white]

    [red]"github.com/gdamore/tcell/v2"[white]
    [red]"code.rocketnine.space/tslocum/view"[white]
)

[green]func[white] [yellow]main[white]() {
    slider := view.[yellow]NewSlider[white]()
    slider.[yellow]SetLabel[white]([red]"Volume:   0%"[white])
    slider.[yellow][yellow]SetChangedFunc[white]([yellow]func[white](key tcell.Key) {
        label := fmt.[yellow]Sprintf[white]("Volume: %3d%%", value)
        slider.[yellow]SetLabel[white](label)
    })
    slider.[yellow][yellow]SetDoneFunc[white]([yellow]func[white](key tcell.Key) {
        [yellow]nextSlide[white]()
    })
    app := view.[yellow]NewApplication[white]()
    app.[yellow]SetRoot[white](slider, true)
    app.[yellow]Run[white]()
}`

// Slider demonstrates the Slider.
func Slider(nextSlide func()) (title string, info string, content view.Primitive) {
	slider := view.NewSlider()
	slider.SetLabel("Volume:   0%")
	slider.SetChangedFunc(func(value int) {
		slider.SetLabel(fmt.Sprintf("Volume: %3d%%", value))
	})
	slider.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	return "Slider", sliderInfo, Code(slider, 30, 1, sliderCode)
}
