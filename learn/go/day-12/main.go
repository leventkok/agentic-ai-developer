package main

import (
	"fmt"
	"learn/go/day-12/menu"
)

func main() {

	// len vs cap

	s := make([]string, 0, 3)
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))

	s = append(s, "latte")
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))

	// backing array

	original := []string{"latte", "espresso", "filter"}

	view := original[0:2]
	view[0] = "CHANGED"

	fmt.Println(original)

	// order demo

	order := menu.NewOrder()
	order.Add("latte")
	order.Add("espresso")
	order.Add("latte")

	fmt.Printf("Order: %v\n", order.Items)
	fmt.Printf("Total: %d TL\n", order.Total())

	// map demo

	menu.AddItem("mocha", 50)
	menu.PrintMenu()

	price, ok := menu.GetPrice("mocha")
	if ok {
		fmt.Printf("Mocha price: %d TL\n", price)
	}

}
