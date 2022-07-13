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

func (w *win) MonsterUI(app fyne.App) {
	w.window = app.NewWindow("Monster")
	//id2 := 0 //passed to listupdate
	var id int
	monsterpic := widget.NewIcon(theme.AccountIcon())
	Monsters := decodemonsters()

	//materialButtons := container.NewHBox()
	//fmt.Println(Monsters[0].LRMat[0])

	list, id := initlist(Monsters, id)
	matlist := initmatlist()

	//materialButtons := w.initmaterialbuttons(app, id, matlist, Monsters)

	addbutton := w.monster_addbutton(app, id, matlist)
	buttons := container.NewVBox(addbutton)

	gbox := container.New(layout.NewGridLayout(3), list, buttons) //list with no data displayed as long as theres no item selected

	list.OnSelected = func(id widget.ListItemID) {
		//id2 = id
		updatebutton := w.monster_updatebutton(app, id, matlist, Monsters[id])
		deletebutton := w.monster_deletebutton(app, id, matlist, Monsters[id])

		buttons = container.NewVBox(addbutton, updatebutton, deletebutton)
		monsterpic = widget.NewIcon(fyne.NewStaticResource("icon", Monsters[id].pic))
		weakness := w.weakness(app, list, Monsters[id]) //assigns fyne.CanvasObject(GridWithColumns) to variable weakness
		matlist := initmatlist()
		materialButtons := w.initmaterialbuttons(app, id, matlist, Monsters)
		materials := container.NewGridWithRows(2, materialButtons, matlist)
		gbox = container.New(layout.NewGridLayout(3), list, monsterpic, materials, buttons, weakness) //display gbox
		w.window.SetContent(gbox)
		w.window.Show()

	}

	list.OnUnselected = func(id widget.ListItemID) {

		gbox = container.New(layout.NewGridLayout(3), list, buttons) //remove additional widgets
		w.window.SetContent(gbox)                                    //display gbox
		w.window.Show()
	}
	list.Select(125)

	w.window.SetContent(gbox) //Layout for the whole Monster-Window
	w.window.Resize(fyne.NewSize(400, 600))
	w.window.Show()

}

