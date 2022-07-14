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

	addbutton := w.weapon_addbutton(app, id)

	list, id := initList_Weapon(Weapons, id)
	list.Resize(fyne.NewSize(25, 25))
	list.OnSelected = func(id widget.ListItemID) {
		//icon := widget.NewIcon(fyne.NewStaticResource("icon", Weapons[id].icon))

		weaponname := widget.NewLabel(Weapons[id].Name)

		updatebutton := w.weapon_updatebutton(app, Weapons[id], id)

		deletebutton := w.weapon_deletebutton(app, Weapons[id], id)

		cancel := widget.NewButton("Cancel", func() {
			w.window.Close()
		})
		buttons := container.New(layout.NewVBoxLayout(), weaponname, addbutton, updatebutton, deletebutton, cancel)

		weaponmaterialForgeLabel := widget.NewLabel("Forge")
		weaponmaterialUpgradeLabel := widget.NewLabel("Upgrade")
		weaponmaterialForge := w.weaponmaterialForge(app, list, Weapons[id])
		weaponmaterialUpgrade := w.weaponmaterialUpgrade(app, list, Weapons[id])
		weaponForgePriceLabel := widget.NewLabel("Price:")
		weaponUpgradePriceLabel := widget.NewLabel("Price:")
		weaponForgePrice := widget.NewLabel(strconv.Itoa(Weapons[id].PriceForge))
		weaponUpgradePrice := widget.NewLabel(strconv.Itoa(Weapons[id].PriceUpgrade))
		weaponForgePriceBox := container.NewHBox(weaponForgePriceLabel, weaponForgePrice)
		weaponUpgradePriceBox := container.NewHBox(weaponUpgradePriceLabel, weaponUpgradePrice)

		weaponmaterialForgeBox := container.NewGridWithRows(3, weaponmaterialForgeLabel, weaponForgePriceBox, weaponmaterialForge)
		weaponmaterialUpgradeBox := container.NewGridWithRows(3, weaponmaterialUpgradeLabel, weaponUpgradePriceBox, weaponmaterialUpgrade)

		gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), buttons, weaponmaterialForgeBox, weaponmaterialUpgradeBox) //display gbox
		w.window.SetContent(gbox)
		w.window.Show()

	}

	list.OnUnselected = func(id widget.ListItemID) {

		gbox := container.New(layout.NewGridLayout(3), list, addbutton) //remove additional widgets
		w.window.SetContent(gbox)
		w.window.Resize(fyne.NewSize(400, 600)) //display gbox
		w.window.Show()
	}

	gbox := container.New(layout.NewGridLayout(3), container.NewHScroll(list), addbutton)

	w.window.SetContent(gbox)
	w.window.Resize(fyne.NewSize(100, 600))
	w.window.Show()
}

