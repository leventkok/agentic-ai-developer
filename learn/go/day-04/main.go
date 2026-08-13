package main

import "fmt"

func depositValue(balance int, amount int) {
	balance += amount
}

func depositPointer(balance *int, amount int) {
	*balance += amount
}

func safeDouble(n *int) {
	if n == nil {
		fmt.Println("Error: nil pointer")
		return
	}
	*n *= 2
}

func main() {
	fmt.Println("=== Task 1: & and * ===")
	price := 45
	ptr := &price
	*ptr = 50
	fmt.Println("Price:", price)

	fmt.Println("\n=== Task 2: value vs pointer ===")
	wallet := 100
	depositValue(wallet, 50)
	fmt.Println("After value:", wallet)
	depositPointer(&wallet, 50)
	fmt.Println("After pointer:", wallet)

	fmt.Println("\n=== Task 3: new ===")
	count := new(int)
	*count = 3
	fmt.Println("Count:", *count)

	fmt.Println("\n=== Task 4: nil check ===")
	var nilPtr *int
	safeDouble(nilPtr)
	safeDouble(&wallet)
	fmt.Println("Wallet doubled:", wallet)
}
