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
	"fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (w *win) ArmorUI(app fyne.App) {
	w.window = app.NewWindow("Armors")
	var id widget.ListItemID
	Armors := decodearmors()

	addbutton := w.armor_addbutton(app, id)

	list, id := initList_Armors(Armors, id)
	//list.Resize(fyne.NewSize(25, 25))
	list.OnSelected = func(id widget.ListItemID) {
		//icon := widget.NewIcon(fyne.NewStaticResource("icon", Weapons[id].icon))

		weaponname := widget.NewLabel(Armors[id].Name)

		updatebutton := w.armor_updatebutton(app, Armors[id], id)

		deletebutton := w.armor_deletebutton(app, Armors[id], id)

		cancel := widget.NewButton("Cancel", func() {
			w.window.Close()
		})
		buttons := container.New(layout.NewVBoxLayout(), weaponname, addbutton, updatebutton, deletebutton, cancel)

		weaponmaterialForgeLabel := widget.NewLabel("Forge")
		weaponmaterialUpgradeLabel := widget.NewLabel("Upgrade")
		weaponmaterialForge := w.armormaterialForge(app, list, Armors[id])
		weaponmaterialUpgrade := w.armormaterialUpgrade(app, list, Armors[id])
		weaponForgePriceLabel := widget.NewLabel("Price:")
		weaponUpgradePriceLabel := widget.NewLabel("Price:")
		weaponForgePrice := widget.NewLabel(strconv.Itoa(Armors[id].PriceForge))
		weaponUpgradePrice := widget.NewLabel(strconv.Itoa(Armors[id].PriceUpgrade))
		weaponForgePriceBox := container.NewHBox(weaponForgePriceLabel, weaponForgePrice)
		weaponUpgradePriceBox := container.NewHBox(weaponUpgradePriceLabel, weaponUpgradePrice)

		weaponmaterialForgeBox := container.NewGridWithRows(3, weaponmaterialForgeLabel, weaponForgePriceBox, weaponmaterialForge)
		weaponmaterialUpgradeBox := container.NewGridWithRows(3, weaponmaterialUpgradeLabel, weaponUpgradePriceBox, weaponmaterialUpgrade)

		gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), buttons, weaponmaterialForgeBox, weaponmaterialUpgradeBox) //display gbox
		w.window.SetContent(gbox)
		w.window.Show()

	}

	list.OnUnselected = func(id widget.ListItemID) {

		gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), addbutton) //remove additional widgets
		w.window.SetContent(gbox)                                                             //display gbox

		w.window.Show()
	}

	gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), addbutton)

	w.window.SetContent(gbox)
	w.window.Resize(fyne.NewSize(100, 600))
	w.window.Show()
}

