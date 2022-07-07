package main

import (
	//"io"

	"encoding/base64"
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"os"
	"strconv"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	//"fyne.io/fyne/v2/data/binding"
	//"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (w *win) ItemUI(app fyne.App) {
	w.window = app.NewWindow("Item")
	empty := widget.NewLabel("text")

	//number := Count(client, ctx, "Items")

	Items := ReadAllItems(client, ctx)

	//data := make([]string, 1000)
	//i := 1
	/*for i := range data {
		data[i] = strconv.Itoa(i+1) + " Test Item "
	}*/

	for i := range Items {

		decoded, err := base64.StdEncoding.DecodeString((Items[i].encoded))
		if err != nil {
			fmt.Printf("Error decoding", err.Error())
			return
		}

		f, err := os.Create("Item " + strconv.Itoa(i))

		if _, err := f.Write(decoded); err != nil {
			panic(err)
		}

		if err := f.Sync(); err != nil {
			panic(err)
		}

		Items[i].icon = canvas.NewImageFromReader(f, string(decoded))
	}

	list := widget.NewList(
		func() int {
			return len(Items)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("Template Object"), canvas.NewImageFromResource(theme.CancelIcon()))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(Items[id].name)
		},
	)

	buttons := item_funcbuttons(app)
	//test := []string{"Brachydios", "Rathalos", "Rathian"}

	gbox := container.New(layout.NewGridLayout(3), list, buttons, empty)

	w.window.SetContent(gbox)
	//w.window= SetContent(grid)
	w.window.Show()
}

func item_funcbuttons(app fyne.App) fyne.CanvasObject {

	//F[0] = widget.NewEntry().Text
	add := widget.NewButton("Add", func() { //Button to Add Data
		wInput := app.NewWindow("Add Data")

		//ID := widget.NewEntry()
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

				//imageOpened(reader)
			}, wInput)
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
			fd.Show()

		})

		addData := widget.NewButton("Add", func() { //Button to add into ItemName typed Data
			tempitem.name = InputItemName.Text
			tempitem.rarity, _ = strconv.Atoi(InputItemRarity.Text)
			//tempitem.rarity, _= strconv.Atoi(InputItemRarity.Text)

			tempitem.qty, _ = strconv.Atoi(InputItemQty.Text)

			tempitem.sell, _ = strconv.Atoi(InputItemSell.Text)

			tempitem.buy, _ = strconv.Atoi(InputItemBuy.Text)

			InsertOne(client, ctx)
			fmt.Println(tempmonster.encoded)
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

	/*exit := widget.NewButton("Close", func() { //added close button for whatever reason...cross-platform , maybe? idk

		w.window.Close()
	})*/

	return container.NewVBox(add)
}
