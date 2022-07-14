package main

import (
	"image/color"
	_ "image/png"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (w *win) WeaponUI(app fyne.App) {
	w.window = app.NewWindow("Weapons")
	var id widget.ListItemID
	Weapons := decodeweapons()

	weaponbuttons := w.weapon_addbutton(app, id)

	list, id := initList_Weapon(Weapons, id)
	list.Resize(fyne.NewSize(25, 25))
	list.OnSelected = func(id widget.ListItemID) {
		icon := widget.NewIcon(fyne.NewStaticResource("icon", Weapons[id].icon))

		weaponname := widget.NewLabel(Weapons[id].Name)

		updatebutton := w.weapon_updatebutton(app, Weapons[id], id)

		deletebutton := w.weapon_deletebutton(app, Weapons[id], id)

		cancel := widget.NewButton("Cancel", func() {
			w.window.Close()
		})
		choice := container.New(layout.NewVBoxLayout(), weaponname, updatebutton, deletebutton, cancel)
		weaponmaterial := w.weaponmaterial(app, list, Weapons[id])
		gbox := container.New(layout.NewGridLayout(3), container.NewHScroll(list), choice, weaponbuttons, weaponmaterial, icon) //display gbox
		w.window.SetContent(gbox)
		w.window.Show()

	}

	list.OnUnselected = func(id widget.ListItemID) {

		gbox := container.New(layout.NewGridLayout(3), list, weaponbuttons) //remove additional widgets
		w.window.SetContent(gbox)
		w.window.Resize(fyne.NewSize(400, 600)) //display gbox
		w.window.Show()
	}

	gbox := container.New(layout.NewGridLayout(3), container.NewHScroll(list), weaponbuttons)

	w.window.SetContent(gbox)
	w.window.Resize(fyne.NewSize(100, 600))
	w.window.Show()
}

/*func (w *win) materialsforgeWeapon(app fyne.App) fyne.CanvasObject {

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

}*/

func (w *win) weapon_addbutton(app fyne.App, id widget.ListItemID) fyne.CanvasObject {

	add := widget.NewButton("Add", func() { //Button to Add Data
		wInput := app.NewWindow("Add Data")

		TextWeaponName := canvas.NewText("Name:", color.White)
		TextWeaponKind := canvas.NewText("Type:", color.White)
		TextWeaponRarity := canvas.NewText("Rarity:", color.White)
		TextWeaponAttack := canvas.NewText("Attack:", color.White)
		TextWeaponElement := canvas.NewText("Element:", color.White)
		TextWeaponElementvalue := canvas.NewText("Element-Value:", color.White)
		TextWeaponSharpness := canvas.NewText("Sharpness:", color.White)
		TextWeaponAffinity := canvas.NewText("Affinity:", color.White)
		TextWeaponDefense := canvas.NewText("Defense:", color.White)
		TextWeaponPrice := canvas.NewText("Price:", color.White)

		InputWeaponName := widget.NewEntry()
		InputWeaponKind := widget.NewEntry()
		InputWeaponRarity := widget.NewEntry()
		InputWeaponAttack := widget.NewEntry()
		InputWeaponElement := widget.NewEntry()
		InputWeaponElementvalue := widget.NewEntry()
		InputWeaponSharpness := widget.NewEntry()
		InputWeaponAffinity := widget.NewEntry()
		InputWeaponDefense := widget.NewEntry()
		InputWeaponPrice := widget.NewEntry()

		//Object
		InputMaterial1 := widget.NewEntry()
		InputMaterial2 := widget.NewEntry()
		InputMaterial3 := widget.NewEntry()
		InputMaterial4 := widget.NewEntry()

		InputQtyMat1 := widget.NewEntry()
		InputQtyMat2 := widget.NewEntry()
		InputQtyMat3 := widget.NewEntry()
		InputQtyMat4 := widget.NewEntry()

		InputWeaponObject := container.NewGridWithColumns(2,
			container.NewGridWithRows(5,
				widget.NewLabel("Material:"),
				InputMaterial1,
				InputMaterial2,
				InputMaterial3,
				InputMaterial4),

			container.NewGridWithRows(5,
				widget.NewLabel("Quantity:"),
				InputQtyMat1,
				InputQtyMat2,
				InputQtyMat3,
				InputQtyMat4),
		)

		InputWeaponIcon := widget.NewButton("File Open With Filter (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, wInput)
					return
				}
				if reader == nil {
					log.Println("Cancelled")
					return
				}

				imageOpenedWeaponIcon(reader)
			}, wInput)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()

		})

		addData := widget.NewButton("Add", func() { //Button to add into ItemName typed Data
			tempweapon.Name = InputWeaponName.Text
			tempweapon.Kind = InputWeaponKind.Text
			tempweapon.Rarity, _ = strconv.Atoi(InputWeaponRarity.Text)

			tempweapon.Attack, _ = strconv.Atoi(InputWeaponAttack.Text)
			tempweapon.Element = InputWeaponElement.Text
			tempweapon.Elementvalue, _ = strconv.Atoi(InputWeaponElementvalue.Text)

			tempweapon.Sharpness = InputWeaponSharpness.Text
			tempweapon.Affinity, _ = strconv.Atoi(InputWeaponAffinity.Text)
			tempweapon.Defense, _ = strconv.Atoi(InputWeaponDefense.Text)
			tempweapon.Price, _ = strconv.Atoi(InputWeaponPrice.Text)

			tempweapon.Material[0] = InputMaterial1.Text
			tempweapon.Material[1] = InputMaterial2.Text
			tempweapon.Material[2] = InputMaterial3.Text
			tempweapon.Material[3] = InputMaterial4.Text

			tempweapon.Quantity[0], _ = strconv.Atoi(InputQtyMat1.Text)
			tempweapon.Quantity[1], _ = strconv.Atoi(InputQtyMat2.Text)
			tempweapon.Quantity[2], _ = strconv.Atoi(InputQtyMat3.Text)
			tempweapon.Quantity[3], _ = strconv.Atoi(InputQtyMat4.Text)

			InsertOneWeapon(client, ctx) //needs to be coded
			w.listUpdateWeapon(app, id)
			wInput.Close()

		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wInput.Close()

		})

		wInput.SetContent(container.New(layout.NewVBoxLayout(), TextWeaponName, InputWeaponName, InputWeaponIcon, TextWeaponKind, InputWeaponKind, TextWeaponRarity, InputWeaponRarity, TextWeaponAttack, InputWeaponAttack, TextWeaponElement, InputWeaponElement, TextWeaponElementvalue, InputWeaponElementvalue, TextWeaponSharpness, InputWeaponSharpness, TextWeaponAffinity, InputWeaponAffinity, TextWeaponDefense, InputWeaponDefense, TextWeaponPrice, InputWeaponPrice, InputWeaponObject, addData, cancel)) //Layout for the "Insertion-Window"
		wInput.Resize(fyne.NewSize(400, 200))
		wInput.CenterOnScreen()
		wInput.Show()
	})

	return container.NewVBox(add)
}