func (w *win) monster_addbutton(app fyne.App, id int, matlist *widget.List) fyne.CanvasObject {
	var tempmonster TempMonsterStruct
	for i := 0; i < 7; i++ {
		tempmonster.Fire[i] = 99
		tempmonster.Thunder[i] = 99
		tempmonster.Water[i] = 99
		tempmonster.Ice[i] = 99
		tempmonster.Dragon[i] = 99

	}
	/*for i := 0; i < 10; i++ {
		tempmonster.LRMat[i] = ""
		tempmonster.HRMat[i] = "a"
		tempmonster.GouRMat[i] = "a"
		tempmonster.GRMat[i] = "a"
		tempmonster.ZRMat[i] = "a"
	}*/

	add := widget.NewButton("Add", func() { //Button to Add Data

		wInput := app.NewWindow("Add Data")

		//Fire
		InputFireHead := widget.NewEntry()
		InputFireWings := widget.NewEntry()
		InputFireWingTailTip := widget.NewEntry()
		InputFireBelly := widget.NewEntry()
		InputFireBack := widget.NewEntry()
		InputFireTail := widget.NewEntry()
		InputFireLegs := widget.NewEntry()
		//Thunder
		InputThunderHead := widget.NewEntry()
		InputThunderWings := widget.NewEntry()
		InputThunderWingTailTip := widget.NewEntry()
		InputThunderBelly := widget.NewEntry()
		InputThunderBack := widget.NewEntry()
		InputThunderTail := widget.NewEntry()
		InputThunderLegs := widget.NewEntry()
		//Water
		InputWaterHead := widget.NewEntry()
		InputWaterWings := widget.NewEntry()
		InputWaterWingTailTip := widget.NewEntry()
		InputWaterBelly := widget.NewEntry()
		InputWaterBack := widget.NewEntry()
		InputWaterTail := widget.NewEntry()
		InputWaterLegs := widget.NewEntry()
		//Ice
		InputIceHead := widget.NewEntry()
		InputIceWings := widget.NewEntry()
		InputIceWingTailTip := widget.NewEntry()
		InputIceBelly := widget.NewEntry()
		InputIceBack := widget.NewEntry()
		InputIceTail := widget.NewEntry()
		InputIceLegs := widget.NewEntry()
		//Dragon
		InputDragonHead := widget.NewEntry()
		InputDragonWings := widget.NewEntry()
		InputDragonWingTailTip := widget.NewEntry()
		InputDragonBelly := widget.NewEntry()
		InputDragonBack := widget.NewEntry()
		InputDragonTail := widget.NewEntry()
		InputDragonLegs := widget.NewEntry()

		//LRMat
		InputLRMat0 := widget.NewEntry()
		InputLRMat1 := widget.NewEntry()
		InputLRMat2 := widget.NewEntry()
		InputLRMat3 := widget.NewEntry()
		InputLRMat4 := widget.NewEntry()
		InputLRMat5 := widget.NewEntry()
		InputLRMat6 := widget.NewEntry()
		InputLRMat7 := widget.NewEntry()
		InputLRMat8 := widget.NewEntry()
		InputLRMat9 := widget.NewEntry()
		//HRMat
		InputHRMat0 := widget.NewEntry()
		InputHRMat1 := widget.NewEntry()
		InputHRMat2 := widget.NewEntry()
		InputHRMat3 := widget.NewEntry()
		InputHRMat4 := widget.NewEntry()
		InputHRMat5 := widget.NewEntry()
		InputHRMat6 := widget.NewEntry()
		InputHRMat7 := widget.NewEntry()
		InputHRMat8 := widget.NewEntry()
		InputHRMat9 := widget.NewEntry()
		//GouRMat
		InputGouRMat0 := widget.NewEntry()
		InputGouRMat1 := widget.NewEntry()
		InputGouRMat2 := widget.NewEntry()
		InputGouRMat3 := widget.NewEntry()
		InputGouRMat4 := widget.NewEntry()
		InputGouRMat5 := widget.NewEntry()
		InputGouRMat6 := widget.NewEntry()
		InputGouRMat7 := widget.NewEntry()
		InputGouRMat8 := widget.NewEntry()
		InputGouRMat9 := widget.NewEntry()
		//GRMat
		InputGRMat0 := widget.NewEntry()
		InputGRMat1 := widget.NewEntry()
		InputGRMat2 := widget.NewEntry()
		InputGRMat3 := widget.NewEntry()
		InputGRMat4 := widget.NewEntry()
		InputGRMat5 := widget.NewEntry()
		InputGRMat6 := widget.NewEntry()
		InputGRMat7 := widget.NewEntry()
		InputGRMat8 := widget.NewEntry()
		InputGRMat9 := widget.NewEntry()
		//ZRMat
		InputZRMat0 := widget.NewEntry()
		InputZRMat1 := widget.NewEntry()
		InputZRMat2 := widget.NewEntry()
		InputZRMat3 := widget.NewEntry()
		InputZRMat4 := widget.NewEntry()
		InputZRMat5 := widget.NewEntry()
		InputZRMat6 := widget.NewEntry()
		InputZRMat7 := widget.NewEntry()
		InputZRMat8 := widget.NewEntry()
		InputZRMat9 := widget.NewEntry()

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
				InputFireHead,
				InputFireWings,
				InputFireWingTailTip,
				InputFireBelly,
				InputFireBack,
				InputFireTail,
				InputFireLegs),

			container.NewGridWithRows(8,
				widget.NewLabel("Thunder"),
				InputThunderHead,
				InputThunderWings,
				InputThunderWingTailTip,
				InputThunderBelly,
				InputThunderBack,
				InputThunderTail,
				InputThunderLegs),

			container.NewGridWithRows(8,
				widget.NewLabel("Water"),
				InputWaterHead,
				InputWaterWings,
				InputWaterWingTailTip,
				InputWaterBelly,
				InputWaterBack,
				InputWaterTail,
				InputWaterLegs),

			container.NewGridWithRows(8,
				widget.NewLabel("Ice"),
				InputIceHead,
				InputIceWings,
				InputIceWingTailTip,
				InputIceBelly,
				InputIceBack,
				InputIceTail,
				InputIceLegs),

			container.NewGridWithRows(8,
				widget.NewLabel("Dragon"),
				InputDragonHead,
				InputDragonWings,
				InputDragonWingTailTip,
				InputDragonBelly,
				InputDragonBack,
				InputDragonTail,
				InputDragonLegs),
		)

		inputmats := container.NewGridWithColumns(6,
			container.NewGridWithRows(10,
				widget.NewLabel("Rank:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:")),

			container.NewGridWithRows(10,
				widget.NewLabel("Low:"),
				InputLRMat0,
				InputLRMat1,
				InputLRMat2,
				InputLRMat3,
				InputLRMat4,
				InputLRMat5,
				InputLRMat6,
				InputLRMat7,
				InputLRMat8,
				InputLRMat9),

			container.NewGridWithRows(10,
				widget.NewLabel("High:"),
				InputHRMat0,
				InputHRMat1,
				InputHRMat2,
				InputHRMat3,
				InputHRMat4,
				InputHRMat5,
				InputHRMat6,
				InputHRMat7,
				InputHRMat8,
				InputHRMat9),

			container.NewGridWithRows(10,
				widget.NewLabel("Gou:"),
				InputGouRMat0,
				InputGouRMat1,
				InputGouRMat2,
				InputGouRMat3,
				InputGouRMat4,
				InputGouRMat5,
				InputGouRMat6,
				InputGouRMat7,
				InputGouRMat8,
				InputGouRMat9),

			container.NewGridWithRows(10,
				widget.NewLabel("G:"),
				InputGRMat0,
				InputGRMat1,
				InputGRMat2,
				InputGRMat3,
				InputGRMat4,
				InputGRMat5,
				InputGRMat6,
				InputGRMat7,
				InputGRMat8,
				InputGRMat9),

			container.NewGridWithRows(10,
				widget.NewLabel("Zenith:"),
				InputZRMat0,
				InputZRMat1,
				InputZRMat2,
				InputZRMat3,
				InputZRMat4,
				InputZRMat5,
				InputZRMat6,
				InputZRMat7,
				InputZRMat8,
				InputZRMat9))

		monstericon := widget.NewButton("Choose Monster-Icon (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, wInput)
					return
				}
				if reader == nil {
					log.Println("Cancelled")
					return
				}

				tempmonster.EncodedIcon = imageOpenedMonsterIcon(reader, tempmonster)
			}, wInput)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()
		})

		monsterpic := widget.NewButton("Choose Monster-Pic... (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, wInput)
					return
				}
				if reader == nil {
					log.Println("Cancelled")
					return
				}

				tempmonster.EncodedPic = imageOpenedMonsterPic(reader, tempmonster)
			}, wInput)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()
		})

		addData := widget.NewButton("Add", func() { //Button to add into MonsterName typed Data

			tempmonster.Name = MonsterName.Text

			/*Get Element-Weaknesses*/
			//Get Fire-Values
			tempmonster.Fire[0], _ = strconv.Atoi(InputFireHead.Text)
			tempmonster.Fire[1], _ = strconv.Atoi(InputFireWings.Text)
			tempmonster.Fire[2], _ = strconv.Atoi(InputFireWingTailTip.Text)
			tempmonster.Fire[3], _ = strconv.Atoi(InputFireBelly.Text)
			tempmonster.Fire[4], _ = strconv.Atoi(InputFireBack.Text)
			tempmonster.Fire[5], _ = strconv.Atoi(InputFireTail.Text)
			tempmonster.Fire[6], _ = strconv.Atoi(InputFireLegs.Text)
			//Get Thunder-Values
			tempmonster.Thunder[0], _ = strconv.Atoi(InputThunderHead.Text)
			tempmonster.Thunder[1], _ = strconv.Atoi(InputThunderWings.Text)
			tempmonster.Thunder[2], _ = strconv.Atoi(InputThunderWingTailTip.Text)
			tempmonster.Thunder[3], _ = strconv.Atoi(InputThunderBelly.Text)
			tempmonster.Thunder[4], _ = strconv.Atoi(InputThunderBack.Text)
			tempmonster.Thunder[5], _ = strconv.Atoi(InputThunderTail.Text)
			tempmonster.Thunder[6], _ = strconv.Atoi(InputThunderLegs.Text)
			//Get Water-Values
			tempmonster.Water[0], _ = strconv.Atoi(InputWaterHead.Text)
			tempmonster.Water[1], _ = strconv.Atoi(InputWaterWings.Text)
			tempmonster.Water[2], _ = strconv.Atoi(InputWaterWingTailTip.Text)
			tempmonster.Water[3], _ = strconv.Atoi(InputWaterBelly.Text)
			tempmonster.Water[4], _ = strconv.Atoi(InputWaterBack.Text)
			tempmonster.Water[5], _ = strconv.Atoi(InputWaterTail.Text)
			tempmonster.Water[6], _ = strconv.Atoi(InputWaterLegs.Text)
			//Get Ice-Values
			tempmonster.Ice[0], _ = strconv.Atoi(InputIceHead.Text)
			tempmonster.Ice[1], _ = strconv.Atoi(InputIceWings.Text)
			tempmonster.Ice[2], _ = strconv.Atoi(InputIceWingTailTip.Text)
			tempmonster.Ice[3], _ = strconv.Atoi(InputIceBelly.Text)
			tempmonster.Ice[4], _ = strconv.Atoi(InputIceBack.Text)
			tempmonster.Ice[5], _ = strconv.Atoi(InputIceTail.Text)
			tempmonster.Ice[6], _ = strconv.Atoi(InputIceLegs.Text)
			//Get Dragon-Values
			tempmonster.Dragon[0], _ = strconv.Atoi(InputDragonHead.Text)
			tempmonster.Dragon[1], _ = strconv.Atoi(InputDragonWings.Text)
			tempmonster.Dragon[2], _ = strconv.Atoi(InputDragonWingTailTip.Text)
			tempmonster.Dragon[3], _ = strconv.Atoi(InputDragonBelly.Text)
			tempmonster.Dragon[4], _ = strconv.Atoi(InputDragonBack.Text)
			tempmonster.Dragon[5], _ = strconv.Atoi(InputDragonTail.Text)
			tempmonster.Dragon[6], _ = strconv.Atoi(InputDragonLegs.Text)

			/*Get Materials*/
			//Low Rank Mats
			tempmonster.LRMat[0] = InputLRMat0.Text
			tempmonster.LRMat[2] = InputLRMat2.Text
			tempmonster.LRMat[3] = InputLRMat3.Text
			tempmonster.LRMat[4] = InputLRMat4.Text
			tempmonster.LRMat[5] = InputLRMat5.Text
			tempmonster.LRMat[6] = InputLRMat6.Text
			tempmonster.LRMat[7] = InputLRMat7.Text
			tempmonster.LRMat[8] = InputLRMat8.Text
			tempmonster.LRMat[9] = InputLRMat9.Text
			//High Rank Mats
			tempmonster.HRMat[0] = InputHRMat0.Text
			tempmonster.HRMat[1] = InputHRMat1.Text
			tempmonster.HRMat[2] = InputHRMat2.Text
			tempmonster.HRMat[3] = InputHRMat3.Text
			tempmonster.HRMat[4] = InputHRMat4.Text
			tempmonster.HRMat[5] = InputHRMat5.Text
			tempmonster.HRMat[6] = InputHRMat6.Text
			tempmonster.HRMat[7] = InputHRMat7.Text
			tempmonster.HRMat[8] = InputHRMat8.Text
			tempmonster.HRMat[9] = InputHRMat9.Text
			//Gou Rank Mats
			tempmonster.GouRMat[0] = InputGouRMat0.Text
			tempmonster.GouRMat[1] = InputGouRMat1.Text
			tempmonster.GouRMat[2] = InputGouRMat2.Text
			tempmonster.GouRMat[3] = InputGouRMat3.Text
			tempmonster.GouRMat[4] = InputGouRMat4.Text
			tempmonster.GouRMat[5] = InputGouRMat5.Text
			tempmonster.GouRMat[6] = InputGouRMat6.Text
			tempmonster.GouRMat[7] = InputGouRMat7.Text
			tempmonster.GouRMat[8] = InputGouRMat8.Text
			tempmonster.GouRMat[9] = InputGouRMat9.Text
			//G Rank Mats
			tempmonster.GRMat[0] = InputGRMat0.Text
			tempmonster.GRMat[1] = InputGRMat1.Text
			tempmonster.GRMat[2] = InputGRMat2.Text
			tempmonster.GRMat[3] = InputGRMat3.Text
			tempmonster.GRMat[4] = InputGRMat4.Text
			tempmonster.GRMat[5] = InputGRMat5.Text
			tempmonster.GRMat[6] = InputGRMat6.Text
			tempmonster.GRMat[7] = InputGRMat7.Text
			tempmonster.GRMat[8] = InputGRMat8.Text
			tempmonster.GRMat[9] = InputGRMat9.Text
			//Z Rank Mats
			tempmonster.ZRMat[0] = InputZRMat0.Text
			tempmonster.ZRMat[1] = InputZRMat1.Text
			tempmonster.ZRMat[2] = InputZRMat2.Text
			tempmonster.ZRMat[3] = InputZRMat3.Text
			tempmonster.ZRMat[4] = InputZRMat4.Text
			tempmonster.ZRMat[5] = InputZRMat5.Text
			tempmonster.ZRMat[6] = InputZRMat6.Text
			tempmonster.ZRMat[7] = InputZRMat7.Text
			tempmonster.ZRMat[8] = InputZRMat8.Text
			tempmonster.ZRMat[9] = InputZRMat9.Text
			InsertOneMonster(client, ctx, tempmonster)
			w.listUpdate(app, id, matlist)
			wInput.Close()
		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wInput.Close()

		})

		wInput.SetContent(container.New(layout.NewVBoxLayout(), MonsterName, monstericon, monsterpic, Weaknesses, inputmats, addData, cancel)) //Layout for the "Insertion-Window"
		wInput.Resize(fyne.NewSize(400, 200))
		wInput.CenterOnScreen()
		wInput.Show()
	})

	return container.NewVBox(add)
}

