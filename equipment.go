package main

import (
	//"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/data/binding"
	//"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (w *win) EquipUI(app fyne.App) {
	w.window = app.NewWindow("Equipment")
	empty := widget.NewLabel("text")
	list := widget.NewList(
		func() int {
			return len(Equipdata)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("EqipmentList")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(Equipdata[i])
		})
	w.window.SetContent(container.New(layout.NewGridLayout(2), list, empty))
	//w.window= SetContent(grid)
	w.window.Show()
}
