package main

import (
	"fmt"
	"learn/go/day-07/menu"
)

func main() {
	fmt.Println("=== RotaCoffee Menu ===")
	drinks := []menu.Drink{
		{Name: "Espresso", Price: 30},
		{Name: "Latte", Price: 45},
		{Name: "Filter", Price: 30},
	}
	for _, d := range drinks {
		fmt.Println(d.String())
	}

	premium := drinks[1].IsPremium()
	fmt.Printf("Is Latte premium? %v\n", premium)

	drinks[1].ApplyDiscount(10)
	fmt.Printf("Discounted Latte: %v\n", drinks[1])

	special := &menu.Drink{Name: "Flat White", Price: 50}
	special.SetPrice(55)
	fmt.Printf("Flat White: %v\n", special.String())

	fmt.Println("--- after value range (no change) ---")
	for _, d := range drinks {
    	fmt.Println(d.String())
	}

	for i := range drinks {
    	drinks[i].ApplyDiscount(10)
	}

	fmt.Println("--- after index range (changed) ---")
	for _, d := range drinks {
    	fmt.Println(d.String())
	}

}
