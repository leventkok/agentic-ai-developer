

package main

import (
	"fmt"
	"time"
	"sync"
	"learn/go/day-16/counter"
)


func brew(drink string){
	fmt.Println("Start: ", drink)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done: ", drink)
}


func main(){


	
	go brew("latte")
	go brew("espresso")
	go brew("filter")

	fmt.Println("main: all orders done")


	


	
	var wg sync.WaitGroup

	drinks := []string{"latte", "espresso", "filter"}

	for _, drink := range drinks{
		wg.Add(1)
		go func(d string){
			defer wg.Done()
			brew(d)
		}(drink)
	}


	wg.Wait()
	fmt.Println("main: all orders done")
	

	c := &counter.Counter{}

	counter.UnsafeAdd(c, 1000)

	time.Sleep(time.Second)

	fmt.Println("Counter (unsafe): ", c.Value)
}