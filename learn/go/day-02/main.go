package main

import "fmt"

func main() {
	balance := 1000
	price := 80

	if balance >= price {
		fmt.Println("Can pay")
	} else {
		fmt.Println("Can't pay")
	}

	prices := []int{45, 35, 50}

	total := 0

	for _, p := range prices {
		total += p
	}

	fmt.Println("Total:", total)

	drink := "Latte"
	switch drink {
	case "Latte":
		fmt.Println("Latte is 45 TL")
	case "Espresso":
		fmt.Println("Espresso is 35 TL")
	case "Americano":
		fmt.Println("Americano is 50 TL")
	default:
		fmt.Println("Unknown drink")
	}
}
