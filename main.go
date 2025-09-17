package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println(i + 1)
			Sleep(2)
		}
	}()
	wg.Wait()
}

func Sleep(duration time.Duration) {
	timer := time.NewTimer(duration * time.Second)
	<-timer.C
}