func (w *win) armor_addbutton(app fyne.App, id widget.ListItemID) fyne.CanvasObject {
	var temparmor TempArmorStruct
	add := widget.NewButton("Add", func() { //Button to Add Data
		wInput := app.NewWindow("Add Data")

		TextArmorName := canvas.NewText("Armorname:", color.White)
		TextArmorKind := canvas.NewText("Type of Armor:", color.White)
		TextArmorRarity := canvas.NewText("Rarity:", color.White)
		TextArmorDefense := canvas.NewText("Defense:", color.White)
		TextArmorFireRes := canvas.NewText("Fire-Res:", color.White)
		TextArmorThunderRes := canvas.NewText("Thunder-Res:", color.White)
		TextArmorWaterRes := canvas.NewText("Water-Res:", color.White)
		TextArmorIceRes := canvas.NewText("Ice-Res:", color.White)
		TextArmorDragonRes := canvas.NewText("Dragon-Res:", color.White)

		InputArmorName := widget.NewEntry()
		InputArmorKind := widget.NewEntry()
		InputArmorRarity := widget.NewEntry()
		InputArmorDefense := widget.NewEntry()
		InputArmorFireRes := widget.NewEntry()
		InputArmorThunderRes := widget.NewEntry()
		InputArmorWaterRes := widget.NewEntry()
		InputArmorIceRes := widget.NewEntry()
		InputArmorDragonRes := widget.NewEntry()

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

		InputArmorMaterialForge := container.NewGridWithColumns(2,
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

		InputArmorMaterialUpgrade := container.NewGridWithColumns(2,
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
		ArmorMaterialLabels := container.NewGridWithColumns(2, ForgeLabel, UpgradeLabel)
		ForgePriceBox := container.NewHBox(ForgePriceLabel, ForgePriceEntry)
		UpgradePriceBox := container.NewHBox(UpgradePriceLabel, UpgradePriceEntry)
		PriceGrid := container.NewGridWithColumns(2, ForgePriceBox, UpgradePriceBox)

		//WeaponPrice := container.NewGridWithColumns(2,ForgePrice,UpgradePrice)
		InputArmorMaterial := container.NewHSplit(InputArmorMaterialForge, InputArmorMaterialUpgrade)
		InputArmorMaterialBox := container.NewVBox(ArmorMaterialLabels, PriceGrid, InputArmorMaterial)

		InputArmorIcon := widget.NewButton("File Open With Filter (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, wInput)
					return
				}
				if reader == nil {
					log.Println("Cancelled")
					return
				}

				temparmor.EncodedIcon = imageOpenedArmorIcon(reader, temparmor)
			}, wInput)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()

		})

		addData := widget.NewButton("Add", func() { //Button to add into ItemName typed Data
			if InputArmorName.Text != "" {
				temparmor.Name = InputArmorName.Text
			}
			if InputArmorKind.Text != "" {
				temparmor.Kind = InputArmorKind.Text
			}
			if InputArmorRarity.Text != "" {
				temparmor.Rarity, _ = strconv.Atoi(InputArmorRarity.Text)
			}
			if InputArmorDefense.Text != "" {
				temparmor.Defense, _ = strconv.Atoi(InputArmorDefense.Text)
			}
			if InputArmorFireRes.Text != "" {
				temparmor.FireRes, _ = strconv.Atoi(InputArmorFireRes.Text)
			}
			if InputArmorThunderRes.Text != "" {
				temparmor.ThunderRes, _ = strconv.Atoi(InputArmorThunderRes.Text)
			}
			if InputArmorWaterRes.Text != "" {
				temparmor.WaterRes, _ = strconv.Atoi(InputArmorWaterRes.Text)
			}
			if InputArmorIceRes.Text != "" {
				temparmor.IceRes, _ = strconv.Atoi(InputArmorIceRes.Text)
			}
			if InputArmorDragonRes.Text != "" {
				temparmor.DragonRes, _ = strconv.Atoi(InputArmorDragonRes.Text)
			}

			//Prices
			if ForgePriceEntry.Text != "" {
				temparmor.PriceForge, _ = strconv.Atoi(ForgePriceEntry.Text)
			}

			if UpgradePriceEntry.Text != "" {
				temparmor.PriceUpgrade, _ = strconv.Atoi(UpgradePriceEntry.Text)
			}

			//Material-Forge
			if InputMaterialForge1.Text != "" {
				temparmor.MaterialForge[0] = InputMaterialForge1.Text
			}
			if InputMaterialForge2.Text != "" {
				temparmor.MaterialForge[1] = InputMaterialForge2.Text
			}
			if InputMaterialForge3.Text != "" {
				temparmor.MaterialForge[2] = InputMaterialForge3.Text
			}
			if InputMaterialForge4.Text != "" {
				temparmor.MaterialForge[3] = InputMaterialForge4.Text
			}
			if InputQtyMatForge1.Text != "" {
				temparmor.QuantityForge[0], _ = strconv.Atoi(InputQtyMatForge1.Text)
			}
			if InputQtyMatForge2.Text != "" {
				temparmor.QuantityForge[1], _ = strconv.Atoi(InputQtyMatForge2.Text)
			}
			if InputQtyMatForge3.Text != "" {
				temparmor.QuantityForge[2], _ = strconv.Atoi(InputQtyMatForge3.Text)
			}
			if InputQtyMatForge4.Text != "" {
				temparmor.QuantityForge[3], _ = strconv.Atoi(InputQtyMatForge4.Text)
			}

			//Material-Upgrade
			if InputMaterialForge1.Text != "" {
				temparmor.MaterialUpgrade[0] = InputMaterialUpgrade1.Text
			}
			if InputMaterialUpgrade2.Text != "" {
				temparmor.MaterialUpgrade[1] = InputMaterialUpgrade2.Text
			}
			if InputMaterialUpgrade3.Text != "" {
				temparmor.MaterialUpgrade[2] = InputMaterialUpgrade3.Text
			}
			if InputMaterialUpgrade4.Text != "" {
				temparmor.MaterialUpgrade[3] = InputMaterialUpgrade4.Text
			}
			if InputQtyMatUpgrade1.Text != "" {
				temparmor.QuantityUpgrade[0], _ = strconv.Atoi(InputQtyMatUpgrade1.Text)
			}
			if InputQtyMatUpgrade2.Text != "" {
				temparmor.QuantityUpgrade[1], _ = strconv.Atoi(InputQtyMatUpgrade2.Text)
			}
			if InputQtyMatUpgrade3.Text != "" {
				temparmor.QuantityUpgrade[2], _ = strconv.Atoi(InputQtyMatUpgrade3.Text)
			}
			if InputQtyMatUpgrade4.Text != "" {
				temparmor.QuantityUpgrade[3], _ = strconv.Atoi(InputQtyMatUpgrade4.Text)
			}

			InsertOneArmor(client, ctx, temparmor) //needs to be coded
			w.listUpdateArmors(app, id)
			wInput.Close()

		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wInput.Close()

		})

		wInput.SetContent(container.NewScroll(container.New(layout.NewVBoxLayout(), TextArmorName, InputArmorName, InputArmorIcon, TextArmorKind, InputArmorKind, TextArmorRarity, InputArmorRarity, TextArmorDefense, InputArmorDefense, TextArmorFireRes, InputArmorFireRes, TextArmorThunderRes, InputArmorThunderRes, TextArmorWaterRes, InputArmorWaterRes, TextArmorIceRes, InputArmorIceRes, TextArmorDragonRes, InputArmorDragonRes, InputArmorMaterialBox, addData, cancel))) //Layout for the "Insertion-Window"
		wInput.Resize(fyne.NewSize(400, 200))
		wInput.CenterOnScreen()
		wInput.Show()
	})

	return container.NewVBox(add)
}

