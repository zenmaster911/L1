package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	clicker int64
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	var incrementer int

	for i := 0; i < 4; i++ {
		wg.Add(2)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&counter.clicker, 1)
			}
		}(&wg)
		go withMutex(&incrementer, &wg)
	}
	wg.Wait()
	fmt.Println(counter.clicker, incrementer)
}

func withMutex(num *int, wg *sync.WaitGroup) {
	defer wg.Done()
	var mx sync.Mutex
	for i := 0; i < 2000; i++ {
		mx.Lock()
		*num++
		mx.Unlock()
	}
}
