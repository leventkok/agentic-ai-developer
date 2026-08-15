package main

import (
	"fmt"
	"learn/go/day-06/menu"
)

func main() {
	drinks := []menu.Drink{
		{Name: "Espresso", Price: 30},
		{Name: "Latte", Price: 45},
		{Name: "Filter", Price: 30},
	}

	fmt.Println("=== RotaCoffee Menu ===")

	for _, d := range drinks {
		fmt.Println(d.String())
	}

	var empty menu.Drink
	fmt.Printf("Zero value: %+v", empty)

}