func (w *win) weapon_addbutton(app fyne.App, id widget.ListItemID) fyne.CanvasObject {
	var tempweapon TempWeaponStruct
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
		//TextWeaponPrice := canvas.NewText("Price:", color.White)

		InputWeaponName := widget.NewEntry()
		InputWeaponKind := widget.NewEntry()
		InputWeaponRarity := widget.NewEntry()
		InputWeaponAttack := widget.NewEntry()
		InputWeaponElement := widget.NewEntry()
		InputWeaponElementvalue := widget.NewEntry()
		InputWeaponSharpness := widget.NewEntry()
		InputWeaponAffinity := widget.NewEntry()
		InputWeaponDefense := widget.NewEntry()
		//InputWeaponPrice := widget.NewEntry()

		//Material-Forge
		InputMaterialForge1 := widget.NewEntry()
		InputMaterialForge2 := widget.NewEntry()
		InputMaterialForge3 := widget.NewEntry()
		InputMaterialForge4 := widget.NewEntry()

		InputQtyMatForge1 := widget.NewEntry()
		InputQtyMatForge2 := widget.NewEntry()
		InputQtyMatForge3 := widget.NewEntry()
		InputQtyMatForge4 := widget.NewEntry()

		//Material-Upgrade
		InputMaterialUpgrade1 := widget.NewEntry()
		InputMaterialUpgrade2 := widget.NewEntry()
		InputMaterialUpgrade3 := widget.NewEntry()
		InputMaterialUpgrade4 := widget.NewEntry()

		InputQtyMatUpgrade1 := widget.NewEntry()
		InputQtyMatUpgrade2 := widget.NewEntry()
		InputQtyMatUpgrade3 := widget.NewEntry()
		InputQtyMatUpgrade4 := widget.NewEntry()

		InputWeaponMaterialForge := container.NewGridWithColumns(2,
			container.NewGridWithRows(5,
				widget.NewLabel("Material:"),
				InputMaterialForge1,
				InputMaterialForge2,
				InputMaterialForge3,
				InputMaterialForge4),

			container.NewGridWithRows(5,
				widget.NewLabel("Quantity:"),
				InputQtyMatForge1,
				InputQtyMatForge2,
				InputQtyMatForge3,
				InputQtyMatForge4),
		)

		InputWeaponMaterialUpgrade := container.NewGridWithColumns(2,
			container.NewGridWithRows(5,
				widget.NewLabel("Material:"),
				InputMaterialUpgrade1,
				InputMaterialUpgrade2,
				InputMaterialUpgrade3,
				InputMaterialUpgrade4),

			container.NewGridWithRows(5,
				widget.NewLabel("Quantity:"),
				InputQtyMatUpgrade1,
				InputQtyMatUpgrade2,
				InputQtyMatUpgrade3,
				InputQtyMatUpgrade4),
		)

		ForgeLabel := widget.NewLabel("Forge")
		UpgradeLabel := widget.NewLabel("Upgrade")
		ForgePriceLabel := widget.NewLabel("Price:")
		UpgradePriceLabel := widget.NewLabel("Price:")
		ForgePriceEntry := widget.NewEntry()
		UpgradePriceEntry := widget.NewEntry()
		WeaponMaterialLabels := container.NewGridWithColumns(2, ForgeLabel, UpgradeLabel)
		ForgePriceBox := container.NewHBox(ForgePriceLabel, ForgePriceEntry)
		UpgradePriceBox := container.NewHBox(UpgradePriceLabel, UpgradePriceEntry)
		PriceGrid := container.NewGridWithColumns(2, ForgePriceBox, UpgradePriceBox)

		//WeaponPrice := container.NewGridWithColumns(2,ForgePrice,UpgradePrice)
		InputWeaponMaterial := container.NewHSplit(InputWeaponMaterialForge, InputWeaponMaterialUpgrade)
		InputWeaponMaterialBox := container.NewVBox(WeaponMaterialLabels, PriceGrid, InputWeaponMaterial)

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

				tempweapon.EncodedIcon = imageOpenedWeaponIcon(reader, tempweapon)
			}, wInput)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()

		})

		addData := widget.NewButton("Add", func() { //Button to add into ItemName typed Data
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
			/*if InputWeaponPrice.Text != "" {
				tempweapon.PriceForge, _ = strconv.Atoi(InputWeaponPrice.Text)
			}*/

			//Prices
			if ForgePriceEntry.Text != "" {
				tempweapon.PriceForge, _ = strconv.Atoi(ForgePriceEntry.Text)
			}

			if UpgradePriceEntry.Text != "" {
				tempweapon.PriceUpgrade, _ = strconv.Atoi(UpgradePriceEntry.Text)
			}

			//Material-Forge
			if InputMaterialForge1.Text != "" {
				tempweapon.MaterialForge[0] = InputMaterialForge1.Text
			}
			if InputMaterialForge2.Text != "" {
				tempweapon.MaterialForge[1] = InputMaterialForge2.Text
			}
			if InputMaterialForge3.Text != "" {
				tempweapon.MaterialForge[2] = InputMaterialForge3.Text
			}
			if InputMaterialForge4.Text != "" {
				tempweapon.MaterialForge[3] = InputMaterialForge4.Text
			}
			if InputQtyMatForge1.Text != "" {
				tempweapon.QuantityForge[0], _ = strconv.Atoi(InputQtyMatForge1.Text)
			}
			if InputQtyMatForge2.Text != "" {
				tempweapon.QuantityForge[1], _ = strconv.Atoi(InputQtyMatForge2.Text)
			}
			if InputQtyMatForge3.Text != "" {
				tempweapon.QuantityForge[2], _ = strconv.Atoi(InputQtyMatForge3.Text)
			}
			if InputQtyMatForge4.Text != "" {
				tempweapon.QuantityForge[3], _ = strconv.Atoi(InputQtyMatForge4.Text)
			}

			//Material-Upgrade
			if InputMaterialForge1.Text != "" {
				tempweapon.MaterialUpgrade[0] = InputMaterialUpgrade1.Text
			}
			if InputMaterialUpgrade2.Text != "" {
				tempweapon.MaterialUpgrade[1] = InputMaterialUpgrade2.Text
			}
			if InputMaterialUpgrade3.Text != "" {
				tempweapon.MaterialUpgrade[2] = InputMaterialUpgrade3.Text
			}
			if InputMaterialUpgrade4.Text != "" {
				tempweapon.MaterialUpgrade[3] = InputMaterialUpgrade4.Text
			}
			if InputQtyMatUpgrade1.Text != "" {
				tempweapon.QuantityUpgrade[0], _ = strconv.Atoi(InputQtyMatUpgrade1.Text)
			}
			if InputQtyMatUpgrade2.Text != "" {
				tempweapon.QuantityUpgrade[1], _ = strconv.Atoi(InputQtyMatUpgrade2.Text)
			}
			if InputQtyMatUpgrade3.Text != "" {
				tempweapon.QuantityUpgrade[2], _ = strconv.Atoi(InputQtyMatUpgrade3.Text)
			}
			if InputQtyMatUpgrade4.Text != "" {
				tempweapon.QuantityUpgrade[3], _ = strconv.Atoi(InputQtyMatUpgrade4.Text)
			}

			InsertOneWeapon(client, ctx, tempweapon) //needs to be coded
			w.listUpdateWeapon(app, id)
			wInput.Close()

		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wInput.Close()

		})

		wInput.SetContent(container.NewScroll(container.New(layout.NewVBoxLayout(), TextWeaponName, InputWeaponName, InputWeaponIcon, TextWeaponKind, InputWeaponKind, TextWeaponRarity, InputWeaponRarity, TextWeaponAttack, InputWeaponAttack, TextWeaponElement, InputWeaponElement, TextWeaponElementvalue, InputWeaponElementvalue, TextWeaponSharpness, InputWeaponSharpness, TextWeaponAffinity, InputWeaponAffinity, TextWeaponDefense, InputWeaponDefense, InputWeaponMaterialBox, addData, cancel))) //Layout for the "Insertion-Window"
		wInput.Resize(fyne.NewSize(400, 200))
		wInput.CenterOnScreen()
		wInput.Show()
	})

	return container.NewVBox(add)
}

