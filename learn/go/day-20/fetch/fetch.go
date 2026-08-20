package fetch

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	URL    string
	Status string
	Err    error
}

func Fetch(url string) Result {
	time.Sleep(200 * time.Millisecond)
	return Result{URL: url, Status: "200 OK", Err: nil}
}

func FetchAll(urls []string) []Result {
	results := make([]Result, len(urls))
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			results[i] = Fetch(url)
		}(i, url)
	}

	wg.Wait()
	return results
}

func FetchWithTimeout(url string, timeout time.Duration) (Result, error) {
	ch := make(chan Result, 1)

	go func() {
		ch <- Fetch(url)
	}()

	select {
	case r := <-ch:
		return r, nil
	case <-time.After(timeout):
		return Result{URL: url}, fmt.Errorf("timeout after %v", timeout)
	}
}

func FetchSlow(url string) Result {
	time.Sleep(2 * time.Second)
	return Result{URL: url, Status: "200 OK"}
}

func FetchSlowWithTimeout(url string, timeout time.Duration) (Result, error) {
	ch := make(chan Result, 1)

	go func() {
		ch <- FetchSlow(url)
	}()

	select {
	case r := <-ch:
		return r, nil
	case <-time.After(timeout):
		return Result{URL: url}, fmt.Errorf("timeout after %v", timeout)
	}
}
