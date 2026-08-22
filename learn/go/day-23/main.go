package main

import (
	"fmt"

	"learn/go/day-23/pricing"
)

func main() {
	total, err := pricing.DiscountedTotal(45, 2, 10)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Discounted total:", total)
}
