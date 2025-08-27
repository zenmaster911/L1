package main

import (
	"fmt"
	"sync"
)

func main() {
	var workerAmount int
	var wg sync.WaitGroup
	fmt.Scan(&workerAmount)
	dataflow := make(chan any, workerAmount)

	wg.Add(workerAmount)
	for i := 1; i <= workerAmount; i++ {
		go worker(i, dataflow, &wg)
	}
	// go func() {
	var data int
	for {
		_, err := fmt.Scan(&data)
		if err != nil {
			fmt.Printf("неверный формат ввода %v", err)
			break
		}
		dataflow <- data
	}
	// }()

}

func worker(id int, jobs chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range jobs {
		fmt.Printf("воркер %d получил работу %v\n ", id, v)
	}
}
