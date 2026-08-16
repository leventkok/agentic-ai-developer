package main

import (
	"encoding/json"
	"fmt"
	"learn/go/day-14/jsonio"
)

func main() {
	menu := jsonio.Menu{
		Items: []jsonio.MenuItem{
			{Name: "latte", Price: 45},
			{Name: "espresso", Price: 35},
			{Name: "filter", Price: 30},
		},
	}
	data, err := jsonio.MenuToJSON(menu)
	if err != nil {
		fmt.Println("Error converting menu to JSON:", err)
		return
	}
	fmt.Println("JSON bytes:")
	fmt.Println(string(data))

	loaded, err := jsonio.MenuFromJSON(data)
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}
	fmt.Println("Loaded Menu:", loaded)

	orderA := jsonio.Order{Items: []string{"latte", "espresso"}, Total: 80}
	orderJSON, _ := json.Marshal(orderA)
	fmt.Println("Order A:", string(orderJSON))

	note := "extra hot"
	orderB := jsonio.Order{Items: []string{"latte"}, Total: 45, Note: &note}
	orderJSON2, _ := json.Marshal(orderB)
	fmt.Println("Order B:", string(orderJSON2))

	err = jsonio.SaveMenu("data/menu.json", menu)
	if err != nil {
		fmt.Println("save error:", err)
		return
	}

	loadedFromFile, err := jsonio.LoadMenu("data/menu.json")
	if err != nil {
		fmt.Println("load error:", err)
		return
	}
	fmt.Println("From file:", loadedFromFile)
}