func (w *win) weakness(app fyne.App, li *widget.List, Monster MonsterStruct) fyne.CanvasObject {

	var data = [8][6]string{[6]string{"Hitzone", "Fire", "Thunder", "Water", "Ice", "Dragon"}, //this inits the table-data, it'll be replaced by MongoDB data later on
		[6]string{"Head", strconv.Itoa(Monster.Fire[0]), strconv.Itoa(Monster.Thunder[0]), strconv.Itoa(Monster.Water[0]), strconv.Itoa(Monster.Ice[0]), strconv.Itoa(Monster.Dragon[0])},
		[6]string{"Wings", strconv.Itoa(Monster.Fire[1]), strconv.Itoa(Monster.Thunder[1]), strconv.Itoa(Monster.Water[1]), strconv.Itoa(Monster.Ice[1]), strconv.Itoa(Monster.Dragon[1])},
		[6]string{"Wing/Tail Wip", strconv.Itoa(Monster.Fire[2]), strconv.Itoa(Monster.Thunder[2]), strconv.Itoa(Monster.Water[2]), strconv.Itoa(Monster.Ice[2]), strconv.Itoa(Monster.Dragon[2])},
		[6]string{"Belly", strconv.Itoa(Monster.Fire[3]), strconv.Itoa(Monster.Thunder[3]), strconv.Itoa(Monster.Water[3]), strconv.Itoa(Monster.Ice[3]), strconv.Itoa(Monster.Dragon[3])},
		[6]string{"Back", strconv.Itoa(Monster.Fire[4]), strconv.Itoa(Monster.Thunder[4]), strconv.Itoa(Monster.Water[4]), strconv.Itoa(Monster.Ice[4]), strconv.Itoa(Monster.Dragon[4])},
		[6]string{"Tail", strconv.Itoa(Monster.Fire[5]), strconv.Itoa(Monster.Thunder[5]), strconv.Itoa(Monster.Water[5]), strconv.Itoa(Monster.Ice[5]), strconv.Itoa(Monster.Dragon[5])},
		[6]string{"Legs", strconv.Itoa(Monster.Fire[6]), strconv.Itoa(Monster.Thunder[6]), strconv.Itoa(Monster.Water[6]), strconv.Itoa(Monster.Ice[6]), strconv.Itoa(Monster.Dragon[6])}}

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

func (w *win) materialsLR(app fyne.App, Monster MonsterStruct) *widget.List {

	matlist := widget.NewList(
		func() int {
			return len(Monster.LRMat)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Monster", theme.AccountIcon().Content())), widget.NewLabel("Template Label")) //creates a HBox for every row of the list
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Monster.Name, Monster.icon)) //assigns monster icon to list
			c.Objects[1].(*widget.Label).SetText(Monster.LRMat[id])                                     //assigns monster name to list
		},
	)

	return matlist

}

