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

func (w *win) ItemUI(app fyne.App) {
	w.window = app.NewWindow("Item")
	var id widget.ListItemID
	Items := decodeitems()

	addbutton := w.item_addbutton(app, id)

	list, id := initList_Item(Items, id)
	list.Resize(fyne.NewSize(25, 25))

	list.OnSelected = func(id widget.ListItemID) {
		icon := widget.NewIcon(fyne.NewStaticResource("icon", Items[id].icon))

		itemname := widget.NewLabel(Items[id].Name)

		updatebutton := w.item_updatebutton(app, Items[id], id)

		deletebutton := w.item_deletebutton(app, Items[id], id)

		cancel := widget.NewButton("Cancel", func() {
			w.window.Close()
		})
		buttons := container.New(layout.NewVBoxLayout(), itemname, addbutton, updatebutton, deletebutton, cancel)

		gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), buttons, icon) //display gbox
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

	w.window.Show()
}

func (w *win) item_addbutton(app fyne.App, id widget.ListItemID) fyne.CanvasObject {
	var tempitem TempItemStruct
	add := widget.NewButton("Add", func() { //Button to Add Data
		wInput := app.NewWindow("Add Data")

		TextItemName := canvas.NewText("Itemname:", color.White)
		TextItemRarity := canvas.NewText("Rarity:", color.White)
		TextItemQty := canvas.NewText("Quantity:", color.White)
		TextItemSell := canvas.NewText("Sell:", color.White)
		TextItemBuy := canvas.NewText("Buy:", color.White)
		InputItemName := widget.NewEntry()
		InputItemRarity := widget.NewEntry()
		InputItemQty := widget.NewEntry()
		InputItemSell := widget.NewEntry()
		InputItemBuy := widget.NewEntry()

		InputItemIcon := widget.NewButton("Choose Item-Icon (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, wInput)
					return
				}
				if reader == nil {
					log.Println("Cancelled")
					return
				}

				imageOpenedItemIcon(reader, tempitem)
			}, wInput)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()

		})

		addData := widget.NewButton("Add", func() { //Button to add into ItemName typed Data
			tempitem.Name = InputItemName.Text
			tempitem.Rarity, _ = strconv.Atoi(InputItemRarity.Text)
			tempitem.Qty, _ = strconv.Atoi(InputItemQty.Text)
			tempitem.Sell, _ = strconv.Atoi(InputItemSell.Text)
			tempitem.Buy, _ = strconv.Atoi(InputItemBuy.Text)

			InsertOneItem(client, ctx, tempitem)
			w.listUpdateItem(app, id)
			wInput.Close()

		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wInput.Close()

		})

		wInput.SetContent(container.NewScroll(container.New(layout.NewVBoxLayout(), TextItemName, InputItemName, InputItemIcon, TextItemRarity, InputItemRarity, TextItemQty, InputItemQty, TextItemSell, InputItemSell, TextItemBuy, InputItemBuy, addData, cancel))) //Layout for the "Insertion-Window"
		wInput.Resize(fyne.NewSize(400, 200))
		wInput.CenterOnScreen()
		wInput.Show()
	})

	return container.NewVBox(add)
}

func (w *win) listUpdateItem(app fyne.App, id widget.ListItemID) { //function updates materials when another Rank was selected
	Items := decodeitems()
	addbutton := w.item_addbutton(app, id)

	Itemicon := widget.NewIcon(fyne.NewStaticResource("icon", Items[id].icon))

	itemname := widget.NewLabel(Items[id].Name)

	updatebutton := w.item_updatebutton(app, Items[id], id)

	deletebutton := w.item_deletebutton(app, Items[id], id)

	cancel := widget.NewButton("Cancel", func() {
		w.window.Close()
	})
	buttons := container.New(layout.NewVBoxLayout(), itemname, addbutton, updatebutton, deletebutton, cancel)

	list, id := initList_Item(Items, id)

	list.OnSelected = func(id widget.ListItemID) {
		icon := widget.NewIcon(fyne.NewStaticResource("icon", Items[id].icon))

		itemname := widget.NewLabel(Items[id].Name)

		updatebutton := w.item_updatebutton(app, Items[id], id)

		deletebutton := w.item_deletebutton(app, Items[id], id)

		cancel := widget.NewButton("Cancel", func() {
			w.window.Close()
		})
		buttons := container.New(layout.NewVBoxLayout(), itemname, addbutton, updatebutton, deletebutton, cancel)

		gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), buttons, icon) //display gbox
		w.window.SetContent(gbox)
		w.window.Show()

	}

	list.OnUnselected = func(id widget.ListItemID) {

		gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), addbutton) //remove additional widgets
		w.window.SetContent(gbox)                                                             //display gbox
		w.window.Show()
	}

	gbox := container.New(layout.NewGridLayout(2), container.NewHScroll(list), buttons, Itemicon)

	w.window.SetContent(gbox)
	w.window.Show()
}

