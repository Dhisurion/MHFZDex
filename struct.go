package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var selectedItemicon *canvas.Image
var selectedMonstericon *canvas.Image
var selectedWeaponicon *canvas.Image
var selectedArmoricon *canvas.Image  //used in file dialog
var selectedMonsterPic *canvas.Image //used as monster pic

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

//Item
type ItemStruct struct {
	icon    []byte
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `json:"Name" bson:"Name"`
	Rarity  int                `json:"Rarity" bson:"Rarity"`
	Qty     int                `json:"Qty" bson:"Qty"`
	Sell    int                `json:"Sell" bson:"Sell"`
	Buy     int                `json:"Buy" bson:"Buy"`
	Encoded string             `json:"Icon" bson:"Icon"`
}

type TempItemStruct struct {
	Name        string `json:"Name" bson:"Name,omitempty"`
	Rarity      int    `json:"Rarity" bson:"Rarity,omitempty"`
	Qty         int    `json:"Qty" bson:"Qty,omitempty"`
	Sell        int    `json:"Sell" bson:"Sell,omitempty"`
	Buy         int    `json:"Buy" bson:"Buy,omitempty"`
	EncodedIcon string `json:"Icon" bson:"Icon,omitempty"`
}

//Weapon
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

//Monster

type MonsterStruct struct {
	icon        []byte
	pic         []byte
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `json:"Name" bson:"Name"`
	Fire        [7]int             `json:"FireWeakness" bson:"FireWeakness"`
	Thunder     [7]int             `json:"ThunderWeakness" bson:"ThunderWeakness"`
	Water       [7]int             `json:"WaterWeakness" bson:"WaterWeakness"`
	Ice         [7]int             `json:"IceWeakness" bson:"IceWeakness"`
	Dragon      [7]int             `json:"DragonWeakness" bson:"DragonWeakness"`
	LRMat       [10]string         `json:"MaterialsLowRank" bson:"MaterialsLowRank"`
	HRMat       [10]string         `json:"MaterialsHighRank" bson:"MaterialsHighRank"`
	GouRMat     [10]string         `json:"MaterialsGouRank" bson:"MaterialsGouRank"`
	GRMat       [10]string         `json:"MaterialsGRank" bson:"MaterialsGRank"`
	ZRMat       [10]string         `json:"MaterialsZenithRank" bson:"MaterialsZenithRank"`
	EncodedIcon string             `json:"Icon" bson:"Icon"`
	EncodedPic  string             `json:"Pic" bson:"Pic"`
}

type TempMonsterStruct struct {
	Name        string     `json:"Name" bson:"Name,omitempty"`
	Fire        [7]int     `json:"FireWeakness" bson:"FireWeakness,omitempty"`
	Thunder     [7]int     `json:"ThunderWeakness" bson:"ThunderWeakness,omitempty"`
	Water       [7]int     `json:"WaterWeakness" bson:"WaterWeakness,omitempty"`
	Ice         [7]int     `json:"IceWeakness" bson:"IceWeakness,omitempty"`
	Dragon      [7]int     `json:"DragonWeakness" bson:"DragonWeakness,omitempty"`
	LRMat       [10]string `json:"MaterialsLowRank" bson:"MaterialsLowRank,omitempty"`
	HRMat       [10]string `json:"MaterialsHighRank" bson:"MaterialsHighRank,omitempty"`
	GouRMat     [10]string `json:"MaterialsGouRank" bson:"MaterialsGouRank,omitempty"`
	GRMat       [10]string `json:"MaterialsGRank" bson:"MaterialsGRank,omitempty"`
	ZRMat       [10]string `json:"MaterialsZenithRank" bson:"MaterialsZenithRank,omitempty"`
	EncodedIcon string     `json:"Icon" bson:"Icon,omitempty"`
	EncodedPic  string     `json:"Pic" bson:"Pic,omitempty"`
}