func (w *win) listUpdateArmors(app fyne.App, id widget.ListItemID) {
	Armors := decodearmors()
	addbutton := w.armor_addbutton(app, id)

	armorname := widget.NewLabel(Armors[id].Name)

	updatebutton := w.armor_updatebutton(app, Armors[id], id)

	deletebutton := w.armor_deletebutton(app, Armors[id], id)

	cancel := widget.NewButton("Cancel", func() {
		w.window.Close()
	})

	buttons := container.New(layout.NewVBoxLayout(), armorname, addbutton, updatebutton, deletebutton, cancel)

	list, id := initList_Armors(Armors, id)

	armormaterialForgeLabel := widget.NewLabel("Forge")
	armormaterialUpgradeLabel := widget.NewLabel("Upgrade")
	armormaterialForge := w.armormaterialForge(app, list, Armors[id])
	armormaterialUpgrade := w.armormaterialUpgrade(app, list, Armors[id])
	armorForgePriceLabel := widget.NewLabel("Price:")
	armorUpgradePriceLabel := widget.NewLabel("Price:")
	armorForgePrice := widget.NewLabel(strconv.Itoa(Armors[id].PriceForge))
	armorUpgradePrice := widget.NewLabel(strconv.Itoa(Armors[id].PriceUpgrade))
	armorForgePriceBox := container.NewHBox(armorForgePriceLabel, armorForgePrice)
	armorUpgradePriceBox := container.NewHBox(armorUpgradePriceLabel, armorUpgradePrice)

	armormaterialForgeBox := container.NewGridWithRows(3, armormaterialForgeLabel, armorForgePriceBox, armormaterialForge)
	armormaterialUpgradeBox := container.NewGridWithRows(3, armormaterialUpgradeLabel, armorUpgradePriceBox, armormaterialUpgrade)

	list.OnSelected = func(id widget.ListItemID) {

		armorname := widget.NewLabel(Armors[id].Name)

		updatebutton := w.armor_updatebutton(app, Armors[id], id)

		deletebutton := w.armor_deletebutton(app, Armors[id], id)

		cancel := widget.NewButton("Cancel", func() {
			w.window.Close()
		})

		buttons := container.New(layout.NewVBoxLayout(), armorname, addbutton, updatebutton, deletebutton, cancel)

		armormaterialForgeLabel := widget.NewLabel("Forge")
		armormaterialUpgradeLabel := widget.NewLabel("Upgrade")
		armormaterialForge := w.armormaterialForge(app, list, Armors[id])
		armormaterialUpgrade := w.armormaterialUpgrade(app, list, Armors[id])
		armorForgePriceLabel := widget.NewLabel("Price:")
		armorUpgradePriceLabel := widget.NewLabel("Price:")
		armorForgePrice := widget.NewLabel(strconv.Itoa(Armors[id].PriceForge))
		armorUpgradePrice := widget.NewLabel(strconv.Itoa(Armors[id].PriceUpgrade))
		armorForgePriceBox := container.NewHBox(armorForgePriceLabel, armorForgePrice)
		armorUpgradePriceBox := container.NewHBox(armorUpgradePriceLabel, armorUpgradePrice)

		armormaterialForgeBox := container.NewGridWithRows(3, armormaterialForgeLabel, armorForgePriceBox, armormaterialForge)
		armormaterialUpgradeBox := container.NewGridWithRows(3, armormaterialUpgradeLabel, armorUpgradePriceBox, armormaterialUpgrade)

		gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), buttons, armormaterialForgeBox, armormaterialUpgradeBox)
		w.window.SetContent(gbox)
		w.window.Show()
	}

	list.OnUnselected = func(id widget.ListItemID) {
		gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), addbutton)
		w.window.SetContent(gbox)
		w.window.Show()
	}

	gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), buttons, armormaterialForgeBox, armormaterialUpgradeBox)
	w.window.SetContent(gbox)
	w.window.Show()
}

