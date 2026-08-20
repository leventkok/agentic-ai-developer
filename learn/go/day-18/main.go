package main

import (
	"fmt"
	"time"

	"learn/go/day-18/kitchen"
)

func main() {
	fmt.Println("=== Task 1: select (wait on multiple channels) ===")
	demoSelect()

	fmt.Println("\n=== Task 2: default (non-blocking) ===")
	demoDefault()

	fmt.Println("\n=== Task 3: timeout ===")
	demoTimeout()

	fmt.Println("\n=== Task 4: done channel (cancel) ===")
	demoCancel()

	fmt.Println("\n=== Task 5: backpressure (drop when full) ===")
	demoBackpressure()
}

func demoSelect() {
	barista := make(chan string)
	kitchenCh := make(chan string)

	go func() { barista <- "latte" }()
	go func() { kitchenCh <- "sandwich" }()

	select {
	case drink := <-barista:
		fmt.Println("From barista:", drink)
	case food := <-kitchenCh:
		fmt.Println("From kitchen:", food)
	}
}

func demoDefault() {
	orders := make(chan string)

	select {
	case drink := <-orders:
		fmt.Println("Got order:", drink)
	default:
		fmt.Println("No orders right now (non-blocking)")
	}
}

func demoTimeout() {
	orders := make(chan string)

	select {
	case drink := <-orders:
		fmt.Println("Received:", drink)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout: no order in 500ms")
	}
}

func demoCancel() {
	orders := make(chan string)
	done := make(chan struct{})

	go kitchen.BrewUntilDone(orders, done)

	orders <- "latte"
	orders <- "espresso"

	close(done)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("main: shut down")
}

func demoBackpressure() {
	results := make(chan string, 1)
	results <- "latte" // buffer full

	select {
	case results <- "espresso":
		fmt.Println("Sent espresso")
	default:
		fmt.Println("Dropped espresso — consumer too slow")
	}
}
