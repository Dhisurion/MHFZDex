package main

import (

	//"os"

	//"fyne.io/fyne/dialog"

	//"container/list"
	"io"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/*MonsterStruct := struct {
	ID string
	//icon fyne.Resource
	Name string
}{}*/

func (w *win) MonsterUI(app fyne.App) {
	w.window = app.NewWindow("Monster")
	iconScreen(w.window)
	//initIcons()
	//empty := widget.NewLabel("text")
	test := make([]string, 3)

	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")

	//hbox := container.NewHBox(icon, label)

	test[0] = "Brachydios"
	test[1] = "Rathalos"
	test[2] = "Rathian"

	//test := []string{"Brachydios", "Rathalos", "Rathian"}

	data := make([]string, 3)

	for i := range data {
		data[i] = strconv.Itoa(i+1) + " " + test[i]
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

	gbox := container.New(layout.NewGridLayout(3), list)

	list.OnSelected = func(id widget.ListItemID) {
		label.SetText(data[id])
		icon.SetResource(theme.DocumentIcon())

		buttons := w.funcbuttons(app, list)                                                         //assigns fyne.CanvasObject(HBOX) to variable buttons
		weakness := w.weakness(app, list)                                                           //assigns fyne.CanvasObject(GridWithColumns) to variable weakness
		weaknesswidget := container.New(layout.NewGridWrapLayout(fyne.NewSize(600, 600)), weakness) //add additional widgets with Wrap to adjust TextSize
		materials := w.materials(app)
		gbox = container.New(layout.NewAdaptiveGridLayout(3), list, buttons, weaknesswidget, materials) //display gbox
		w.window.SetContent(gbox)
		w.window.Show()

	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
		gbox = container.New(layout.NewAdaptiveGridLayout(3), list) //remove additional widgets
		w.window.SetContent(gbox)                                   //display gbox
		w.window.Show()
	}
	list.Select(125)

	w.window.SetContent(gbox) //Layout for the whole Monster-Window
	w.window.Resize(fyne.NewSize(400, 600))
	w.window.Show()

}

func (w *win) funcbuttons(app fyne.App, li *widget.List) fyne.CanvasObject {
	var F [7]string

	FireHead := widget.NewEntry()
	FireWings := widget.NewEntry()
	FireWingTailTip := widget.NewEntry()
	FireBelly := widget.NewEntry()
	FireBack := widget.NewEntry()
	FireTail := widget.NewEntry()
	FireLegs := widget.NewEntry()
	//F[0] = widget.NewEntry().Text
	add := widget.NewButton("Add", func() { //Button to Add Data
		w.window = app.NewWindow("Add Data")
		//FH := widget.NewEntry()
		ID := widget.NewEntry()
		MonsterName := widget.NewEntry()
		Weaknesses := container.NewGridWithColumns(6,
			container.NewGridWithRows(8,
				widget.NewLabel("Hitzone"),
				widget.NewLabel("Head"),
				widget.NewLabel("Wings"),
				widget.NewLabel("Wing/Tail Tip"),
				widget.NewLabel("Belly"),
				widget.NewLabel("Back"),
				widget.NewLabel("Tail"),
				widget.NewLabel("Legs")),

			container.NewGridWithRows(8,
				widget.NewLabel("Fire"),
				FireHead,
				FireWings,
				FireWingTailTip,
				FireBelly,
				FireBack,
				FireTail,
				FireLegs),

			container.NewGridWithRows(8,
				widget.NewLabel("Thunder"),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry()),

			container.NewGridWithRows(8,
				widget.NewLabel("Water"),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry()),

			container.NewGridWithRows(8,
				widget.NewLabel("Ice"),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry()),

			container.NewGridWithRows(8,
				widget.NewLabel("Dragon"),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry(),
				widget.NewEntry()),
		)

		F[0] = FireHead.Text
		F[1] = FireWings.Text
		F[2] = FireWingTailTip.Text
		F[3] = FireBelly.Text
		F[4] = FireBack.Text
		F[5] = FireTail.Text
		F[6] = FireLegs.Text

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

		w.window.SetContent(container.New(layout.NewVBoxLayout(), ID, MonsterName, Selector, Weaknesses, addData, cancel)) //Layout for the "Insertion-Window"
		w.window.Resize(fyne.NewSize(400, 200))
		w.window.CenterOnScreen()
		w.window.Show()
	})

	exit := widget.NewButton("Close", func() { //added close button for whatever reason...cross-platform , maybe? idk

		w.window.Close()
	})

	return container.NewVBox(add, exit)
}

func (w *win) weakness(app fyne.App, li *widget.List) fyne.CanvasObject {
	//w.window = app.NewWindow("Weakness")

	return container.NewGridWithColumns(6,
		container.NewGridWithRows(8,
			widget.NewLabel("Hitzone"),
			widget.NewLabel("Head"),
			widget.NewLabel("Wings"),
			widget.NewLabel("Wing/Tail Tip"),
			widget.NewLabel("Belly"),
			widget.NewLabel("Back"),
			widget.NewLabel("Tail"),
			widget.NewLabel("Legs"),
		),
		container.NewGridWithRows(8,
			widget.NewLabel("Fire"),
			widget.NewLabel("FH"),
			widget.NewLabel("FWs"),
			widget.NewLabel("FW/T T"),
			widget.NewLabel("FBe"),
			widget.NewLabel("FBck"),
			widget.NewLabel("FT"),
			widget.NewLabel("FL"),
		),

		container.NewGridWithRows(8,
			widget.NewLabel("Thunder"),
			widget.NewLabel("ThH"),
			widget.NewLabel("ThW"),
			widget.NewLabel("ThW/T T"),
			widget.NewLabel("ThBe"),
			widget.NewLabel("ThBck"),
			widget.NewLabel("ThT"),
			widget.NewLabel("ThL"),
		),

		container.NewGridWithRows(8,
			widget.NewLabel("Water"),
			widget.NewLabel("WH"),
			widget.NewLabel("WW"),
			widget.NewLabel("WW/T T"),
			widget.NewLabel("WBe"),
			widget.NewLabel("WBck"),
			widget.NewLabel("WT"),
			widget.NewLabel("WL"),
		),

		container.NewGridWithRows(8,
			widget.NewLabel("Ice"),
			widget.NewLabel("IH"),
			widget.NewLabel("IW"),
			widget.NewLabel("IW/T T"),
			widget.NewLabel("IBe"),
			widget.NewLabel("IBck"),
			widget.NewLabel("IT"),
			widget.NewLabel("IL"),
		),

		container.NewGridWithRows(8,
			widget.NewLabel("Dragon"),
			widget.NewLabel("DrH"),
			widget.NewLabel("DrnW"),
			widget.NewLabel("DrW/T T"),
			widget.NewLabel("DrBe"),
			widget.NewLabel("DrBck"),
			widget.NewLabel("DrT"),
			widget.NewLabel("DrL"),
		),
	)

}

func (w *win) materials(app fyne.App) fyne.CanvasObject {

	var data = [][]string{[]string{"top left", "top right"},
		[]string{"bottom left", "bottom right"}}

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	return list

}