func (w *win) materialsHR(app fyne.App, Monster MonsterStruct) *widget.List {

	matlist := widget.NewList(
		func() int {
			return len(Monster.HRMat)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Monster", theme.AccountIcon().Content())), widget.NewLabel("Template Label")) //creates a HBox for every row of the list
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Monster.Name, Monster.icon)) //assigns monster icon to list
			c.Objects[1].(*widget.Label).SetText(Monster.HRMat[id])                                     //assigns monster name to list
		},
	)

	return matlist

}

func (w *win) materialsGouR(app fyne.App, Monster MonsterStruct) *widget.List {

	matlist := widget.NewList(
		func() int {
			return len(Monster.GouRMat)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Monster", theme.AccountIcon().Content())), widget.NewLabel("Template Label")) //creates a HBox for every row of the list
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Monster.Name, Monster.icon)) //assigns monster icon to list
			c.Objects[1].(*widget.Label).SetText(Monster.GouRMat[id])                                   //assigns monster name to list
		},
	)

	return matlist

}

func (w *win) materialsG(app fyne.App, Monster MonsterStruct) *widget.List {

	matlist := widget.NewList(
		func() int {
			return len(Monster.ZRMat)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Monster", theme.AccountIcon().Content())), widget.NewLabel("Template Label")) //creates a HBox for every row of the list
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Monster.Name, Monster.icon)) //assigns monster icon to list
			c.Objects[1].(*widget.Label).SetText(Monster.ZRMat[id])                                     //assigns monster name to list
		},
	)

	return matlist

}

func (w *win) materialsZenith(app fyne.App, Monster MonsterStruct) *widget.List {

	matlist := widget.NewList(
		func() int {
			return len(Monster.GRMat)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Monster", theme.AccountIcon().Content())), widget.NewLabel("Template Label")) //creates a HBox for every row of the list
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Monster.Name, Monster.icon)) //assigns monster icon to list
			c.Objects[1].(*widget.Label).SetText(Monster.GRMat[id])                                     //assigns monster name to list
		},
	)

	return matlist

}

func (w *win) listUpdate(app fyne.App, id widget.ListItemID,
	matlist *widget.List) { //function updates materials when another Rank was selected
	Monsters := decodemonsters()
	list, id := initlist(Monsters, id)
	materialButtons := w.initmaterialbuttons(app, id, matlist, Monsters)
	addbutton := w.monster_addbutton(app, id, matlist)
	updatebutton := w.monster_updatebutton(app, id, matlist, Monsters[id])
	deletebutton := w.monster_deletebutton(app, id, matlist, Monsters[id])

	monsterpic := widget.NewIcon(fyne.NewStaticResource("icon", Monsters[id].pic))
	buttons := container.NewVBox(addbutton, updatebutton, deletebutton)
	weakness := w.weakness(app, list, Monsters[id])
	materials := container.NewGridWithRows(2, materialButtons, matlist)
	gbox := container.New(layout.NewGridLayout(3), list, monsterpic, materials, buttons, weakness)

	list.OnSelected = func(id widget.ListItemID) {
		//id2 = id
		updatebutton := w.monster_updatebutton(app, id, matlist, Monsters[id])
		deletebutton := w.monster_deletebutton(app, id, matlist, Monsters[id])

		buttons = container.NewVBox(addbutton, updatebutton, deletebutton)
		monsterpic = widget.NewIcon(fyne.NewStaticResource("icon", Monsters[id].pic))
		weakness := w.weakness(app, list, Monsters[id]) //assigns fyne.CanvasObject(GridWithColumns) to variable weakness
		matlist := initmatlist()
		materialButtons := w.initmaterialbuttons(app, id, matlist, Monsters)
		materials := container.NewGridWithRows(2, materialButtons, matlist)
		gbox = container.New(layout.NewGridLayout(3), list, monsterpic, materials, buttons, weakness) //display gbox
		w.window.SetContent(gbox)
		w.window.Show()

	}

	w.window.SetContent(gbox)
	w.window.Show()
}

