package main

import (
	"fmt"
	"sync"

	"learn/go/day-19/config"
	"learn/go/day-19/counter"
)

func main() {
	fmt.Println("=== Task 1: Mutex ===")
	demoMutex()

	fmt.Println("\n=== Task 2: sync.Once ===")
	demoOnce()

	fmt.Println("\n=== Task 3: Atomic ===")
	demoAtomic()

	fmt.Println("\n=== Task 4: Channels vs Mutex ===")
	demoCompare()
}

func demoMutex() {
	c := &counter.SafeCounter{}
	c.Add(1000, 10)
	fmt.Println("Safe counter (mutex):", c.Value)
}

func demoOnce() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(config.ShopName())
		}()
	}
	wg.Wait()
}

func demoAtomic() {
	var wg sync.WaitGroup
	ac := &counter.AtomicCounter{}
	counter.AtomicAdd(ac, 1000, &wg)
	wg.Wait()
	fmt.Println("Atomic counter:", ac.Get())
}

func demoCompare() {
	fmt.Println("Channels  → pass data between goroutines (order queue, worker pipeline)")
	fmt.Println("Mutex     → protect shared struct/map many goroutines touch")
	fmt.Println("Atomic    → single int counter, high frequency")
	fmt.Println("sync.Once → expensive setup runs exactly once")
}
