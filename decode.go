package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

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

		Items[i].icon = b

	}
	return Items
}

func decodeweapons() []WeaponStruct {
	Weapons, err := ReadAllWeapons(client, ctx)
	if err != nil {
		panic(err)
	}

	for i := range Weapons {

		decoded, err := base64.StdEncoding.DecodeString((Weapons[i].Encoded))
		if err != nil {
			fmt.Printf("Error decoding", err.Error())
			panic(err)
		}

		f, err := os.Create("res/weapon/Weapon" + strconv.Itoa(i) + ".jpg")
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

		iconFile, err := os.Open("res/weapon/Weapon" + strconv.Itoa(i) + ".jpg")
		if err != nil {
			log.Fatal(err)
		}

		r := bufio.NewReader(iconFile)

		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		Weapons[i].icon = b

	}
	return Weapons
}

func decodemonsters() []MonsterStruct {
	Monsters, err := ReadAllMonsters(client, ctx)
	if err != nil {
		panic(err)
	}

	for i := range Monsters {

		decodedIcon, err := base64.StdEncoding.DecodeString((Monsters[i].EncodedIcon))
		if err != nil {
			fmt.Printf("Error decoding", err.Error())
			panic(err)
		}

		decodedPic, err := base64.StdEncoding.DecodeString((Monsters[i].EncodedPic))
		if err != nil {
			fmt.Printf("Error decoding", err.Error())
			panic(err)
		}

		f, err := os.Create("res/monster/icon/Monster" + strconv.Itoa(i) + ".jpg")
		if err != nil {
			fmt.Printf("Error creating file", err.Error())
			panic(err)
		}

		if _, err := f.Write(decodedIcon); err != nil {
			panic(err)
		}

		if err := f.Sync(); err != nil {
			panic(err)
		}

		f, err = os.Create("res/monster/pic/Monster" + strconv.Itoa(i) + ".jpg")
		if err != nil {
			fmt.Printf("Error creating file", err.Error())
			panic(err)
		}

		if _, err := f.Write(decodedPic); err != nil {
			panic(err)
		}

		if err := f.Sync(); err != nil {
			panic(err)
		}

		iconFile, err := os.Open("res/monster/icon/Monster" + strconv.Itoa(i) + ".jpg")
		if err != nil {
			log.Fatal(err)
		}

		picFile, err := os.Open("res/monster/pic/Monster" + strconv.Itoa(i) + ".jpg")
		if err != nil {
			log.Fatal(err)
		}

		r := bufio.NewReader(iconFile)
		r2 := bufio.NewReader(picFile)

		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		b2, err := ioutil.ReadAll(r2)
		if err != nil {
			log.Fatal(err)
		}

		Monsters[i].icon = b
		Monsters[i].pic = b2

	}
	return Monsters
}