func (w *win) listUpdateWeapon(app fyne.App, id widget.ListItemID) { //function updates materials when another Rank was selected
	Weapons := decodeweapons()
	addbutton := w.weapon_addbutton(app, id)

	//Weaponicon := widget.NewIcon(fyne.NewStaticResource("icon", Weapons[id].icon))

	weaponname := widget.NewLabel(Weapons[id].Name)

	updatebutton := w.weapon_updatebutton(app, Weapons[id], id)

	deletebutton := w.weapon_deletebutton(app, Weapons[id], id)

	cancel := widget.NewButton("Cancel", func() {
		w.window.Close()
	})

	buttons := container.New(layout.NewVBoxLayout(), weaponname, addbutton, updatebutton, deletebutton, cancel)

	list, id := initList_Weapon(Weapons, id)

	weaponmaterialForgeLabel := widget.NewLabel("Forge")
	weaponmaterialUpgradeLabel := widget.NewLabel("Upgrade")
	weaponmaterialForge := w.weaponmaterialForge(app, list, Weapons[id])
	weaponmaterialUpgrade := w.weaponmaterialUpgrade(app, list, Weapons[id])
	weaponForgePriceLabel := widget.NewLabel("Price:")
	weaponUpgradePriceLabel := widget.NewLabel("Price:")
	weaponForgePrice := widget.NewLabel(strconv.Itoa(Weapons[id].PriceForge))
	weaponUpgradePrice := widget.NewLabel(strconv.Itoa(Weapons[id].PriceUpgrade))
	weaponForgePriceBox := container.NewHBox(weaponForgePriceLabel, weaponForgePrice)
	weaponUpgradePriceBox := container.NewHBox(weaponUpgradePriceLabel, weaponUpgradePrice)

	weaponmaterialForgeBox := container.NewGridWithRows(3, weaponmaterialForgeLabel, weaponForgePriceBox, weaponmaterialForge)
	weaponmaterialUpgradeBox := container.NewGridWithRows(3, weaponmaterialUpgradeLabel, weaponUpgradePriceBox, weaponmaterialUpgrade)

	list.OnSelected = func(id widget.ListItemID) {
		//icon := widget.NewIcon(fyne.NewStaticResource("icon", Weapons[id].icon))

		weaponname := widget.NewLabel(Weapons[id].Name)

		updatebutton := w.weapon_updatebutton(app, Weapons[id], id)

		deletebutton := w.weapon_deletebutton(app, Weapons[id], id)

		cancel := widget.NewButton("Cancel", func() {
			w.window.Close()
		})

		buttons := container.New(layout.NewVBoxLayout(), weaponname, addbutton, updatebutton, deletebutton, cancel)

		weaponmaterialForgeLabel := widget.NewLabel("Forge")
		weaponmaterialUpgradeLabel := widget.NewLabel("Upgrade")
		weaponmaterialForge := w.weaponmaterialForge(app, list, Weapons[id])
		weaponmaterialUpgrade := w.weaponmaterialUpgrade(app, list, Weapons[id])
		weaponForgePriceLabel := widget.NewLabel("Price:")
		weaponUpgradePriceLabel := widget.NewLabel("Price:")
		weaponForgePrice := widget.NewLabel(strconv.Itoa(Weapons[id].PriceForge))
		weaponUpgradePrice := widget.NewLabel(strconv.Itoa(Weapons[id].PriceUpgrade))
		weaponForgePriceBox := container.NewHBox(weaponForgePriceLabel, weaponForgePrice)
		weaponUpgradePriceBox := container.NewHBox(weaponUpgradePriceLabel, weaponUpgradePrice)

		weaponmaterialForgeBox := container.NewGridWithRows(3, weaponmaterialForgeLabel, weaponForgePriceBox, weaponmaterialForge)
		weaponmaterialUpgradeBox := container.NewGridWithRows(3, weaponmaterialUpgradeLabel, weaponUpgradePriceBox, weaponmaterialUpgrade)

		gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), buttons, weaponmaterialForgeBox, weaponmaterialUpgradeBox)
		w.window.SetContent(gbox)
		w.window.Show()
	}

	list.OnUnselected = func(id widget.ListItemID) {
		gbox := container.New(layout.NewGridLayout(3), list, addbutton)
		w.window.SetContent(gbox)
		w.window.Show()
	}

	gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), buttons, weaponmaterialForgeBox, weaponmaterialUpgradeBox)
	w.window.SetContent(gbox)
	w.window.Show()
}

