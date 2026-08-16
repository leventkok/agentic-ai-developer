package main

import (
	"fmt"
	"learn/go/day-13/storage"
)

func main() {
	data, err := storage.ReadMenu("data/menu.txt")

	if err != nil {
		fmt.Println("Error reading menu:", err)
		return
	}

	fmt.Println("Raw File:")
	fmt.Println(string(data))

	menu := storage.ParseMenuLines(data)
	fmt.Println("Parsed Menu:", menu)

	err = storage.SaveOrder("data/order.txt", "latte\nespresso\nlatte\n")
	if err != nil {
		fmt.Println("Error saving order:", err)
		return
	}

	fmt.Println("Order saved successfully, check data/order.txt")

	_, err = storage.ReadMenu("data/missing.txt")
	if err != nil {
		fmt.Println("Expected error: ", err)
	}
}
