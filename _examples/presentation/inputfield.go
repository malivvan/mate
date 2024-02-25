package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/view"
)

const inputField = `[green]package[white] main

[green]import[white] (
    [red]"strconv"[white]

    [red]"github.com/gdamore/tcell/v2"[white]
    [red]"code.rocketnine.space/tslocum/view"[white]
)

[green]func[white] [yellow]main[white]() {
    input := view.[yellow]NewInputField[white]().
        [yellow]SetLabel[white]([red]"Enter a number: "[white]).
        [yellow]SetAcceptanceFunc[white](
            view.InputFieldInteger,
        ).[yellow]SetDoneFunc[white]([yellow]func[white](key tcell.Key) {
            text := input.[yellow]GetText[white]()
            n, _ := strconv.[yellow]Atoi[white](text)
            [blue]// We have a number.[white]
        })
    view.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](input, true).
        [yellow]Run[white]()
}`

// InputField demonstrates the InputField.
func InputField(nextSlide func()) (title string, info string, content view.Primitive) {
	input := view.NewInputField()
	input.SetLabel("Enter a number: ")
	input.SetAcceptanceFunc(view.InputFieldInteger)
	input.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	return "InputField", "", Code(input, 30, 1, inputField)
}