func initmatlist() *widget.List {
	matlist := widget.NewList(
		func() int {
			return 1
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Item", theme.AccountIcon().Content())), widget.NewLabel("Template Label")) //creates a HBox for every row of the list
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(theme.AccountIcon()) //assigns monster icon to list
			c.Objects[1].(*widget.Label).SetText("")                     //assigns monster name to list
		},
	)

	return matlist
}

func initlist(Monsters []MonsterStruct, id widget.ListItemID) (*widget.List, widget.ListItemID) {

	list := widget.NewList(
		func() int {
			return len(Monsters)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Monster", theme.AccountIcon().Content())), widget.NewLabel("Template Label")) //creates a HBox for every row of the list
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Monsters[id].Name, Monsters[id].icon)) //assigns monster icon to list
			c.Objects[1].(*widget.Label).SetText(Monsters[id].Name)                                               //assigns monster name to list
		},
	)

	return list, id
}

func (w *win) monster_updatebutton(app fyne.App, id int, matlist *widget.List, Monster MonsterStruct) fyne.CanvasObject {
	var tempmonster TempMonsterStruct

	tempmonster.Name = Monster.Name
	tempmonster.EncodedIcon = Monster.EncodedIcon
	tempmonster.EncodedPic = Monster.EncodedPic
	for i := 0; i <= 6; i++ {
		tempmonster.Fire[i] = Monster.Fire[i]
		tempmonster.Thunder[i] = Monster.Thunder[i]
		tempmonster.Water[i] = Monster.Water[i]
		tempmonster.Ice[i] = Monster.Ice[i]
		tempmonster.Dragon[i] = Monster.Dragon[i]

	}

	for i := 0; i <= 9; i++ {
		tempmonster.LRMat[i] = Monster.LRMat[i]
		tempmonster.HRMat[i] = Monster.HRMat[i]
		tempmonster.GouRMat[i] = Monster.GouRMat[i]
		tempmonster.GRMat[i] = Monster.GRMat[i]
		tempmonster.ZRMat[i] = Monster.ZRMat[i]
	}

	update := widget.NewButton("Update", func() { //Button to Add Data
		wInput := app.NewWindow("Update Data")

		//Fire
		InputFireHead := widget.NewEntry()
		InputFireWings := widget.NewEntry()
		InputFireWingTailTip := widget.NewEntry()
		InputFireBelly := widget.NewEntry()
		InputFireBack := widget.NewEntry()
		InputFireTail := widget.NewEntry()
		InputFireLegs := widget.NewEntry()
		//Thunder
		InputThunderHead := widget.NewEntry()
		InputThunderWings := widget.NewEntry()
		InputThunderWingTailTip := widget.NewEntry()
		InputThunderBelly := widget.NewEntry()
		InputThunderBack := widget.NewEntry()
		InputThunderTail := widget.NewEntry()
		InputThunderLegs := widget.NewEntry()
		//Water
		InputWaterHead := widget.NewEntry()
		InputWaterWings := widget.NewEntry()
		InputWaterWingTailTip := widget.NewEntry()
		InputWaterBelly := widget.NewEntry()
		InputWaterBack := widget.NewEntry()
		InputWaterTail := widget.NewEntry()
		InputWaterLegs := widget.NewEntry()
		//Ice
		InputIceHead := widget.NewEntry()
		InputIceWings := widget.NewEntry()
		InputIceWingTailTip := widget.NewEntry()
		InputIceBelly := widget.NewEntry()
		InputIceBack := widget.NewEntry()
		InputIceTail := widget.NewEntry()
		InputIceLegs := widget.NewEntry()
		//Dragon
		InputDragonHead := widget.NewEntry()
		InputDragonWings := widget.NewEntry()
		InputDragonWingTailTip := widget.NewEntry()
		InputDragonBelly := widget.NewEntry()
		InputDragonBack := widget.NewEntry()
		InputDragonTail := widget.NewEntry()
		InputDragonLegs := widget.NewEntry()

		//LRMat
		InputLRMat0 := widget.NewEntry()
		InputLRMat1 := widget.NewEntry()
		InputLRMat2 := widget.NewEntry()
		InputLRMat3 := widget.NewEntry()
		InputLRMat4 := widget.NewEntry()
		InputLRMat5 := widget.NewEntry()
		InputLRMat6 := widget.NewEntry()
		InputLRMat7 := widget.NewEntry()
		InputLRMat8 := widget.NewEntry()
		InputLRMat9 := widget.NewEntry()
		//HRMat
		InputHRMat0 := widget.NewEntry()
		InputHRMat1 := widget.NewEntry()
		InputHRMat2 := widget.NewEntry()
		InputHRMat3 := widget.NewEntry()
		InputHRMat4 := widget.NewEntry()
		InputHRMat5 := widget.NewEntry()
		InputHRMat6 := widget.NewEntry()
		InputHRMat7 := widget.NewEntry()
		InputHRMat8 := widget.NewEntry()
		InputHRMat9 := widget.NewEntry()
		//GouRMat
		InputGouRMat0 := widget.NewEntry()
		InputGouRMat1 := widget.NewEntry()
		InputGouRMat2 := widget.NewEntry()
		InputGouRMat3 := widget.NewEntry()
		InputGouRMat4 := widget.NewEntry()
		InputGouRMat5 := widget.NewEntry()
		InputGouRMat6 := widget.NewEntry()
		InputGouRMat7 := widget.NewEntry()
		InputGouRMat8 := widget.NewEntry()
		InputGouRMat9 := widget.NewEntry()
		//GRMat
		InputGRMat0 := widget.NewEntry()
		InputGRMat1 := widget.NewEntry()
		InputGRMat2 := widget.NewEntry()
		InputGRMat3 := widget.NewEntry()
		InputGRMat4 := widget.NewEntry()
		InputGRMat5 := widget.NewEntry()
		InputGRMat6 := widget.NewEntry()
		InputGRMat7 := widget.NewEntry()
		InputGRMat8 := widget.NewEntry()
		InputGRMat9 := widget.NewEntry()
		//ZRMat
		InputZRMat0 := widget.NewEntry()
		InputZRMat1 := widget.NewEntry()
		InputZRMat2 := widget.NewEntry()
		InputZRMat3 := widget.NewEntry()
		InputZRMat4 := widget.NewEntry()
		InputZRMat5 := widget.NewEntry()
		InputZRMat6 := widget.NewEntry()
		InputZRMat7 := widget.NewEntry()
		InputZRMat8 := widget.NewEntry()
		InputZRMat9 := widget.NewEntry()

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
				InputFireHead,
				InputFireWings,
				InputFireWingTailTip,
				InputFireBelly,
				InputFireBack,
				InputFireTail,
				InputFireLegs),

			container.NewGridWithRows(8,
				widget.NewLabel("Thunder"),
				InputThunderHead,
				InputThunderWings,
				InputThunderWingTailTip,
				InputThunderBelly,
				InputThunderBack,
				InputThunderTail,
				InputThunderLegs),

			container.NewGridWithRows(8,
				widget.NewLabel("Water"),
				InputWaterHead,
				InputWaterWings,
				InputWaterWingTailTip,
				InputWaterBelly,
				InputWaterBack,
				InputWaterTail,
				InputWaterLegs),

			container.NewGridWithRows(8,
				widget.NewLabel("Ice"),
				InputIceHead,
				InputIceWings,
				InputIceWingTailTip,
				InputIceBelly,
				InputIceBack,
				InputIceTail,
				InputIceLegs),

			container.NewGridWithRows(8,
				widget.NewLabel("Dragon"),
				InputDragonHead,
				InputDragonWings,
				InputDragonWingTailTip,
				InputDragonBelly,
				InputDragonBack,
				InputDragonTail,
				InputDragonLegs),
		)

		inputmats := container.NewGridWithColumns(6,
			container.NewGridWithRows(10,
				widget.NewLabel("Rank:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:"),
				widget.NewLabel("Item:")),

			container.NewGridWithRows(10,
				widget.NewLabel("Low:"),
				InputLRMat0,
				InputLRMat1,
				InputLRMat2,
				InputLRMat3,
				InputLRMat4,
				InputLRMat5,
				InputLRMat6,
				InputLRMat7,
				InputLRMat8,
				InputLRMat9),

			container.NewGridWithRows(10,
				widget.NewLabel("High:"),
				InputHRMat0,
				InputHRMat1,
				InputHRMat2,
				InputHRMat3,
				InputHRMat4,
				InputHRMat5,
				InputHRMat6,
				InputHRMat7,
				InputHRMat8,
				InputHRMat9),

			container.NewGridWithRows(10,
				widget.NewLabel("Gou:"),
				InputGouRMat0,
				InputGouRMat1,
				InputGouRMat2,
				InputGouRMat3,
				InputGouRMat4,
				InputGouRMat5,
				InputGouRMat6,
				InputGouRMat7,
				InputGouRMat8,
				InputGouRMat9),

			container.NewGridWithRows(10,
				widget.NewLabel("G:"),
				InputGRMat0,
				InputGRMat1,
				InputGRMat2,
				InputGRMat3,
				InputGRMat4,
				InputGRMat5,
				InputGRMat6,
				InputGRMat7,
				InputGRMat8,
				InputGRMat9),

			container.NewGridWithRows(10,
				widget.NewLabel("Zenith:"),
				InputZRMat0,
				InputZRMat1,
				InputZRMat2,
				InputZRMat3,
				InputZRMat4,
				InputZRMat5,
				InputZRMat6,
				InputZRMat7,
				InputZRMat8,
				InputZRMat9))

		monstericon := widget.NewButton("Choose Monster-Icon (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, wInput)
					return
				}
				if reader == nil {
					log.Println("Cancelled")
					return
				}

				tempmonster.EncodedIcon = imageOpenedMonsterIcon(reader, tempmonster)
			}, wInput)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()
		})

		monsterpic := widget.NewButton("Choose Monster-Pic... (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, wInput)
					return
				}
				if reader == nil {
					log.Println("Cancelled")
					return
				}

				tempmonster.EncodedPic = imageOpenedMonsterPic(reader, tempmonster)
			}, wInput)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()
		})

		updateData := widget.NewButton("Update", func() { //Button to add into MonsterName typed Data
			tempmonster.Name = MonsterName.Text

			/*Get Element-Weaknesses*/
			//Get Fire-Values
			if InputFireHead.Text != "" {
				tempmonster.Fire[0], _ = strconv.Atoi(InputFireHead.Text)
			}
			if InputFireWings.Text != "" {
				tempmonster.Fire[1], _ = strconv.Atoi(InputFireWings.Text)
			}
			if InputFireWingTailTip.Text != "" {
				tempmonster.Fire[2], _ = strconv.Atoi(InputFireWingTailTip.Text)
			}
			if InputFireBelly.Text != "" {
				tempmonster.Fire[3], _ = strconv.Atoi(InputFireBelly.Text)
			}
			if InputFireBack.Text != "" {
				tempmonster.Fire[4], _ = strconv.Atoi(InputFireBack.Text)
			}
			if InputFireTail.Text != "" {
				tempmonster.Fire[5], _ = strconv.Atoi(InputFireTail.Text)
			}
			if InputFireLegs.Text != "" {
				tempmonster.Fire[6], _ = strconv.Atoi(InputFireLegs.Text)
			}
			//Get Thunder-Values
			if InputThunderHead.Text != "" {
				tempmonster.Thunder[0], _ = strconv.Atoi(InputThunderHead.Text)
			}
			if InputThunderWings.Text != "" {
				tempmonster.Thunder[1], _ = strconv.Atoi(InputThunderWings.Text)
			}
			if InputThunderWingTailTip.Text != "" {
				tempmonster.Thunder[2], _ = strconv.Atoi(InputThunderWingTailTip.Text)
			}
			if InputThunderBelly.Text != "" {
				tempmonster.Thunder[3], _ = strconv.Atoi(InputThunderBelly.Text)
			}
			if InputThunderBack.Text != "" {
				tempmonster.Thunder[4], _ = strconv.Atoi(InputThunderBack.Text)
			}
			if InputThunderTail.Text != "" {
				tempmonster.Thunder[5], _ = strconv.Atoi(InputThunderTail.Text)
			}
			if InputThunderLegs.Text != "" {
				tempmonster.Thunder[6], _ = strconv.Atoi(InputThunderLegs.Text)
			}
			//Get Water-Values
			if InputWaterHead.Text != "" {
				tempmonster.Water[0], _ = strconv.Atoi(InputWaterHead.Text)
			}
			if InputWaterWings.Text != "" {
				tempmonster.Water[1], _ = strconv.Atoi(InputWaterWings.Text)
			}
			if InputWaterWingTailTip.Text != "" {
				tempmonster.Water[2], _ = strconv.Atoi(InputWaterWingTailTip.Text)
			}
			if InputWaterBelly.Text != "" {
				tempmonster.Water[3], _ = strconv.Atoi(InputWaterBelly.Text)
			}
			if InputWaterBack.Text != "" {
				tempmonster.Water[4], _ = strconv.Atoi(InputWaterBack.Text)
			}
			if InputWaterTail.Text != "" {
				tempmonster.Water[5], _ = strconv.Atoi(InputWaterTail.Text)
			}
			if InputWaterLegs.Text != "" {
				tempmonster.Water[6], _ = strconv.Atoi(InputWaterLegs.Text)
			}
			//Get Ice-Values
			if InputIceHead.Text != "" {
				tempmonster.Ice[0], _ = strconv.Atoi(InputIceHead.Text)
			}
			if InputIceWings.Text != "" {
				tempmonster.Ice[1], _ = strconv.Atoi(InputIceWings.Text)
			}
			if InputIceWingTailTip.Text != "" {
				tempmonster.Ice[2], _ = strconv.Atoi(InputIceWingTailTip.Text)
			}
			if InputIceBelly.Text != "" {
				tempmonster.Ice[3], _ = strconv.Atoi(InputIceBelly.Text)
			}
			if InputIceBack.Text != "" {
				tempmonster.Ice[4], _ = strconv.Atoi(InputIceBack.Text)
			}
			if InputIceTail.Text != "" {
				tempmonster.Ice[5], _ = strconv.Atoi(InputIceTail.Text)
			}
			if InputIceLegs.Text != "" {
				tempmonster.Ice[6], _ = strconv.Atoi(InputIceLegs.Text)
			}
			//Get Dragon-Values
			if InputDragonHead.Text != "" {
				tempmonster.Dragon[0], _ = strconv.Atoi(InputDragonHead.Text)
			}
			if InputDragonWings.Text != "" {
				tempmonster.Dragon[1], _ = strconv.Atoi(InputDragonWings.Text)
			}
			if InputDragonWingTailTip.Text != "" {
				tempmonster.Dragon[2], _ = strconv.Atoi(InputDragonWingTailTip.Text)
			}
			if InputDragonBelly.Text != "" {
				tempmonster.Dragon[3], _ = strconv.Atoi(InputDragonBelly.Text)
			}
			if InputDragonBack.Text != "" {
				tempmonster.Dragon[4], _ = strconv.Atoi(InputDragonBack.Text)
			}
			if InputDragonTail.Text != "" {
				tempmonster.Dragon[5], _ = strconv.Atoi(InputDragonTail.Text)
			}
			if InputDragonLegs.Text != "" {
				tempmonster.Dragon[6], _ = strconv.Atoi(InputDragonLegs.Text)
			}

			/*Get Materials*/
			//Low Rank Mats
			if InputLRMat0.Text != "" {
				tempmonster.LRMat[0] = InputLRMat0.Text
			}
			if InputLRMat1.Text != "" {
				tempmonster.LRMat[1] = InputLRMat1.Text
			}
			if InputLRMat2.Text != "" {
				tempmonster.LRMat[2] = InputLRMat2.Text
			}
			if InputLRMat3.Text != "" {
				tempmonster.LRMat[3] = InputLRMat3.Text
			}
			if InputLRMat4.Text != "" {
				tempmonster.LRMat[4] = InputLRMat4.Text
			}
			if InputLRMat5.Text != "" {
				tempmonster.LRMat[5] = InputLRMat5.Text
			}
			if InputLRMat6.Text != "" {
				tempmonster.LRMat[6] = InputLRMat6.Text
			}
			if InputLRMat7.Text != "" {
				tempmonster.LRMat[7] = InputLRMat7.Text
			}
			if InputLRMat8.Text != "" {
				tempmonster.LRMat[8] = InputLRMat8.Text
			}
			if InputLRMat9.Text != "" {
				tempmonster.LRMat[9] = InputLRMat9.Text
			}
			//High Rank Mats
			if InputHRMat0.Text != "" {
				tempmonster.HRMat[0] = InputHRMat0.Text
			}
			if InputHRMat1.Text != "" {
				tempmonster.HRMat[1] = InputHRMat1.Text
			}
			if InputHRMat2.Text != "" {
				tempmonster.HRMat[2] = InputHRMat2.Text
			}
			if InputHRMat3.Text != "" {
				tempmonster.HRMat[3] = InputHRMat3.Text
			}
			if InputHRMat4.Text != "" {
				tempmonster.HRMat[4] = InputHRMat4.Text
			}
			if InputHRMat5.Text != "" {
				tempmonster.HRMat[5] = InputHRMat5.Text
			}
			if InputHRMat6.Text != "" {
				tempmonster.HRMat[6] = InputHRMat6.Text
			}
			if InputHRMat7.Text != "" {
				tempmonster.HRMat[7] = InputHRMat7.Text
			}
			if InputHRMat8.Text != "" {
				tempmonster.HRMat[8] = InputHRMat8.Text
			}
			if InputHRMat9.Text != "" {
				tempmonster.HRMat[9] = InputHRMat9.Text
			}
			//Gou Rank Mats
			if InputGouRMat0.Text != "" {
				tempmonster.GouRMat[0] = InputGouRMat0.Text
			}
			if InputGouRMat1.Text != "" {
				tempmonster.GouRMat[1] = InputGouRMat1.Text
			}
			if InputGouRMat2.Text != "" {
				tempmonster.GouRMat[2] = InputGouRMat2.Text
			}
			if InputGouRMat3.Text != "" {
				tempmonster.GouRMat[3] = InputGouRMat3.Text
			}
			if InputGouRMat4.Text != "" {
				tempmonster.GouRMat[4] = InputGouRMat4.Text
			}
			if InputGouRMat5.Text != "" {
				tempmonster.GouRMat[5] = InputGouRMat5.Text
			}
			if InputGouRMat6.Text != "" {
				tempmonster.GouRMat[6] = InputGouRMat6.Text
			}
			if InputGouRMat7.Text != "" {
				tempmonster.GouRMat[7] = InputGouRMat7.Text
			}
			if InputGouRMat8.Text != "" {
				tempmonster.GouRMat[8] = InputGouRMat8.Text
			}
			if InputGouRMat9.Text != "" {
				tempmonster.GouRMat[9] = InputGouRMat9.Text
			}
			//G Rank Mats
			if InputGRMat0.Text != "" {
				tempmonster.GRMat[0] = InputGRMat0.Text
			}
			if InputGRMat1.Text != "" {
				tempmonster.GRMat[1] = InputGRMat1.Text
			}
			if InputGRMat2.Text != "" {
				tempmonster.GRMat[2] = InputGRMat2.Text
			}
			if InputGRMat3.Text != "" {
				tempmonster.GRMat[3] = InputGRMat3.Text
			}
			if InputGRMat4.Text != "" {
				tempmonster.GRMat[4] = InputGRMat4.Text
			}
			if InputGRMat5.Text != "" {
				tempmonster.GRMat[5] = InputGRMat5.Text
			}
			if InputGRMat6.Text != "" {
				tempmonster.GRMat[6] = InputGRMat6.Text
			}
			if InputGRMat7.Text != "" {
				tempmonster.GRMat[7] = InputGRMat7.Text
			}
			if InputGRMat8.Text != "" {
				tempmonster.GRMat[8] = InputGRMat8.Text
			}
			if InputGRMat9.Text != "" {
				tempmonster.GRMat[9] = InputGRMat9.Text
			}
			//Z Rank Mats
			if InputZRMat0.Text != "" {
				tempmonster.ZRMat[0] = InputZRMat0.Text
			}
			if InputZRMat1.Text != "" {
				tempmonster.ZRMat[1] = InputZRMat1.Text
			}
			if InputZRMat2.Text != "" {
				tempmonster.ZRMat[2] = InputZRMat2.Text
			}
			if InputZRMat3.Text != "" {
				tempmonster.ZRMat[3] = InputZRMat3.Text
			}
			if InputZRMat4.Text != "" {
				tempmonster.ZRMat[4] = InputZRMat4.Text
			}
			if InputZRMat5.Text != "" {
				tempmonster.ZRMat[5] = InputZRMat5.Text
			}
			if InputZRMat6.Text != "" {
				tempmonster.ZRMat[6] = InputZRMat6.Text
			}
			if InputZRMat7.Text != "" {
				tempmonster.ZRMat[7] = InputZRMat7.Text
			}
			if InputZRMat8.Text != "" {
				tempmonster.ZRMat[8] = InputZRMat8.Text
			}
			if InputZRMat9.Text != "" {
				tempmonster.ZRMat[9] = InputZRMat9.Text
			}
			UpdateOneMonster(client, ctx, Monster, tempmonster)
			w.listUpdate(app, id, matlist)
			wInput.Close()
		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wInput.Close()

		})

		wInput.SetContent(container.New(layout.NewVBoxLayout(), MonsterName, monstericon, monsterpic, Weaknesses, inputmats, updateData, cancel)) //Layout for the "Insertion-Window"
		wInput.Resize(fyne.NewSize(400, 200))
		wInput.CenterOnScreen()
		wInput.Show()
	})

	return container.NewVBox(update)
}

