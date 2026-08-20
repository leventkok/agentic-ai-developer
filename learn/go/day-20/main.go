package main

import (
	"fmt"
	"time"

	"learn/go/day-20/fetch"
	"learn/go/day-20/pipeline"
)

func main() {
	fmt.Println("=== Task 1: Concurrent fetch ===")
	demoFetchAll()

	fmt.Println("\n=== Task 2: Pipeline ===")
	demoPipeline()

	fmt.Println("\n=== Task 3: Timeout ===")
	demoTimeout()

	fmt.Println("\n=== Task 4: Race check ===")
	fmt.Println("Run: go run -race .")
}

func demoFetchAll() {
	urls := []string{
		"https://example.com/menu",
		"https://example.com/orders",
		"https://example.com/prices",
	}
	results := fetch.FetchAll(urls)
	for _, r := range results {
		fmt.Printf("%s → %s\n", r.URL, r.Status)
	}
}

func demoPipeline() {
	nums := pipeline.Generate(1, 2, 3, 4, 5)
	squared := pipeline.Square(nums)
	pipeline.Print(squared)
}

func demoTimeout() {
	r, err := fetch.FetchWithTimeout("https://example.com/menu", 1*time.Second)
	if err != nil {
		fmt.Println("Fast fetch error:", err)
	} else {
		fmt.Printf("Fast fetch OK: %s → %s\n", r.URL, r.Status)
	}

	_, err = fetch.FetchSlowWithTimeout("https://example.com/slow", 500*time.Millisecond)
	if err != nil {
		fmt.Println("Slow fetch (expected timeout):", err)
	} else {
		fmt.Println("Slow fetch: unexpected success")
	}
}