func (w *win) armor_updatebutton(app fyne.App, Armor ArmorStruct, id widget.ListItemID) fyne.CanvasObject {
	var temparmor TempArmorStruct

	temparmor.Name = Armor.Name
	temparmor.Kind = Armor.Kind
	temparmor.Rarity = Armor.Rarity
	temparmor.Defense = Armor.Defense
	temparmor.FireRes = Armor.FireRes
	temparmor.ThunderRes = Armor.ThunderRes
	temparmor.WaterRes = Armor.WaterRes
	temparmor.IceRes = Armor.IceRes
	temparmor.DragonRes = Armor.DragonRes
	temparmor.PriceForge = Armor.PriceForge
	temparmor.PriceUpgrade = Armor.PriceUpgrade
	temparmor.EncodedIcon = Armor.EncodedIcon

	for i := 0; i <= 3; i++ {
		temparmor.MaterialForge[i] = Armor.MaterialForge[i]
		temparmor.QuantityForge[i] = Armor.QuantityForge[i]
		temparmor.MaterialUpgrade[i] = Armor.MaterialUpgrade[i]
		temparmor.QuantityUpgrade[i] = Armor.QuantityUpgrade[i]
	}

	update := widget.NewButton("Update", func() { //Button to Update Data
		wUpdate := app.NewWindow("Update Data")

		TextArmorName := canvas.NewText("Armorname:", color.White)
		TextArmorKind := canvas.NewText("Type of Armor:", color.White)
		TextArmorRarity := canvas.NewText("Rarity:", color.White)
		TextArmorDefense := canvas.NewText("Defense:", color.White)
		TextArmorFireRes := canvas.NewText("Fire-Res:", color.White)
		TextArmorThunderRes := canvas.NewText("Thunder-Res:", color.White)
		TextArmorWaterRes := canvas.NewText("Water-Res:", color.White)
		TextArmorIceRes := canvas.NewText("Ice-Res:", color.White)
		TextArmorDragonRes := canvas.NewText("Dragon-Res:", color.White)

		InputArmorName := widget.NewEntry()
		InputArmorKind := widget.NewEntry()
		InputArmorRarity := widget.NewEntry()
		InputArmorDefense := widget.NewEntry()
		InputArmorFireRes := widget.NewEntry()
		InputArmorThunderRes := widget.NewEntry()
		InputArmorWaterRes := widget.NewEntry()
		InputArmorIceRes := widget.NewEntry()
		InputArmorDragonRes := widget.NewEntry()

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

		InputArmorMaterialForge := container.NewGridWithColumns(2,
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

		InputArmorMaterialUpgrade := container.NewGridWithColumns(2,
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
		ArmorMaterialLabels := container.NewGridWithColumns(2, ForgeLabel, UpgradeLabel)
		ForgePriceBox := container.NewHBox(ForgePriceLabel, ForgePriceEntry)
		UpgradePriceBox := container.NewHBox(UpgradePriceLabel, UpgradePriceEntry)
		PriceGrid := container.NewGridWithColumns(2, ForgePriceBox, UpgradePriceBox)

		InputArmorMaterial := container.NewHSplit(InputArmorMaterialForge, InputArmorMaterialUpgrade)
		InputArmorMaterialBox := container.NewVBox(ArmorMaterialLabels, PriceGrid, InputArmorMaterial)

		InputArmorIcon := widget.NewButton("File Open With Filter (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, wUpdate)
					return
				}
				if reader == nil {
					log.Println("Cancelled")
					return
				}

				temparmor.EncodedIcon = imageOpenedArmorIcon(reader, temparmor)
			}, wUpdate)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()

		})

		updateData := widget.NewButton("Update", func() { //Button to add into ItemName typed Data
			if InputArmorName.Text != "" {
				temparmor.Name = InputArmorName.Text
			}
			if InputArmorKind.Text != "" {
				temparmor.Kind = InputArmorKind.Text
			}
			if InputArmorRarity.Text != "" {
				temparmor.Rarity, _ = strconv.Atoi(InputArmorRarity.Text)
			}
			if InputArmorDefense.Text != "" {
				temparmor.Defense, _ = strconv.Atoi(InputArmorDefense.Text)
			}
			if InputArmorFireRes.Text != "" {
				temparmor.FireRes, _ = strconv.Atoi(InputArmorFireRes.Text)
			}
			if InputArmorThunderRes.Text != "" {
				temparmor.ThunderRes, _ = strconv.Atoi(InputArmorThunderRes.Text)
			}
			if InputArmorWaterRes.Text != "" {
				temparmor.WaterRes, _ = strconv.Atoi(InputArmorWaterRes.Text)
			}
			if InputArmorIceRes.Text != "" {
				temparmor.IceRes, _ = strconv.Atoi(InputArmorIceRes.Text)
			}
			if InputArmorDragonRes.Text != "" {
				temparmor.DragonRes, _ = strconv.Atoi(InputArmorDragonRes.Text)
			}

			//Prices
			if ForgePriceEntry.Text != "" {
				temparmor.PriceForge, _ = strconv.Atoi(ForgePriceEntry.Text)
			}

			if UpgradePriceEntry.Text != "" {
				temparmor.PriceUpgrade, _ = strconv.Atoi(UpgradePriceEntry.Text)
			}

			//Material-Forge
			if InputMaterialForge1.Text != "" {
				temparmor.MaterialForge[0] = InputMaterialForge1.Text
			}
			if InputMaterialForge2.Text != "" {
				temparmor.MaterialForge[1] = InputMaterialForge2.Text
			}
			if InputMaterialForge3.Text != "" {
				temparmor.MaterialForge[2] = InputMaterialForge3.Text
			}
			if InputMaterialForge4.Text != "" {
				temparmor.MaterialForge[3] = InputMaterialForge4.Text
			}
			if InputQtyMatForge1.Text != "" {
				temparmor.QuantityForge[0], _ = strconv.Atoi(InputQtyMatForge1.Text)
			}
			if InputQtyMatForge2.Text != "" {
				temparmor.QuantityForge[1], _ = strconv.Atoi(InputQtyMatForge2.Text)
			}
			if InputQtyMatForge3.Text != "" {
				temparmor.QuantityForge[2], _ = strconv.Atoi(InputQtyMatForge3.Text)
			}
			if InputQtyMatForge4.Text != "" {
				temparmor.QuantityForge[3], _ = strconv.Atoi(InputQtyMatForge4.Text)
			}

			//Material-Upgrade
			if InputMaterialUpgrade1.Text != "" {
				temparmor.MaterialUpgrade[0] = InputMaterialUpgrade1.Text
			}
			if InputMaterialUpgrade2.Text != "" {
				temparmor.MaterialUpgrade[1] = InputMaterialUpgrade2.Text
			}
			if InputMaterialUpgrade3.Text != "" {
				temparmor.MaterialUpgrade[2] = InputMaterialUpgrade3.Text
			}
			if InputMaterialUpgrade4.Text != "" {
				temparmor.MaterialUpgrade[3] = InputMaterialUpgrade4.Text
			}
			if InputQtyMatUpgrade1.Text != "" {
				temparmor.QuantityUpgrade[0], _ = strconv.Atoi(InputQtyMatUpgrade1.Text)
			}
			if InputQtyMatUpgrade2.Text != "" {
				temparmor.QuantityUpgrade[1], _ = strconv.Atoi(InputQtyMatUpgrade2.Text)
			}
			if InputQtyMatUpgrade3.Text != "" {
				temparmor.QuantityUpgrade[2], _ = strconv.Atoi(InputQtyMatUpgrade3.Text)
			}
			if InputQtyMatUpgrade4.Text != "" {
				temparmor.QuantityUpgrade[3], _ = strconv.Atoi(InputQtyMatUpgrade4.Text)
			}

			UpdateOneArmor(client, ctx, Armor, temparmor)
			w.listUpdateArmors(app, id)
			wUpdate.Close()

		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wUpdate.Close()

		})

		wUpdate.SetContent(container.NewScroll(container.New(layout.NewVBoxLayout(), TextArmorName, InputArmorName, InputArmorIcon, TextArmorKind, InputArmorKind, TextArmorRarity, InputArmorRarity, TextArmorDefense, InputArmorDefense, TextArmorFireRes, InputArmorFireRes, TextArmorThunderRes, InputArmorThunderRes, TextArmorWaterRes, InputArmorWaterRes, TextArmorIceRes, InputArmorIceRes, TextArmorDragonRes, InputArmorDragonRes, InputArmorMaterialBox, updateData, cancel))) //Layout for the "Insertion-Window"
		wUpdate.Resize(fyne.NewSize(500, 300))
		wUpdate.CenterOnScreen()
		wUpdate.Show()
	})

	return container.NewVBox(update)
}