func (w *win) weapon_updatebutton(app fyne.App, Weapon WeaponStruct, id widget.ListItemID) fyne.CanvasObject {
	var tempweapon TempWeaponStruct

	tempweapon.Name = Weapon.Name
	tempweapon.Kind = Weapon.Kind
	tempweapon.Rarity = Weapon.Rarity
	tempweapon.Attack = Weapon.Attack
	tempweapon.Element = Weapon.Element
	tempweapon.Elementvalue = Weapon.Elementvalue
	tempweapon.Sharpness = Weapon.Sharpness
	tempweapon.Affinity = Weapon.Affinity
	tempweapon.Defense = Weapon.Defense
	tempweapon.PriceForge = Weapon.PriceForge
	tempweapon.PriceUpgrade = Weapon.PriceUpgrade
	tempweapon.EncodedIcon = Weapon.Encoded

	for i := 0; i <= 3; i++ {
		tempweapon.MaterialForge[i] = Weapon.MaterialForge[i]
		tempweapon.QuantityForge[i] = Weapon.QuantityForge[i]
		tempweapon.MaterialUpgrade[i] = Weapon.MaterialUpgrade[i]
		tempweapon.QuantityUpgrade[i] = Weapon.QuantityUpgrade[i]
	}

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
		//TextWeaponPrice := canvas.NewText("Price:", color.White)

		InputWeaponName := widget.NewEntry()
		InputWeaponKind := widget.NewEntry()
		InputWeaponRarity := widget.NewEntry()
		InputWeaponAttack := widget.NewEntry()
		InputWeaponElement := widget.NewEntry()
		InputWeaponElementvalue := widget.NewEntry()
		InputWeaponSharpness := widget.NewEntry()
		InputWeaponAffinity := widget.NewEntry()
		InputWeaponDefense := widget.NewEntry()
		//InputWeaponPrice := widget.NewEntry()

		//Material-Forge
		InputMaterialForge1 := widget.NewEntry()
		InputMaterialForge2 := widget.NewEntry()
		InputMaterialForge3 := widget.NewEntry()
		InputMaterialForge4 := widget.NewEntry()

		InputQtyMatForge1 := widget.NewEntry()
		InputQtyMatForge2 := widget.NewEntry()
		InputQtyMatForge3 := widget.NewEntry()
		InputQtyMatForge4 := widget.NewEntry()

		//Material-Upgrade
		InputMaterialUpgrade1 := widget.NewEntry()
		InputMaterialUpgrade2 := widget.NewEntry()
		InputMaterialUpgrade3 := widget.NewEntry()
		InputMaterialUpgrade4 := widget.NewEntry()

		InputQtyMatUpgrade1 := widget.NewEntry()
		InputQtyMatUpgrade2 := widget.NewEntry()
		InputQtyMatUpgrade3 := widget.NewEntry()
		InputQtyMatUpgrade4 := widget.NewEntry()

		InputWeaponMaterialForge := container.NewGridWithColumns(2,
			container.NewGridWithRows(5,
				widget.NewLabel("Material:"),
				InputMaterialForge1,
				InputMaterialForge2,
				InputMaterialForge3,
				InputMaterialForge4),

			container.NewGridWithRows(5,
				widget.NewLabel("Quantity:"),
				InputQtyMatForge1,
				InputQtyMatForge2,
				InputQtyMatForge3,
				InputQtyMatForge4),
		)

		InputWeaponMaterialUpgrade := container.NewGridWithColumns(2,
			container.NewGridWithRows(5,
				widget.NewLabel("Material:"),
				InputMaterialUpgrade1,
				InputMaterialUpgrade2,
				InputMaterialUpgrade3,
				InputMaterialUpgrade4),

			container.NewGridWithRows(5,
				widget.NewLabel("Quantity:"),
				InputQtyMatUpgrade1,
				InputQtyMatUpgrade2,
				InputQtyMatUpgrade3,
				InputQtyMatUpgrade4),
		)

		ForgeLabel := widget.NewLabel("Forge")
		UpgradeLabel := widget.NewLabel("Upgrade")
		ForgePriceLabel := widget.NewLabel("Price:")
		UpgradePriceLabel := widget.NewLabel("Price:")
		ForgePriceEntry := widget.NewEntry()
		UpgradePriceEntry := widget.NewEntry()
		WeaponMaterialLabels := container.NewGridWithColumns(2, ForgeLabel, UpgradeLabel)
		ForgePriceBox := container.NewHBox(ForgePriceLabel, ForgePriceEntry)
		UpgradePriceBox := container.NewHBox(UpgradePriceLabel, UpgradePriceEntry)
		PriceGrid := container.NewGridWithColumns(2, ForgePriceBox, UpgradePriceBox)

		//WeaponPrice := container.NewGridWithColumns(2,ForgePrice,UpgradePrice)
		InputWeaponMaterial := container.NewHSplit(InputWeaponMaterialForge, InputWeaponMaterialUpgrade)
		InputWeaponMaterialBox := container.NewVBox(WeaponMaterialLabels, PriceGrid, InputWeaponMaterial)

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

				tempweapon.EncodedIcon = imageOpenedWeaponIcon(reader, tempweapon)
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
			/*if InputWeaponPrice.Text != "" {
				tempweapon.PriceForge, _ = strconv.Atoi(InputWeaponPrice.Text)
			}*/

			//Prices
			if ForgePriceEntry.Text != "" {
				tempweapon.PriceForge, _ = strconv.Atoi(ForgePriceEntry.Text)
			}

			if UpgradePriceEntry.Text != "" {
				tempweapon.PriceUpgrade, _ = strconv.Atoi(UpgradePriceEntry.Text)
			}

			//Material-Forge
			if InputMaterialForge1.Text != "" {
				tempweapon.MaterialForge[0] = InputMaterialForge1.Text
			}
			if InputMaterialForge2.Text != "" {
				tempweapon.MaterialForge[1] = InputMaterialForge2.Text
			}
			if InputMaterialForge3.Text != "" {
				tempweapon.MaterialForge[2] = InputMaterialForge3.Text
			}
			if InputMaterialForge4.Text != "" {
				tempweapon.MaterialForge[3] = InputMaterialForge4.Text
			}
			if InputQtyMatForge1.Text != "" {
				tempweapon.QuantityForge[0], _ = strconv.Atoi(InputQtyMatForge1.Text)
			}
			if InputQtyMatForge2.Text != "" {
				tempweapon.QuantityForge[1], _ = strconv.Atoi(InputQtyMatForge2.Text)
			}
			if InputQtyMatForge3.Text != "" {
				tempweapon.QuantityForge[2], _ = strconv.Atoi(InputQtyMatForge3.Text)
			}
			if InputQtyMatForge4.Text != "" {
				tempweapon.QuantityForge[3], _ = strconv.Atoi(InputQtyMatForge4.Text)
			}

			//Material-Upgrade
			if InputMaterialForge1.Text != "" {
				tempweapon.MaterialUpgrade[0] = InputMaterialUpgrade1.Text
			}
			if InputMaterialUpgrade2.Text != "" {
				tempweapon.MaterialUpgrade[1] = InputMaterialUpgrade2.Text
			}
			if InputMaterialUpgrade3.Text != "" {
				tempweapon.MaterialUpgrade[2] = InputMaterialUpgrade3.Text
			}
			if InputMaterialUpgrade4.Text != "" {
				tempweapon.MaterialUpgrade[3] = InputMaterialUpgrade4.Text
			}
			if InputQtyMatUpgrade1.Text != "" {
				tempweapon.QuantityUpgrade[0], _ = strconv.Atoi(InputQtyMatUpgrade1.Text)
			}
			if InputQtyMatUpgrade2.Text != "" {
				tempweapon.QuantityUpgrade[1], _ = strconv.Atoi(InputQtyMatUpgrade2.Text)
			}
			if InputQtyMatUpgrade3.Text != "" {
				tempweapon.QuantityUpgrade[2], _ = strconv.Atoi(InputQtyMatUpgrade3.Text)
			}
			if InputQtyMatUpgrade4.Text != "" {
				tempweapon.QuantityUpgrade[3], _ = strconv.Atoi(InputQtyMatUpgrade4.Text)
			}

			UpdateOneWeapon(client, ctx, Weapon, tempweapon)
			w.listUpdateWeapon(app, id)
			wUpdate.Close()

		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wUpdate.Close()

		})

		wUpdate.SetContent(container.NewScroll(container.New(layout.NewVBoxLayout(), TextWeaponName, InputWeaponName, InputWeaponIcon, TextWeaponKind, InputWeaponKind, TextWeaponRarity, InputWeaponRarity, TextWeaponAttack, InputWeaponAttack, TextWeaponElement, InputWeaponElement, TextWeaponElementvalue, InputWeaponElementvalue, TextWeaponSharpness, InputWeaponSharpness, TextWeaponAffinity, InputWeaponAffinity, TextWeaponDefense, InputWeaponDefense, InputWeaponMaterialBox, updateData, cancel))) //Layout for the "Insertion-Window"
		wUpdate.Resize(fyne.NewSize(500, 300))
		wUpdate.CenterOnScreen()
		wUpdate.Show()
	})

	return container.NewVBox(update)
}

