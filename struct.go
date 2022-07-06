package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var Monsterselectedicon *canvas.Image //used
var Itemselectedicon *canvas.Image    //in
var Armorselectedicon *canvas.Image   //icon
var Weaponselectedicon *canvas.Image  //selector
//var encoded string

type win struct {
	//equation string

	//output  *widget.Label
	buttons map[string]*widget.Button
	window  fyne.Window
}

/*ar apps = []Monster{
	{1, main.Rathalos, "Rathalos"},
	{2, main.Seregios, "Sergios"},
}*/

/*type MonsterList struct {
	Monsters []Monster
}*/

type MonsterStruct struct {
	icon *canvas.Image
	name string
	//encoded string
}

type ItemStruct struct {
	icon    *canvas.Image
	name    string
	rarity  int
	qty     int
	sell    int
	buy     int
	encoded string
}

type TempMonsterStruct struct {
	icon    *canvas.Image
	name    string
	encoded string
}

type TempItemStruct struct {
	icon    *canvas.Image
	name    string
	rarity  int
	qty     int
	sell    int
	buy     int
	encoded string
}
