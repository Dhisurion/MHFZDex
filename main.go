package main

import (
	//go Imports

	//fyne imports

	"context"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/widget"
)

func main() {

	//Start DB
	// Get Client, Context, CancelFunc and
	// err from connect method.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, ctx, cancel, err = connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Release resource when the main
	// function is returned.
	defer close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	ping(client, ctx)
	//exInsertOne(client, ctx)
	//End DB

	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	button1 := widget.NewButton("Monster", func() {
		w := MonsterWindow()
		w.MonsterUI(myApp)

	})
	button2 := widget.NewButton("Weapons", func() {
		w := WeaponWindow()
		w.WeaponUI(myApp)
	})
	button3 := widget.NewButton("Armor", func() {
		w := ArmorWindow()
		w.ArmorUI(myApp)
	})

	button4 := widget.NewButton("Item", func() {
		w := ItemWindow()
		w.ItemUI(myApp)
	})
	grid := container.New(layout.NewGridLayout(2), button1, button2, button3, button4)
	myWindow.SetContent(grid)
	myWindow.SetMaster()
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}

func MonsterWindow() *win { //launchess MonsterWindow
	return &win{
		buttons: make(map[string]*widget.Button, 19),
	}
}

func ItemWindow() *win { //launches ItemWindow
	return &win{
		buttons: make(map[string]*widget.Button, 19)}
}

func WeaponWindow() *win { //launches WeaponWindow
	return &win{
		buttons: make(map[string]*widget.Button, 19),
	}
}

func ArmorWindow() *win { //launches ArmorWindow
	return &win{
		buttons: map[string]*widget.Button{},
	}
}