func (w *win) weapon_deletebutton(app fyne.App, Weapon WeaponStruct, id widget.ListItemID) fyne.CanvasObject {
	delete := widget.NewButton("Delete", func() { //Button to Delete Items
		DeleteOneWeapon(client, ctx, Weapon)
		id = id - 1
		if id >= 0 {
			w.listUpdateItem(app, id)
		}
		if id < 0 {
			id = 0 //set id back to 0 , array out of bounds otherwise -> -1
			list := initemptylist()
			addbutton := w.item_addbutton(app, id)
			gbox := container.NewGridWithColumns(3, list, addbutton)
			w.window.SetContent(gbox)
			w.window.Show()
		}

	})
	return container.NewVBox(delete)
}

func initList_Weapon(Weapons []WeaponStruct, id widget.ListItemID) (*widget.List, widget.ListItemID) {
	list := widget.NewList(
		func() int {
			return len(Weapons)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Weapon", weapon.icon)), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"))
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

		},
	)
	return list, id

}

func (W *win) weaponmaterialForge(app fyne.App, li *widget.List, Weapon WeaponStruct) fyne.CanvasObject {
	var data = [5][2]string{
		[2]string{"Material", "Quantity"},
		[2]string{Weapon.MaterialForge[0], strconv.Itoa(Weapon.QuantityForge[0])},
		[2]string{Weapon.MaterialForge[1], strconv.Itoa(Weapon.QuantityForge[1])},
		[2]string{Weapon.MaterialForge[2], strconv.Itoa(Weapon.QuantityForge[2])},
		[2]string{Weapon.MaterialForge[3], strconv.Itoa(Weapon.QuantityForge[3])}}

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

func (W *win) weaponmaterialUpgrade(app fyne.App, li *widget.List, Weapon WeaponStruct) fyne.CanvasObject {
	var data = [5][2]string{
		[2]string{"Material", "Quantity"},
		[2]string{Weapon.MaterialUpgrade[0], strconv.Itoa(Weapon.QuantityUpgrade[0])},
		[2]string{Weapon.MaterialUpgrade[1], strconv.Itoa(Weapon.QuantityUpgrade[1])},
		[2]string{Weapon.MaterialUpgrade[2], strconv.Itoa(Weapon.QuantityUpgrade[2])},
		[2]string{Weapon.MaterialUpgrade[3], strconv.Itoa(Weapon.QuantityUpgrade[3])}}

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