func (w *win) monster_deletebutton(app fyne.App, id int, matlist *widget.List, Monster MonsterStruct) fyne.CanvasObject {
	delete := widget.NewButton("Delete", func() { //Button to Delete Items
		DeleteOneMonster(client, ctx, Monster)
		id = id - 1
		if id >= 0 {
			w.listUpdate(app, id, matlist)
		}
		if id < 0 {
			id = 0
			list := initemptylist()
			addbutton := w.monster_addbutton(app, id, matlist)
			gbox := container.NewGridWithColumns(3, list, addbutton)
			w.window.SetContent(gbox)
			w.window.Show()

		}
	})

	return container.NewVBox(delete)
}

func (w *win) initmaterialbuttons(app fyne.App, id int, matlist *widget.List, Monsters []MonsterStruct) fyne.CanvasObject {
	LRButton := widget.NewButton("Low Rank", func() { //ex: assign low rank mats to table
		matlist = w.materialsLR(app, Monsters[id])
		//matlist.Refresh()
		w.listUpdate(app, id, matlist)
	})

	HRButton := widget.NewButton("High Rank", func() {
		matlist = w.materialsHR(app, Monsters[id])

		w.listUpdate(app, id, matlist)
	})

	GouRButton := widget.NewButton("Gou Rank", func() {
		matlist = w.materialsGouR(app, Monsters[id])
		matlist.Refresh()
		w.listUpdate(app, id, matlist)
	})

	GRButton := widget.NewButton("G Rank", func() {
		matlist = w.materialsG(app, Monsters[id])
		matlist.Refresh()
		w.listUpdate(app, id, matlist)
	})

	ZenithButton := widget.NewButton("Zenith Rank", func() {
		matlist = w.materialsZenith(app, Monsters[id])
		matlist.Refresh()
		w.listUpdate(app, id, matlist)
	})

	materialbuttons := container.NewHBox(LRButton, HRButton, GouRButton, GRButton, ZenithButton)

	return materialbuttons
}

//inits list ----can be used for every category----
func initemptylist() *widget.List {

	list := widget.NewList(
		func() int {
			return 0
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("Template Label")) //creates a HBox for every row of the list
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Label).SetText("")
		},
	)
	return list

}
