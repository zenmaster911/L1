package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var workerAmount int
	fmt.Scan(&workerAmount)
	//activeWorkers := new(sync.Map)
	wg.Add(workerAmount)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataflow := make(chan any, workerAmount)
	result := make(chan any, workerAmount)
	defer close(dataflow)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("to stop the flow press ctrl+C")

	for i := 1; i <= workerAmount; i++ {
		go worker(ctx, &wg, i, dataflow, result)
	}

	go func() {

		<-sig
		fmt.Println("termination signal received")
		cancel()

	}()

	go func() {
		var data int
		defer close(dataflow)
		for {
			_, err := fmt.Scan(&data)
			if err != nil {
				fmt.Printf("неверный формат ввода %v", err)
				break
			}
			select {
			case dataflow <- data:
			case <-ctx.Done():
				return
			}
		}
	}()

	for {
		select {
		case res := <-result:
			fmt.Printf("result is %v\n", res)
		case <-ctx.Done():
			fmt.Println("waiting for the workers")
			wg.Wait()
			close(result)
			fmt.Println("shutting down")
			return
		}

	}
}

// func ActiveWorkerAwait(m *sync.Map) {
// 	m.Range(func(key any, value any) bool {
// 		workerId := key.(int)
// 		ch := value.(chan struct{})
// 		<-ch
// 		fmt.Printf("worker %d finished it's work\n", workerId)
// 		return true
// 	})

// }

func worker(ctx context.Context, wg *sync.WaitGroup, id int, jobs, result chan any) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("воркер %d получил работу %v\n ", id, job)
			time.Sleep(time.Second * 6)
			result <- job

		}
	}
}
