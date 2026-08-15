package main

import "fmt"
import "learn/go/day-08/menu"

func PrintMenu(items []menu.MenuItem) {
	for _, item := range items {
		fmt.Println(item.String())
	}
}

func Total(items []menu.MenuItem) int {
	sum := 0
	for _, item := range items {
		sum += item.PriceTL()
	}

	return sum
}

func main() {
	items := []menu.MenuItem{
		menu.Drink{Name: "Latte", Price: 45},
		menu.Drink{Name: "Espresso", Price: 35},
		menu.Pastry{Name: "Croissant", Price: 25},
	}

	PrintMenu(items)
	fmt.Printf("Total: %d TL\n", Total(items))

	var item menu.MenuItem
	fmt.Println(item == nil)

	var d *menu.Drink
	var item2 menu.MenuItem = d
	fmt.Println(item2 == nil)
	fmt.Println(item2)

}
