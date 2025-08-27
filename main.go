package main

import (
	"fmt"
)

func main() {
	var workerAmount int
	fmt.Scan(&workerAmount)
	dataflow := make(chan any, workerAmount)
	fmt.Println("to stop the flow write anything but numbers")

	for i := 1; i <= workerAmount; i++ {
		go worker(i, dataflow)
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

func worker(id int, jobs chan any) {

	for v := range jobs {
		fmt.Printf("воркер %d получил работу %v\n ", id, v)
	}
}
