package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	nums := [5]int{2, 4, 6, 8, 10}
	for _, v := range nums {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Println(v * v)
		}(v)
	}
	wg.Wait()
}
