package main

import "fmt"

func main() {

	nums := [5]int{1, 2, 3, 4, 5}
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		defer close(ch1)
		for _, v := range nums {
			ch1 <- v
		}

	}()
	go func() {
		defer close(ch2)
		for val := range ch1 {
			ch2 <- val * 2
		}
	}()

	for result := range ch2 {

		fmt.Println(result)
	}
}