func (w *win) listUpdateWeapon(app fyne.App, id widget.ListItemID) { //function updates materials when another Rank was selected
	Weapons := decodeweapons()
	weaponbuttons := w.weapon_addbutton(app, id)

	Weaponicon := widget.NewIcon(fyne.NewStaticResource("icon", Weapons[id].icon))

	weaponname := widget.NewLabel(Weapons[id].Name)

	updatebutton := w.weapon_updatebutton(app, Weapons[id], id)

	deletebutton := w.weapon_deletebutton(app, Weapons[id], id)

	cancel := widget.NewButton("Cancel", func() {
		w.window.Close()
	})

	choice := container.New(layout.NewVBoxLayout(), weaponname, updatebutton, deletebutton, cancel)

	list, id := initList_Weapon(Weapons, id)

	list.OnSelected = func(id widget.ListItemID) {
		icon := widget.NewIcon(fyne.NewStaticResource("icon", Weapons[id].icon))

		weaponname := widget.NewLabel(Weapons[id].Name)

		updatebutton := w.weapon_updatebutton(app, Weapons[id], id)

		deletebutton := w.weapon_deletebutton(app, Weapons[id], id)

		cancel := widget.NewButton("Cancel", func() {
			w.window.Close()
		})

		choice := container.New(layout.NewVBoxLayout(), weaponname, updatebutton, deletebutton, cancel)
		weaponmaterial := w.weaponmaterial(app, list, Weapons[id])
		gbox := container.New(layout.NewGridLayout(3), list, choice, weaponbuttons, weaponmaterial, icon)
		w.window.SetContent(gbox)
		w.window.Show()
	}

	list.OnUnselected = func(id widget.ListItemID) {
		gbox := container.New(layout.NewGridLayout(3), list, weaponbuttons)
		w.window.SetContent(gbox)
		w.window.Show()
	}

	gbox := container.New(layout.NewGridLayout(3), list, choice, weaponbuttons, Weaponicon)
	w.window.SetContent(gbox)
	w.window.Show()
}

