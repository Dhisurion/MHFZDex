package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var selectedicon *canvas.Image //used in file dialog
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
	iconb   []byte
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `json:"Name" bson:"Name"`
	Rarity  int                `json:"Rarity" bson:"Rarity"`
	Qty     int                `json:"Qty" bson:"Qty"`
	Sell    int                `json:"Sell" bson:"Sell"`
	Buy     int                `json:"Buy" bson:"Buy"`
	Encoded string             `json:"Icon" bson:"Icon"`
}

type WeaponStruct struct {
	iconb        []byte
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `json:"Name" bson:"Name"`
	Kind         string             `json:"Kind" bson:"Kind"`
	Rarity       int                `json:"Rarity" bson:"Rarity"`
	Attack       int                `json:"Attack" bson:"Attack"`
	Element      string             `json:"Element" bson:"Element"`
	Elementvalue int                `json:"Elementvalue" bson:"Elementvalue"`
	Sharpness    string             `json:"Sharpness" bson:"Sharpness"`
	Affinity     int                `json:"Affinity" bson:"Affinity"`
	Defense      int                `json:"Defense" bson:"Defense"`
	Price        int                `json:"Price" bson:"Price"`
	Material     string             `json:"Material" bson:"Material"`
	Quantity     int                `json:"Quantity" bson:"Quantity"`
	Encoded      string             `json:"Icon" bson:"Icon"`
}
type TempMonsterStruct struct {
	Name    string   `json:"Name" bson:"Name"`
	Fire    [6]int   `json:"Fire" bson:"Fire"`
	Thunder [6]int   `json:"Thunder" bson:"Thunder"`
	Water   [6]int   `json:"Ice" bson:"Ice"`
	Ice     [6]int   `json:"Water" bson:"Water"`
	Dragon  [6]int   `json:"Dragon" bson:"Dragon"`
	LRMat   []string `json:"LRMat" bson:"LRMat"`
	HRMat   []string `json:"HRMat" bson:"HRMat"`
	GouMat  []string `json:"GouMat" bson:"GouMat"`
	GMat    []string `json:"GMat" bson:"GMat"`
	ZMat    []string `json:"ZMat" bson:"ZMat"`
	Encoded string   `json:"Icon" bson:"Icon"`
}

type TempItemStruct struct {
	Name    string `json:"Name" bson:"Name"`
	Rarity  int    `json:"Rarity" bson:"Rarity"`
	Qty     int    `json:"Qty" bson:"Qty"`
	Sell    int    `json:"Sell" bson:"Sell"`
	Buy     int    `json:"Buy" bson:"Buy"`
	Encoded string `json:"Icon" bson:"Icon"`
}

type TempWeaponStruct struct {
	Name         string `json:"Name" bson:"Name"`
	Kind         string `json:"Kind" bson:"Kind"` //weapontype but type is a bad declaration
	Rarity       int    `json:"Rarity" bson:"Rarity"`
	Attack       int    `json:"Attack" bson:"Attack"`
	Element      string `json:"Element" bson:"Element"`
	Elementvalue int    `json:"Elementvalue" bson:"Elementvalue"`
	Sharpness    string `json:"Sharpness" bson:"Sharpness"`
	Affinity     int    `json:"Affinity" bson:"Affinity"`
	Defense      int    `json:"Defense" bson:"Defense"`
	Price        int    `json:"Price" bson:"Price"`
	Material     string `json:"Material" bson:"Material"`
	Quantity     int    `json:"Quantity" bson:"Quantity"`
	Encoded      string `json:"Icon" bson:"Icon"`
}
