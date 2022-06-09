package main

import (
	//"io"

	"io"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"

	//"fyne.io/fyne/v2/data/binding"
	//"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (w *win) ArmorUI(app fyne.App) {
	w.window = app.NewWindow("Weapons")
	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")

	test := make([]string, 3)
	Armordata := make([]string, 3)
	test[0] = "DummyHelmet"
	test[1] = "DummyChestpiece"
	test[2] = "DummyLegs"

	for i := range Armordata {
		Armordata[i] = strconv.Itoa(i+1) + " " + test[i]
	}

	list := widget.NewList(
		func() int {
			return len(Armordata)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("ArmorList")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(Armordata[i])
		})

	gbox := container.New(layout.NewGridLayout(3), list)

	list.OnSelected = func(id widget.ListItemID) {
		//id2 = id
		label.SetText(Equipdata[id])
		icon.SetResource(theme.DocumentIcon())

		buttons := w.funcbuttons(app, list) //assigns fyne.CanvasObject(HBOX) to variable buttons
		//assigns fyne.CanvasObject(GridWithColumns) to variable weakness
		//weaknesswidget := container.New(layout.NewGridWrapLayout(fyne.NewSize(600, 600)), weakness) //add additional widgets with Wrap to adjust TextSize
		//materials := w.materials(app)
		TableForgeMatsArmor := w.materialsforgeArmor(app)
		TableUpgradeMatsArmor := w.materialsupgradeArmor(app)

		gbox = container.New(layout.NewGridLayout(2), list, TableForgeMatsArmor, buttons, TableUpgradeMatsArmor) //display gbox
		w.window.SetContent(gbox)
		w.window.Show()

	}

	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
		gbox = container.New(layout.NewGridLayout(3), list) //remove additional widgets
		w.window.SetContent(gbox)                           //display gbox
		w.window.Show()
	}
	list.Select(125)

	w.window.SetContent(container.New(layout.NewGridLayout(2), list))
	//w.window= SetContent(grid)
	w.window.Show()
}

func (w *win) materialsforgeArmor(app fyne.App) fyne.CanvasObject {

	var data = [][]string{[]string{"Icon", "ItemName", "Quantity"},
		[]string{"", "Dummy Forge", "1x"}}

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("wide content"))
		},
		func(id widget.TableCellID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(data[id.Row][id.Col])
		})

	return table

}

func (w *win) materialsupgradeArmor(app fyne.App) fyne.CanvasObject {

	var data = [][]string{[]string{"Icon", "ItemName", "Quantity"},
		[]string{"", "Dummy Upgrade", "1x"}}

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("wide content"))
		},
		func(id widget.TableCellID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(data[id.Row][id.Col])
		})

	return table

}

func (w *win) funcbuttonsArmor(app fyne.App, li *widget.List) fyne.CanvasObject {
	add := widget.NewButton("Add", func() { //Button to Add Data
		w.window = app.NewWindow("Add Data")
		ID := widget.NewEntry()
		WeaponName := widget.NewEntry()

		inputmats := widget.NewButton("Low Rank", func() {
			container.NewGridWithColumns(4,
				container.NewGridWithRows(4,
					widget.NewEntry(),
					widget.NewEntry(),
					widget.NewEntry(),
					widget.NewEntry()),
				container.NewGridWithRows(4,
					widget.NewEntry(),
					widget.NewEntry(),
					widget.NewEntry(),
					widget.NewEntry()),
				container.NewGridWithRows(4,
					widget.NewEntry(),
					widget.NewEntry(),
					widget.NewEntry(),
					widget.NewEntry()),
				container.NewGridWithRows(4,
					widget.NewEntry(),
					widget.NewEntry(),
					widget.NewEntry(),
					widget.NewEntry()))

		})

		Selector := widget.NewButton("Select", func() { //Button to open file dialog
			dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error) {

				if e != nil {
					dialog.ShowInformation("Open Error", uc.URI().Path(), w.window)
					return
				}
				if _, ok := uc.(io.ReadSeeker); ok {
					dialog.ShowInformation("Seeker OK", uc.URI().Path(), w.window)

				}

			}, w.window).Show()
		})

		addData := widget.NewButton("Add", func() { //Button to add into MonsterName typed Data

			w.window.Close()
		})

		cancel := widget.NewButton("Cancel", func() {

			w.window.Close()
		})

		w.window.SetContent(container.New(layout.NewVBoxLayout(), ID, WeaponName, Selector, inputmats, addData, cancel)) //Layout for the "Insertion-Window"
		w.window.Resize(fyne.NewSize(400, 200))
		w.window.CenterOnScreen()
		w.window.Show()
	})

	exit := widget.NewButton("Close", func() { //added close button for whatever reason...cross-platform , maybe? idk

		w.window.Close()
	})

	return container.NewVBox(add, exit)
}