func (w *win) armor_deletebutton(app fyne.App, Armor ArmorStruct, id widget.ListItemID) fyne.CanvasObject {
	delete := widget.NewButton("Delete", func() { //Button to Delete Items
		DeleteOneArmors(client, ctx, Armor)
		id = id - 1
		if id >= 0 {
			w.listUpdateArmors(app, id)
		}
		if id < 0 {
			id = 0 //set id back to 0 , array out of bounds otherwise -> -1
			list := initemptylist()
			addbutton := w.armor_addbutton(app, id)
			gbox := container.NewGridWithColumns(2, container.NewHScroll(list), addbutton)
			w.window.SetContent(gbox)
			w.window.Show()
		}

	})
	return container.NewVBox(delete)
}

func initList_Armors(Armors []ArmorStruct, id widget.ListItemID) (*widget.List, widget.ListItemID) {
	list := widget.NewList(
		func() int {
			return len(Armors)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Armor", theme.AccountIcon().Content())), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Armors[id].Name, Armors[id].icon))
			c.Objects[1].(*widget.Label).SetText(Armors[id].Name)
			c.Objects[2].(*widget.Label).SetText(Armors[id].Kind)
			c.Objects[3].(*widget.Label).SetText("Rarity: " + strconv.Itoa((Armors[id].Rarity)))
			c.Objects[4].(*widget.Label).SetText("Defense: " + strconv.Itoa((Armors[id].Defense)))
			c.Objects[5].(*widget.Label).SetText("FireRes: " + strconv.Itoa(Armors[id].FireRes))
			c.Objects[6].(*widget.Label).SetText("ThunderRes: " + strconv.Itoa((Armors[id].ThunderRes)))
			c.Objects[7].(*widget.Label).SetText("WaterRes: " + strconv.Itoa(Armors[id].WaterRes))
			c.Objects[8].(*widget.Label).SetText("IceRes: " + strconv.Itoa((Armors[id].IceRes)))
			c.Objects[9].(*widget.Label).SetText("DragonRes: " + strconv.Itoa((Armors[id].DragonRes)))

		},
	)
	return list, id

}