func (w *win) item_updatebutton(app fyne.App, Item ItemStruct, id widget.ListItemID) fyne.CanvasObject {
	var tempitem TempItemStruct

	tempitem.Name = Item.Name
	tempitem.Rarity = Item.Rarity
	tempitem.Qty = Item.Qty
	tempitem.Qty = Item.Sell
	tempitem.Buy = Item.Buy
	tempitem.EncodedIcon = Item.Encoded

	update := widget.NewButton("Update", func() { //Button to Update Data
		wUpdate := app.NewWindow("Update Data")

		TextItemName := canvas.NewText("Itemname:", color.White)
		TextItemRarity := canvas.NewText("Rarity:", color.White)
		TextItemQty := canvas.NewText("Quantity:", color.White)
		TextItemSell := canvas.NewText("Sell:", color.White)
		TextItemBuy := canvas.NewText("Buy:", color.White)
		InputItemName := widget.NewEntry()
		InputItemRarity := widget.NewEntry()
		InputItemQty := widget.NewEntry()
		InputItemSell := widget.NewEntry()
		InputItemBuy := widget.NewEntry()

		InputItemIcon := widget.NewButton("File Open With Filter (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, wUpdate)
					return
				}
				if reader == nil {
					log.Println("Cancelled")
					return
				}

				tempitem.EncodedIcon = imageOpenedItemIcon(reader, tempitem)
			}, wUpdate)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()

		})

		updateData := widget.NewButton("Update", func() { //Button to add into ItemName typed Data

			if InputItemName.Text != "" {
				tempitem.Name = InputItemName.Text
			}
			if InputItemRarity.Text != "" {
				tempitem.Rarity, _ = strconv.Atoi(InputItemRarity.Text)
			}
			if InputItemQty.Text != "" {
				tempitem.Qty, _ = strconv.Atoi(InputItemQty.Text)
			}
			if InputItemSell.Text != "" {
				tempitem.Sell, _ = strconv.Atoi(InputItemSell.Text)
			}
			if InputItemBuy.Text != "" {
				tempitem.Buy, _ = strconv.Atoi(InputItemBuy.Text)
			}

			UpdateOneItem(client, ctx, Item, tempitem)
			w.listUpdateItem(app, id)
			wUpdate.Close()

		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wUpdate.Close()

		})

		wUpdate.SetContent(container.NewScroll(container.New(layout.NewVBoxLayout(), TextItemName, InputItemName, InputItemIcon, TextItemRarity, InputItemRarity, TextItemQty, InputItemQty, TextItemSell, InputItemSell, TextItemBuy, InputItemBuy, updateData, cancel))) //Layout for the "Insertion-Window"
		wUpdate.Resize(fyne.NewSize(400, 200))
		wUpdate.CenterOnScreen()
		wUpdate.Show()
	})

	return container.NewVBox(update)
}

func (w *win) item_deletebutton(app fyne.App, Item ItemStruct, id widget.ListItemID) fyne.CanvasObject {
	delete := widget.NewButton("Delete", func() { //Button to Delete Items
		DeleteOneItem(client, ctx, Item)
		id = id - 1
		if id >= 0 {
			w.listUpdateItem(app, id)
		}
		if id < 0 {
			id = 0 //set id back to 0 , array out of bounds otherwise -> -1
			list := initemptylist()
			addbutton := w.item_addbutton(app, id)
			gbox := container.NewGridWithColumns(2, container.NewHScroll(list), addbutton)
			w.window.SetContent(gbox)
			w.window.Show()
		}
	})
	return container.NewVBox(delete)
}

func initList_Item(Items []ItemStruct, id widget.ListItemID) (*widget.List, widget.ListItemID) {

	list := widget.NewList(
		func() int {
			return len(Items)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Item", theme.AccountIcon().Content())), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {

			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Items[id].Name, Items[id].icon))
			c.Objects[1].(*widget.Label).SetText(Items[id].Name)
			c.Objects[2].(*widget.Label).SetText("Rarity: " + strconv.Itoa((Items[id].Rarity)))
			c.Objects[3].(*widget.Label).SetText("Qty: " + strconv.Itoa((Items[id].Qty)))
			c.Objects[4].(*widget.Label).SetText("Sell: " + strconv.Itoa((Items[id].Sell)))
			c.Objects[5].(*widget.Label).SetText("Buy: " + strconv.Itoa((Items[id].Buy)))

		},
	)
	return list, id
}
