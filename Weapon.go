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

func (w *win) WeaponUI(app fyne.App) {
	w.window = app.NewWindow("Weapons")
	empty := widget.NewLabel("text")
	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")

	test := make([]string, 3) //just a test array, it'll probably replaced by data retrieved from MongoDB later on
	Weapondata := make([]string, 3)
	test[0] = "Dummy GreatSword"
	test[1] = "Dummy Lance"
	test[2] = "Dummy Bow"

	for i := range Weapondata {
		Weapondata[i] = strconv.Itoa(i+1) + " " + test[i]
	}

	list := widget.NewList(
		func() int {
			return len(Weapondata)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("Template Object"), widget.NewIcon(theme.DocumentIcon()))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(Weapondata[id])
		},
	)

	gbox := container.New(layout.NewGridLayout(3), list)

	list.OnSelected = func(id widget.ListItemID) {
		//id2 = id
		label.SetText(Weapondata[id])
		icon.SetResource(theme.DocumentIcon())

		buttons := w.funcbuttons(app, list) //assigns fyne.CanvasObject(HBOX) to variable buttons
		//assigns fyne.CanvasObject(GridWithColumns) to variable weakness
		//weaknesswidget := container.New(layout.NewGridWrapLayout(fyne.NewSize(600, 600)), weakness) //add additional widgets with Wrap to adjust TextSize
		//materials := w.materials(app)
		TableForgeMatsWeapon := w.materialsforgeWeapon(app)
		TableUpgradeMatsWeapon := w.materialsupgradeWeapon(app)

		gbox = container.New(layout.NewGridLayout(2), list, TableForgeMatsWeapon, buttons, TableUpgradeMatsWeapon) //display gbox
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

	w.window.SetContent(container.New(layout.NewGridLayout(2), list, empty))
	//w.window= SetContent(grid)
	w.window.Show()
}

func (w *win) materialsforgeWeapon(app fyne.App) fyne.CanvasObject {

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

func (w *win) materialsupgradeWeapon(app fyne.App) fyne.CanvasObject {

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

func (w *win) funcbuttonsWeapon(app fyne.App, li *widget.List) fyne.CanvasObject {
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
