package main

import (

	//"os"

	//"fyne.io/fyne/dialog"

	//"container/list"

	"log"
	"strconv"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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
	id2 := 0 //passed to listupdate
	//initIcons()

	monsterpic := widget.NewIcon(theme.CancelIcon())
	test := make([]string, 3)        //just a test array, it'll probably replaced by data retrieved from MongoDB later on
	Monsterdata := make([]string, 3) //the  Monsterdata array used in list
	test[0] = "Brachydios"
	test[1] = "Rathalos"
	test[2] = "Rathian"
	materialButtons := container.NewHBox()

	//materials := container.NewGridWithRows(2)

	for i := range Monsterdata { //init with data from test array
		Monsterdata[i] = strconv.Itoa(i+1) + " " + test[i]
	}

	list := widget.NewList(
		func() int {
			return len(Monsterdata)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("Template Object"), widget.NewIcon(theme.DocumentIcon())) //creates a HBox for every row of the list
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(Monsterdata[id]) //assigns data to the box
		},
	)

	//icon := widget.NewIcon(nil)
	//label := widget.NewLabel("Select An Item From The List")

	var datat = [][]string{[]string{"Icon", "ItemName", "Quantity", "Price"},
		[]string{"", "", "", ""}}

	table := widget.NewTable( //empty Table, for which there will be data assigned to on ButtonClick
		func() (int, int) {
			return len(datat), len(datat[0])
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("wide content"))
		},
		func(id widget.TableCellID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(datat[id.Row][id.Col])
		})

	LRButton := widget.NewButton("Low Rank", func() { //ex: assign low rank mats to table
		table = w.materialsLR(app, *table)
		w.listUpdate(app, id2, list, monsterpic, table, materialButtons)
	})

	HRButton := widget.NewButton("High Rank", func() {
		table = w.materialsHR(app, *table)
		w.listUpdate(app, id2, list, monsterpic, table, materialButtons)
	})

	GButton := widget.NewButton("G Rank", func() {
		table = w.materialsG(app, *table)
		w.listUpdate(app, id2, list, monsterpic, table, materialButtons)
	})

	ZenithButton := widget.NewButton("Zenith Rank", func() {
		table = w.materialsZenith(app, *table)
		w.listUpdate(app, id2, list, monsterpic, table, materialButtons)
	})

	MusouButton := widget.NewButton("Musou Rank", func() {
		table = w.materialsMusou(app, *table)
		w.listUpdate(app, id2, list, monsterpic, table, materialButtons)
	})

	//hbox := container.NewHBox(icon, label)
	//buttons := w.monster_funcbuttons(app, list)
	buttons := monster_funcbuttons(app)
	//test := []string{"Brachydios", "Rathalos", "Rathian"}

	gbox := container.New(layout.NewGridLayout(3), list, buttons) //list with no data displayed as long as theres no item selected

	list.OnSelected = func(id widget.ListItemID) {
		id2 = id
		//label.SetText(Monsterdata[id])
		//icon.SetResource(theme.DocumentIcon())

		//assigns fyne.CanvasObject(HBOX) to variable buttons
		weakness := w.weakness(app, list) //assigns fyne.CanvasObject(GridWithColumns) to variable weakness
		//weaknesswidget := container.New(layout.NewGridWrapLayout(fyne.NewSize(600, 600)), weakness) //add additional widgets with Wrap to adjust TextSize
		//materials := w.materials(app)
		materialButtons = container.NewHBox(LRButton, HRButton, GButton, ZenithButton, MusouButton)
		materials := container.NewGridWithRows(2, materialButtons, table)
		gbox = container.New(layout.NewGridLayout(3), list, monsterpic, materials, buttons, weakness) //display gbox
		w.window.SetContent(gbox)
		w.window.Show()

	}

	list.OnUnselected = func(id widget.ListItemID) {
		//label.SetText("Select An Item From The List")
		//icon.SetResource(nil)
		gbox = container.New(layout.NewGridLayout(3), list, buttons) //remove additional widgets
		w.window.SetContent(gbox)                                    //display gbox
		w.window.Show()
	}
	list.Select(125)

	w.window.SetContent(gbox) //Layout for the whole Monster-Window
	w.window.Resize(fyne.NewSize(400, 600))
	w.window.Show()

}

