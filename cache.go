package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func loadJsonData() []string {
	fmt.Println("Loading data from JSON file")

	input, _ := ioutil.ReadFile("data.json")
	var data []string
	json.Unmarshal(input, &data)

	return data
}
