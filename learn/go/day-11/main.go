package main

import (
	"errors"
	"fmt"
	"learn/go/day-11/order"
)

func main() {
	// success
	total, err := order.Total([]string{"latte", "espresso"})
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Printf("Total: %d TL \n", total)
	}

	// empty

	_, err = order.Total([]string{})
	if errors.Is(err, order.ErrEmptyOrder) {
		fmt.Println("Caught: empty order")
	}

	// unkown

	_, err = order.Total([]string{"mocha"})

	if errors.Is(err, order.ErrUnknownDrink) {
		fmt.Println("Caught: unknown drink")
	}

}
