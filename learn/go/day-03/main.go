package main

import (
	"fmt"
	"learn/go/day-03/coffe"
)

func calculateTotal(prices []int) int {
	total := 0
	for _, p := range prices {
		total += p

	}

	return total
}

func canPay(balance int, price int) bool {
	return balance >= price
}

func getDrinkPrice(drink string) (int, bool) {
	switch drink {
	case "Latte":
		return 45, true
	case "Espresso":
		return 35, true
	case "Americano":
		return 50, true
	default:
		return 0, false
	}
}

func getOrderSummary(items int, total int) (summary string, itemCount int) {
	summary = fmt.Sprintf("Order: %d items - %d TL", items, total)

	itemCount = items
	return
}

func main() {
	prices := []int{45, 35, 50}
	total := coffe.CalculateTotal(prices)
	fmt.Println("Total:", total)

	price, found := coffe.GetDrinkPrice("Latte")
	if found {
		fmt.Println("Price:", price, "TL")
	}
}