func monster_funcbuttons(app fyne.App) fyne.CanvasObject {

	//F[0] = widget.NewEntry().Text
	add := widget.NewButton("Add", func() { //Button to Add Data
		wInput := app.NewWindow("Add Data")

		var F [7]string
		FireHead := widget.NewEntry()
		FireWings := widget.NewEntry()
		FireWingTailTip := widget.NewEntry()
		FireBelly := widget.NewEntry()
		FireBack := widget.NewEntry()
		FireTail := widget.NewEntry()
		FireLegs := widget.NewEntry()
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

		NumberOfItemsEntryField := widget.NewEntry() //init with "0", crash otherwise, because value is nil on start of the window
		Entries, err := strconv.Atoi("0")
		if err != nil {
			panic(err)
		}
		EntryButton := widget.NewButton("Set", func() { // this is the real data initialization
			Entries, err = strconv.Atoi(NumberOfItemsEntryField.Text)
			if err != nil {
				panic(err)
			}

		})

		EntryContainer := container.NewHBox(widget.NewLabel("Number of Entries: "), NumberOfItemsEntryField, EntryButton) //display everything Entry related in Box

		inputLR := widget.NewButton("Low Rank", func() {
			container.NewGridWithColumns(Entries, //want to create a variable "Table" here for mats input
				container.NewGridWithRows(4,
					widget.NewEntry(),
					widget.NewEntry(),
					widget.NewEntry(),
					widget.NewEntry()))

		})

		inputHR := widget.NewButton("High Rank", func() {

		})

		inputG := widget.NewButton("G Rank", func() {

		})

		inputZenith := widget.NewButton("Zenith Rank", func() {

		})

		inputMusou := widget.NewButton("Musou Rank", func() {

		})

		inputmaterialbuttons := container.NewHBox(inputLR, inputHR, inputG, inputZenith, inputMusou) //container with all matButtons

		Selector := widget.NewButton("File Open With Filter (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, wInput)
					return
				}
				if reader == nil {
					log.Println("Cancelled")
					return
				}

				imageOpened(reader)
			}, wInput)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()
		})

		addData := widget.NewButton("Add", func() { //Button to add into MonsterName typed Data

			wInput.Close()
		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wInput.Close()

		})

		wInput.SetContent(container.New(layout.NewVBoxLayout(), ID, MonsterName, Selector, Weaknesses, EntryContainer, inputmaterialbuttons, addData, cancel)) //Layout for the "Insertion-Window"
		wInput.Resize(fyne.NewSize(400, 200))
		wInput.CenterOnScreen()
		wInput.Show()
	})

	/*exit := widget.NewButton("Close", func() { //added close button for whatever reason...cross-platform , maybe? idk

		w.window.Close()
	})*/

	return container.NewVBox(add)
}

func (w *win) weakness(app fyne.App, li *widget.List) fyne.CanvasObject {

	var data = [][]string{[]string{"Hitzone", "Fire", "Thunder", "Water", "Ice", "Dragon"}, //this inits the table-data, it'll be replaced by MongoDB data later on
		[]string{"Head", "FH", "TH", "WH", "IH", "DH"},
		[]string{"Wings", "FW", "TW", "WW", "IW", "DW"},
		[]string{"Wing/Tail Wip", "FWTW", "TWTW", "WWTW", "IWTW", "DWTW"},
		[]string{"Belly", "FBe", "TBe", "WBe", "IBe", "DBe"},
		[]string{"Back", "FBck", "TBck", "WBck", "IBck", "DBck"},
		[]string{"Tail", "FT", "TT", "WT", "IT", "DT"},
		[]string{"Legs", "FL", "TL", "WL", "IL", "DL"}}

	table := widget.NewTable( //initialization of table
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	return table
}

func (w *win) materialsLR(app fyne.App, tr widget.Table) *widget.Table {

	var data = [][]string{[]string{"Icon", "ItemName", "Quantity", "Price"}, //ex : inits LR mat data
		[]string{"", "Dummy LR", "1x", "100z"}}

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

func (w *win) materialsHR(app fyne.App, tr widget.Table) *widget.Table {

	var data = [][]string{[]string{"Icon", "ItemName", "Quantity", "Price"},
		[]string{"", "Dummy HR", "1x", "100z"}}

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

func (w *win) materialsG(app fyne.App, tr widget.Table) *widget.Table {

	var data = [][]string{[]string{"Icon", "ItemName", "Quantity", "Price"},
		[]string{"", "Dummy G", "1x", "100z"}}

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

func (w *win) materialsZenith(app fyne.App, tr widget.Table) *widget.Table {

	var data = [][]string{[]string{"Icon", "ItemName", "Quantity", "Price"},
		[]string{"", "Dummy Zenith", "1x", "100z"}}

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

func (w *win) materialsMusou(app fyne.App, tr widget.Table) *widget.Table {

	var data = [][]string{[]string{"Icon", "ItemName", "Quantity", "Price"},
		[]string{"", "Dummy Musou", "1x", "100z"}}

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

func (w *win) listUpdate(app fyne.App, id widget.ListItemID, list *widget.List, monsterpic fyne.CanvasObject,
	table *widget.Table, materialButtons fyne.CanvasObject) { //function updates materials when another Rank was selected
	buttons := monster_funcbuttons(app)
	weakness := w.weakness(app, list)
	materials := container.NewGridWithRows(2, materialButtons, table)
	gbox := container.New(layout.NewGridLayout(3), list, monsterpic, materials, buttons, weakness)

	w.window.SetContent(gbox)
	w.window.Show()
}

/*func (w *win) weakness(app fyne.App, li *widget.List) fyne.CanvasObject {
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

}*/
