package main

import (
	"fmt"
	"learn/go/day-17/worker"
)

func main() {
	orders := make(chan string)

	go func() {
		orders <- "latte"
		orders <- "espresso"
	}()

	fmt.Println("Receive:", <-orders)
	fmt.Println("Receive:", <-orders)

	buffered := make(chan string, 2)
	buffered <- "latte"
	buffered <- "espresso"

	fmt.Println(<-buffered)
	fmt.Println(<-buffered)

	drinks := make(chan string)

	go func() {
		drinks <- "latte"
		drinks <- "espresso"
		drinks <- "filter"
		close(drinks)
	}()

	for drink := range drinks {
		fmt.Println("Served:", drink)
	}

	fmt.Println("ALL Done")

	jobs := make(chan worker.Job, 3)
	results := make(chan worker.Result, 3)

	go worker.Brew(jobs, results)

	orderList := []string{"latte", "espresso", "latte"}
	for _, drink := range orderList {
		jobs <- worker.Job{Drink: drink}
	}
	close(jobs)

	for i := 0; i < len(orderList); i++ {
		r := <-results
		fmt.Printf("%s: %d TL\n", r.Drink, r.Price)
	}
}
