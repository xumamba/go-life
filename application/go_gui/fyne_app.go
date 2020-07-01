package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("Hello")
	myWin.SetContent(widget.NewLabel("Go GUI"))
	myWin.Resize(fyne.NewSize(200,200))
	myWin.ShowAndRun()
}