func (W *win) armormaterialForge(app fyne.App, li *widget.List, Armor ArmorStruct) fyne.CanvasObject {
	var data = [5][2]string{
		[2]string{"Material", "Quantity"},
		[2]string{Armor.MaterialForge[0], strconv.Itoa(Armor.QuantityForge[0])},
		[2]string{Armor.MaterialForge[1], strconv.Itoa(Armor.QuantityForge[1])},
		[2]string{Armor.MaterialForge[2], strconv.Itoa(Armor.QuantityForge[2])},
		[2]string{Armor.MaterialForge[3], strconv.Itoa(Armor.QuantityForge[3])}}

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

func (W *win) armormaterialUpgrade(app fyne.App, li *widget.List, Armor ArmorStruct) fyne.CanvasObject {
	var data = [5][2]string{
		[2]string{"Material", "Quantity"},
		[2]string{Armor.MaterialUpgrade[0], strconv.Itoa(Armor.QuantityUpgrade[0])},
		[2]string{Armor.MaterialUpgrade[1], strconv.Itoa(Armor.QuantityUpgrade[1])},
		[2]string{Armor.MaterialUpgrade[2], strconv.Itoa(Armor.QuantityUpgrade[2])},
		[2]string{Armor.MaterialUpgrade[3], strconv.Itoa(Armor.QuantityUpgrade[3])}}

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
