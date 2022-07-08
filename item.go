package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image/color"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (w *win) ItemUI(app fyne.App) {
	w.window = app.NewWindow("Item")

	Items := decodeitems()

	itembuttons := w.item_addbutton(app)

	list := widget.NewList(
		func() int {
			return len(Items)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Item", item.iconb)), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {

			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Items[id].Name, Items[id].iconb))
			c.Objects[1].(*widget.Label).SetText(Items[id].Name)
			c.Objects[2].(*widget.Label).SetText("Rarity: " + strconv.Itoa((Items[id].Rarity)))
			c.Objects[3].(*widget.Label).SetText("Qty: " + strconv.Itoa((Items[id].Qty)))
			c.Objects[4].(*widget.Label).SetText("Sell: " + strconv.Itoa((Items[id].Sell)))
			c.Objects[5].(*widget.Label).SetText("Buy: " + strconv.Itoa((Items[id].Buy)))

		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		icon := widget.NewIcon(fyne.NewStaticResource("icon", Items[id].iconb))

		itemname := widget.NewLabel(Items[id].Name)

		updatebutton := w.item_updatebutton(app, Items[id])

		deletebutton := w.item_deletebutton(app, Items[id])

		cancel := widget.NewButton("Cancel", func() {
			w.window.Close()
		})
		choice := container.New(layout.NewVBoxLayout(), itemname, updatebutton, deletebutton, cancel)

		gbox := container.New(layout.NewGridLayout(3), list, choice, itembuttons, icon) //display gbox
		w.window.SetContent(gbox)
		w.window.Show()

	}

	list.OnUnselected = func(id widget.ListItemID) {

		gbox := container.New(layout.NewGridLayout(3), list, itembuttons) //remove additional widgets
		w.window.SetContent(gbox)                                         //display gbox
		w.window.Show()
	}

	gbox := container.New(layout.NewGridLayout(3), list, itembuttons)

	w.window.SetContent(gbox)

	w.window.Show()
}

func (w *win) item_addbutton(app fyne.App) fyne.CanvasObject {

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

		InputItemIcon := widget.NewButton("File Open With Filter (.jpg or .png)", func() {
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

		addData := widget.NewButton("Add", func() { //Button to add into ItemName typed Data
			tempitem.Name = InputItemName.Text
			tempitem.Rarity, _ = strconv.Atoi(InputItemRarity.Text)
			//tempitem.rarity, _= strconv.Atoi(InputItemRarity.Text)

			tempitem.Qty, _ = strconv.Atoi(InputItemQty.Text)

			tempitem.Sell, _ = strconv.Atoi(InputItemSell.Text)

			tempitem.Buy, _ = strconv.Atoi(InputItemBuy.Text)

			InsertOne(client, ctx)
			w.listUpdateItem(app)
			wInput.Close()

		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wInput.Close()

		})

		wInput.SetContent(container.New(layout.NewVBoxLayout(), TextItemName, InputItemName, InputItemIcon, TextItemRarity, InputItemRarity, TextItemQty, InputItemQty, TextItemSell, InputItemSell, TextItemBuy, InputItemBuy, addData, cancel)) //Layout for the "Insertion-Window"
		wInput.Resize(fyne.NewSize(400, 200))
		wInput.CenterOnScreen()
		wInput.Show()
	})

	return container.NewVBox(add)
}

func decodeitems() []ItemStruct {
	Items, err := ReadAllItems(client, ctx)
	if err != nil {
		panic(err)
	}

	for i := range Items {

		decoded, err := base64.StdEncoding.DecodeString((Items[i].Encoded))
		if err != nil {
			fmt.Printf("Error decoding", err.Error())
			panic(err)
		}

		f, err := os.Create("res/item/Item" + strconv.Itoa(i) + ".jpg")
		if err != nil {
			fmt.Printf("Error creating file", err.Error())
			panic(err)
		}

		if _, err := f.Write(decoded); err != nil {
			panic(err)
		}

		if err := f.Sync(); err != nil {
			panic(err)
		}

		iconFile, err := os.Open("res/item/Item" + strconv.Itoa(i) + ".jpg")
		if err != nil {
			log.Fatal(err)
		}

		r := bufio.NewReader(iconFile)

		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		Items[i].iconb = b

	}
	return Items
}

func (w *win) listUpdateItem(app fyne.App) { //function updates materials when another Rank was selected
	Items := decodeitems()
	itembuttons := w.item_addbutton(app)

	icon := widget.NewIcon(fyne.NewStaticResource("icon", Items[1].iconb))

	list := widget.NewList(
		func() int {
			return len(Items)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(fyne.NewStaticResource("Item", item.iconb)), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {

			c := obj.(*fyne.Container)
			c.Objects[0].(*widget.Icon).SetResource(fyne.NewStaticResource(Items[id].Name, Items[id].iconb))
			c.Objects[1].(*widget.Label).SetText(Items[id].Name)
			c.Objects[2].(*widget.Label).SetText("Rarity: " + strconv.Itoa((Items[id].Rarity)))
			c.Objects[3].(*widget.Label).SetText("Qty: " + strconv.Itoa((Items[id].Qty)))
			c.Objects[4].(*widget.Label).SetText("Sell: " + strconv.Itoa((Items[id].Sell)))
			c.Objects[5].(*widget.Label).SetText("Buy: " + strconv.Itoa((Items[id].Buy)))

		},
	)

	gbox := container.New(layout.NewGridLayout(3), list, itembuttons, icon)

	w.window.SetContent(gbox)
	w.window.Show()
}

func (w *win) item_updatebutton(app fyne.App, Item ItemStruct) fyne.CanvasObject {

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

				imageOpened(reader)
			}, wUpdate)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()

		})

		updateData := widget.NewButton("Update", func() { //Button to add into ItemName typed Data
			tempitem.Name = InputItemName.Text
			tempitem.Rarity, _ = strconv.Atoi(InputItemRarity.Text)
			//tempitem.rarity, _= strconv.Atoi(InputItemRarity.Text)

			tempitem.Qty, _ = strconv.Atoi(InputItemQty.Text)

			tempitem.Sell, _ = strconv.Atoi(InputItemSell.Text)

			tempitem.Buy, _ = strconv.Atoi(InputItemBuy.Text)

			UpdateOne(client, ctx, Item)
			w.listUpdateItem(app)
			wUpdate.Close()

		})

		cancel := widget.NewButton("Cancel", func() { //cancel data-input
			wUpdate.Close()

		})

		wUpdate.SetContent(container.New(layout.NewVBoxLayout(), TextItemName, InputItemName, InputItemIcon, TextItemRarity, InputItemRarity, TextItemQty, InputItemQty, TextItemSell, InputItemSell, TextItemBuy, InputItemBuy, updateData, cancel)) //Layout for the "Insertion-Window"
		wUpdate.Resize(fyne.NewSize(400, 200))
		wUpdate.CenterOnScreen()
		wUpdate.Show()
	})

	return container.NewVBox(update)
}

func (w *win) item_deletebutton(app fyne.App, Item ItemStruct) fyne.CanvasObject {
	delete := widget.NewButton("Delete", func() { //Button to Delete Items
		DeleteOne(client, ctx, Item)

	})
	return container.NewVBox(delete)
}