func (w *win) weapon_updatebutton(app fyne.App, Weapon WeaponStruct, id widget.ListItemID) fyne.CanvasObject {

	update := widget.NewButton("Update", func() { //Button to Update Data
		wUpdate := app.NewWindow("Update Data")

		TextWeaponName := canvas.NewText("Weaponname:", color.White)
		TextWeaponKind := canvas.NewText("Type of Weapon:", color.White)
		TextWeaponRarity := canvas.NewText("Rarity:", color.White)
		TextWeaponAttack := canvas.NewText("Attack:", color.White)
		TextWeaponElement := canvas.NewText("Element:", color.White)
		TextWeaponElementvalue := canvas.NewText("Element-Value:", color.White)
		TextWeaponSharpness := canvas.NewText("Sharpness:", color.White)
		TextWeaponAffinity := canvas.NewText("Affinity:", color.White)
		TextWeaponDefense := canvas.NewText("Defense:", color.White)
		TextWeaponPrice := canvas.NewText("Price:", color.White)

		InputWeaponName := widget.NewEntry()
		InputWeaponKind := widget.NewEntry()
		InputWeaponRarity := widget.NewEntry()
		InputWeaponAttack := widget.NewEntry()
		InputWeaponElement := widget.NewEntry()
		InputWeaponElementvalue := widget.NewEntry()
		InputWeaponSharpness := widget.NewEntry()
		InputWeaponAffinity := widget.NewEntry()
		InputWeaponDefense := widget.NewEntry()
		InputWeaponPrice := widget.NewEntry()

		//Object
		InputMaterial1 := widget.NewEntry()
		InputMaterial2 := widget.NewEntry()
		InputMaterial3 := widget.NewEntry()
		InputMaterial4 := widget.NewEntry()

		InputQtyMat1 := widget.NewEntry()
		InputQtyMat2 := widget.NewEntry()
		InputQtyMat3 := widget.NewEntry()
		InputQtyMat4 := widget.NewEntry()

		InputWeaponObject := container.NewGridWithColumns(2,
			container.NewGridWithRows(5,
				widget.NewLabel("Material:"),
				InputMaterial1,
				InputMaterial2,
				InputMaterial3,
				InputMaterial4),

			container.NewGridWithRows(5,
				widget.NewLabel("Quantity:"),
				InputQtyMat1,
				InputQtyMat2,
				InputQtyMat3,
				InputQtyMat4),
		)

		InputWeaponIcon := widget.NewButton("File Open With Filter (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, wUpdate)
					return
				}
				if reader == nil {
					log.Println("Cancelled")
					return
				}

				imageOpenedWeaponIcon(reader)
			}, wUpdate)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()

		})

		updateData := widget.NewButton("Update", func() { //Button to add into ItemName typed Data
			if InputWeaponName.Text != "" {
				tempweapon.Name = InputWeaponName.Text
			}
			if InputWeaponKind.Text != "" {
				tempweapon.Kind = InputWeaponKind.Text
			}
			if InputWeaponRarity.Text != "" {
				tempweapon.Rarity, _ = strconv.Atoi(InputWeaponRarity.Text)
			}
			if InputWeaponAttack.Text != "" {
				tempweapon.Attack, _ = strconv.Atoi(InputWeaponAttack.Text)
			}
			if InputWeaponElement.Text != "" {
				tempweapon.Element = InputWeaponElement.Text
			}
			if InputWeaponElementvalue.Text != "" {
				tempweapon.Elementvalue, _ = strconv.Atoi(InputWeaponElementvalue.Text)
			}
			if InputWeaponSharpness.Text != "" {
				tempweapon.Sharpness = InputWeaponSharpness.Text
			}
			if InputWeaponAffinity.Text != "" {
				tempweapon.Affinity, _ = strconv.Atoi(InputWeaponAffinity.Text)
			}
			if InputWeaponDefense.Text != "" {
				tempweapon.Defense, _ = strconv.Atoi(InputWeaponDefense.Text)
			}
			if InputWeaponPrice.Text != "" {
				tempweapon.Price, _ = strconv.Atoi(InputWeaponPrice.Text)
			}
			//Material
			if InputMaterial1.Text != "" {
				tempweapon.Material[0] = InputMaterial1.Text
			}
			if InputMaterial1.Text != "" {
				tempweapon.Material[1] = InputMaterial2.Text
			}
			if InputMaterial1.Text != "" {
				tempweapon.Material[2] = InputMaterial3.Text
			}
			if InputMaterial1.Text != "" {
				tempweapon.Material[3] = InputMaterial4.Text
			}
			if InputQtyMat1.Text != "" {
				tempweapon.Quantity[0], _ = strconv.Atoi(InputQtyMat1.Text)
			}
			if InputQtyMat1.Text != "" {
				tempweapon.Quantity[1], _ = strconv.Atoi(InputQtyMat2.Text)
			}
			if InputQtyMat1.Text != "" {
				tempweapon.Quantity[2], _ = strconv.Atoi(InputQtyMat3.Text)
			}
			if InputQtyMat1.Text != "" {
				tempweapon.Quantity[3], _ = strconv.Atoi(InputQtyMat4.Text)
			}

			UpdateOneWeapon(client, ctx, Weapon)
			w.listUpdateWeapon(app, id)
			wUpdate.Close()

		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wUpdate.Close()

		})

		wUpdate.SetContent(container.New(layout.NewVBoxLayout(), TextWeaponName, InputWeaponName, InputWeaponIcon, TextWeaponKind, InputWeaponKind, TextWeaponRarity, InputWeaponRarity, TextWeaponAttack, InputWeaponAttack, TextWeaponElement, InputWeaponElement, TextWeaponElementvalue, InputWeaponElementvalue, TextWeaponSharpness, InputWeaponSharpness, TextWeaponAffinity, InputWeaponAffinity, TextWeaponDefense, InputWeaponDefense, TextWeaponPrice, InputWeaponPrice, InputWeaponObject, updateData, cancel)) //Layout for the "Insertion-Window"
		wUpdate.Resize(fyne.NewSize(500, 300))
		wUpdate.CenterOnScreen()
		wUpdate.Show()
	})

	return container.NewVBox(update)
}

