package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Data struct {
	InventoryId int       `json:"inventory_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags"`
	PurchaseAt  time.Time `json:"purchased_at"`
	Placement   Placement `json:"placement"`
}

type Placement struct {
	RoomId        int    `json:"room_id"`
	PlacementName string `json:"name"`
}

func main() {

	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var x []*Data

	json.Unmarshal(byteValue, &x)

	var x2 []*Data
	for _, v := range x {
		purchased, _ := time.Parse("2020-01-16", "2020-01-16")
		purchasedat, _ := time.Parse("2020-01-16", fmt.Sprint(v.PurchaseAt))
		if purchasedat == purchased {
			x2 = append(x2, v)
		}
	}

	data, _ := json.MarshalIndent(x2, "", " ")
	fmt.Println(string(data))
}
