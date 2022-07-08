package main

import (
	"encoding/base64"

	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

//Item Dialog

func imageOpenedItemIcon(f fyne.URIReadCloser) {
	if f == nil {
		log.Println("Cancelled")
		return
	}
	defer f.Close()

	selectedItemicon = loadImage(f)
	tempitem.EncodedIcon = base64.StdEncoding.EncodeToString([]byte(selectedItemicon.Resource.Content()))
	//showImage(f)
}

//Monster Dialog
func imageOpenedMonsterIcon(f fyne.URIReadCloser) {
	if f == nil {
		log.Println("Cancelled")
		return
	}
	defer f.Close()

	selectedMonstericon = loadImage(f)
	tempmonster.EncodedIcon = base64.StdEncoding.EncodeToString([]byte(selectedMonstericon.Resource.Content()))

}

func imageOpenedMonsterPic(f fyne.URIReadCloser) {
	if f == nil {
		log.Println("Cancelled")
		return
	}
	defer f.Close()

	selectedMonsterPic = loadImage(f)
	tempmonster.EncodedPic = base64.StdEncoding.EncodeToString([]byte(selectedMonsterPic.Resource.Content()))

}

//Weapon Dialog

func imageOpenedWeaponIcon(f fyne.URIReadCloser) {
	if f == nil {
		log.Println("Cancelled")
		return
	}
	defer f.Close()

	selectedWeaponicon = loadImage(f)
	tempitem.EncodedIcon = base64.StdEncoding.EncodeToString([]byte(selectedWeaponicon.Resource.Content()))

}

//Armor Dialog

func imageOpenedArmorIcon(f fyne.URIReadCloser) {
	if f == nil {
		log.Println("Cancelled")
		return
	}
	defer f.Close()

	selectedArmoricon = loadImage(f)
	tempitem.EncodedIcon = base64.StdEncoding.EncodeToString([]byte(selectedArmoricon.Resource.Content()))

}

//Other Image Functions

func showImage(f fyne.URIReadCloser) {
	img := loadImage(f)
	if img == nil {
		return
	}
	img.FillMode = canvas.ImageFillOriginal

	w := fyne.CurrentApp().NewWindow(f.URI().Name())
	w.SetContent(container.NewScroll(img))
	w.Resize(fyne.NewSize(320, 240))
	w.Show()
}

func loadImage(f fyne.URIReadCloser) *canvas.Image {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fyne.LogError("Failed to load image data", err)
		return nil
	}
	res := fyne.NewStaticResource(f.URI().Name(), data)

	return canvas.NewImageFromResource(res)
}
