package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	var N int64 //lifetime of this app
	fmt.Scan(&N)
	ch := make(chan string)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(N)*time.Second))
	defer cancel()

	go func() { //routine to receive data
		for {

			select {
			case <-ctx.Done():
				return

			default:
				var data string
				fmt.Fscanln(os.Stdin, &data)
				fmt.Println("data received")
				ch <- data

			}
		}
	}()
	Worker(ch, ctx)

}

func Worker(ch chan string, ctx context.Context) error {

	for {

		select {
		case data := <-ch:
			fmt.Println("magic was made with", data)
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return ctx.Err()
		}

	}

}
