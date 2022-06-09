package main

import (
	//"io"

	"strconv"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/data/binding"
	//"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (w *win) ItemUI(app fyne.App) {
	w.window = app.NewWindow("Item")
	empty := widget.NewLabel("text")

	data := make([]string, 1000)
	//i := 1
	for i := range data {
		data[i] = strconv.Itoa(i+1) + " Test Item "
	}

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("Template Object"), widget.NewIcon(theme.DocumentIcon()))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(data[id])
		},
	)
	w.window.SetContent(container.New(layout.NewGridLayout(2), list, empty))
	//w.window= SetContent(grid)
	w.window.Show()
}
