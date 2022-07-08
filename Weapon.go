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

	Weapons := decodeweapons()

	weaponbuttons := w.weapon_addbutton(app)

	list := widget.NewList(
		func() int {
			return len(Weapons)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Weapon", weapon.iconb)), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Weapons[id].Name, Weapons[id].iconb))
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
			c.Objects[11].(*widget.Label).SetText("Material: " + (Weapons[id].Material))
			c.Objects[12].(*widget.Label).SetText("Quantity: " + strconv.Itoa((Weapons[id].Quantity)))
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		icon := widget.NewIcon(fyne.NewStaticResource("icon", Weapons[id].iconb))

		weaponname := widget.NewLabel(Weapons[id].Name)

		updatebutton := w.weapon_updatebutton(app, Weapons[id])

		deletebutton := w.weapon_deletebutton(app, Weapons[id])

		cancel := widget.NewButton("Cancel", func() {
			w.window.Close()
		})
		choice := container.New(layout.NewVBoxLayout(), weaponname, updatebutton, deletebutton, cancel)

		gbox := container.New(layout.NewGridLayout(3), list, choice, weaponbuttons, icon) //display gbox
		w.window.SetContent(gbox)
		w.window.Show()

	}

	list.OnUnselected = func(id widget.ListItemID) {

		gbox := container.New(layout.NewGridLayout(3), list, weaponbuttons) //remove additional widgets
		w.window.SetContent(gbox)                                           //display gbox
		w.window.Show()
	}

	gbox := container.New(layout.NewGridLayout(3), list, weaponbuttons)

	w.window.SetContent(gbox)

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

func (w *win) weapon_addbutton(app fyne.App) fyne.CanvasObject {

	add := widget.NewButton("Add", func() { //Button to Add Data
		wInput := app.NewWindow("Add Data")

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
		TextWeaponMaterial := canvas.NewText("Material:", color.White)
		TextWeaponQuantity := canvas.NewText("Quantity of Material:", color.White)
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
		InputWeaponMaterial := widget.NewEntry()
		InputWeaponQuantity := widget.NewEntry()

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
			tempweapon.Material = InputWeaponMaterial.Text
			tempweapon.Quantity, _ = strconv.Atoi(InputWeaponQuantity.Text)

			//InsertOneWeapon(client, ctx) //needs to be coded
			w.listUpdateWeapon(app)
			wInput.Close()

		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wInput.Close()

		})

		wInput.SetContent(container.New(layout.NewVBoxLayout(), TextWeaponName, InputWeaponName, InputWeaponIcon, TextWeaponKind, InputWeaponKind, TextWeaponRarity, InputWeaponRarity, TextWeaponAttack, InputWeaponAttack, TextWeaponElement, InputWeaponElement, TextWeaponElementvalue, InputWeaponElementvalue, TextWeaponSharpness, InputWeaponSharpness, TextWeaponAffinity, InputWeaponAffinity, TextWeaponDefense, InputWeaponDefense, TextWeaponPrice, InputWeaponPrice, TextWeaponMaterial, InputWeaponMaterial, TextWeaponQuantity, InputWeaponQuantity, addData, cancel)) //Layout for the "Insertion-Window"
		wInput.Resize(fyne.NewSize(500, 300))
		wInput.CenterOnScreen()
		wInput.Show()
	})

	return container.NewVBox(add)
}

func (w *win) listUpdateWeapon(app fyne.App) { //function updates materials when another Rank was selected
	Weapons := decodeweapons()
	weaponbuttons := w.weapon_addbutton(app)

	icon := widget.NewIcon(fyne.NewStaticResource("icon", Weapons[1].iconb))

	list := widget.NewList(
		func() int {
			return len(Weapons)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Weapon", weapon.iconb)), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {

			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Weapons[id].Name, Weapons[id].iconb))
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
			c.Objects[11].(*widget.Label).SetText("Material: " + (Weapons[id].Material))
			c.Objects[12].(*widget.Label).SetText("Quantity: " + strconv.Itoa((Weapons[id].Quantity)))
		},
	)

	gbox := container.New(layout.NewGridLayout(3), list, weaponbuttons, icon)

	w.window.SetContent(gbox)
	w.window.Show()
}

func (w *win) weapon_updatebutton(app fyne.App, Weapon WeaponStruct) fyne.CanvasObject {

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
		TextWeaponMaterial := canvas.NewText("Material:", color.White)
		TextWeaponQuantity := canvas.NewText("Quantity of Material:", color.White)
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
		InputWeaponMaterial := widget.NewEntry()
		InputWeaponQuantity := widget.NewEntry()

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
			tempweapon.Material = InputWeaponMaterial.Text
			tempweapon.Quantity, _ = strconv.Atoi(InputWeaponQuantity.Text)

			//UpdateOneWeapon(client, ctx, Weapon)
			w.listUpdateItem(app)
			wUpdate.Close()

		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wUpdate.Close()

		})

		wUpdate.SetContent(container.New(layout.NewVBoxLayout(), TextWeaponName, InputWeaponName, InputWeaponIcon, TextWeaponKind, InputWeaponKind, TextWeaponRarity, InputWeaponRarity, TextWeaponAttack, InputWeaponAttack, TextWeaponElement, InputWeaponElement, TextWeaponElementvalue, InputWeaponElementvalue, TextWeaponSharpness, InputWeaponSharpness, TextWeaponAffinity, InputWeaponAffinity, TextWeaponDefense, InputWeaponDefense, TextWeaponPrice, InputWeaponPrice, TextWeaponMaterial, InputWeaponMaterial, TextWeaponQuantity, InputWeaponQuantity, updateData, cancel)) //Layout for the "Insertion-Window"
		wUpdate.Resize(fyne.NewSize(500, 300))
		wUpdate.CenterOnScreen()
		wUpdate.Show()
	})

	return container.NewVBox(update)
}

func (w *win) weapon_deletebutton(app fyne.App, Weapon WeaponStruct) fyne.CanvasObject {
	delete := widget.NewButton("Delete", func() { //Button to Delete Items
		//DeleteOne(client, ctx, Weapon)

	})
	return container.NewVBox(delete)
}
