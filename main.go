package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("你好，世界！")
	label := widget.NewLabel("Hello, world! 你好，世界!")
	w.SetContent(label)
	w.ShowAndRun()
}
