package main

// To Run me:
// FYNE_THEME=light go run .

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewLabel("Hello Fyne!"))

	w.ShowAndRun()
}
