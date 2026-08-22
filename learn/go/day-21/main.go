package main

import (
	"fmt"

	"learn/go/day-21/pricing"
)

func main() {
	fmt.Println("Latte x2:", pricing.Total(45, 2))
	fmt.Println("10% off:", pricing.ApplyDiscount(90, 10))
}