func (w *win) weapon_deletebutton(app fyne.App, Weapon WeaponStruct, id widget.ListItemID) fyne.CanvasObject {
	delete := widget.NewButton("Delete", func() { //Button to Delete Items
		DeleteOneWeapon(client, ctx, Weapon)

	})
	return container.NewVBox(delete)
}

func initList_Weapon(Weapons []WeaponStruct, id widget.ListItemID) (*widget.List, widget.ListItemID) {
	list := widget.NewList(
		func() int {
			return len(Weapons)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Weapon", weapon.icon)), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Weapons[id].Name, Weapons[id].icon))
			c.Objects[1].(*widget.Label).SetText(Weapons[id].Name)
			c.Objects[2].(*widget.Label).SetText(Weapons[id].Kind)
			c.Objects[3].(*widget.Label).SetText("Rarity: " + strconv.Itoa((Weapons[id].Rarity)))
			c.Objects[4].(*widget.Label).SetText("Attack: " + strconv.Itoa((Weapons[id].Attack)))
			c.Objects[5].(*widget.Label).SetText("Element: " + (Weapons[id].Element))
			c.Objects[6].(*widget.Label).SetText("Elementvalue: " + strconv.Itoa((Weapons[id].Elementvalue)))
			c.Objects[7].(*widget.Label).SetText("Sharpness: " + (Weapons[id].Sharpness))
			c.Objects[8].(*widget.Label).SetText("Affinity: " + strconv.Itoa((Weapons[id].Affinity)))
			c.Objects[9].(*widget.Label).SetText("Defense: " + strconv.Itoa((Weapons[id].Defense)))
			c.Objects[10].(*widget.Label).SetText("Price: " + strconv.Itoa((Weapons[id].Price)))
		},
	)
	return list, id

}

func (W *win) weaponmaterial(app fyne.App, li *widget.List, Weapon WeaponStruct) fyne.CanvasObject {
	var data = [5][2]string{
		[2]string{"Material", "Quantity"},
		[2]string{Weapon.Material[0], strconv.Itoa(Weapon.Quantity[0])},
		[2]string{Weapon.Material[1], strconv.Itoa(Weapon.Quantity[1])},
		[2]string{Weapon.Material[2], strconv.Itoa(Weapon.Quantity[2])},
		[2]string{Weapon.Material[3], strconv.Itoa(Weapon.Quantity[3])}}

	table := widget.NewTable(
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
