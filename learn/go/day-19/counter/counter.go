package counter

import (
	"sync"
	"sync/atomic"
)

type SafeCounter struct {
	mu    sync.Mutex
	Value int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	c.Value++
	c.mu.Unlock()
}

func (c *SafeCounter) Add(n int, workers int) {
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < n/workers; j++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()
}

type AtomicCounter struct {
	Value int64
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.Value, 1)
}

func (c *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&c.Value)
}

func AtomicAdd(c *AtomicCounter, n int, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
}
