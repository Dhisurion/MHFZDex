package main

import (
	//go Imports

	//fyne imports

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	loadIcons()

	//Start DB
	// Get Client, Context, CancelFunc and
	// err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Release resource when the main
	// function is returned.
	defer close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	ping(client, ctx)
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

/*func func2() {
	log.Println("2tapped")
}

func func3() {
	log.Println("3tapped")
}*/

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
		window:  nil,
	}
}

/*IconSelectButton := widget.NewButton("Icon...", func() {
	//w.window = (f *FileDialog) Show()
	//FileDialog :=
	dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error),fyne.Window {
		if e != nil {
			dialog.ShowInformation("Open Error", uc.URI().Path(), w)
			return
		}
		if _, ok := uc.(io.ReadSeeker); ok {
			dialog.ShowInformation("Seeker OK", uc.URI().Path(), w)
		}
		stat, err := os.Stat(uc.URI().Path())
		if err != nil {
			dialog.ShowInformation("Stat Error", uc.URI().Path()+err.Error(), w)
		} else {
			dialog.ShowInformation("Stat Success", fmt.Sprintf("%s %v", uc.URI().Path(), stat.Size()), w)
		}

	})
w.window.Show()
})*/
