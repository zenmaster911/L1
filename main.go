package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	deadline, deadlineCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(10)*time.Second))
	defer deadlineCancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Printf("\ntermination signal received\n")
		cancel()
	}()

	go withRuntime()
	go ctxDeadline(deadline)
	go withCondition()
	go func() {
		defer func() {
			recover()
		}()
		badDecision()
	}()
	go withCtx(ctx)

	stopCh := make(chan struct{})
	individualStopCh1 := make(chan struct{})
	individualStopCh2 := make(chan struct{})
	commonStopCh := make(chan int)
	go withStopCh(stopCh)
	go individualStop1(individualStopCh1)
	go individualStop2(individualStopCh2)
	for i := 1; i <= 2; i++ {
		go commonStop(i, commonStopCh)
	}

	for {
		fmt.Print("\nSelect which routine would you like to cancel:\n\t1) usual stop chanel\n\t2) individual routine stop Channel(1 or 2)\n\t",
			"3) common routine stop channel\n\t4)stop main\n")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Printf("\n\nyour choice might be incorrect this is the problem-> %v\nTry one more time\n\n", err)
			continue
		}
		switch choice {
		case 1:
			stopCh <- struct{}{}
		case 2:
			fmt.Printf("\nselect which routine would you like to stop first(press 1) or second(press 2)\n")
			var subchoice int
			_, err := fmt.Scan(&subchoice)
			if err != nil {
				fmt.Printf("\n\nyour choice might be incorrect this is the problem-> %v\nTry one more time\n\n", err)
				continue
			}
			switch subchoice {
			case 1:
				individualStopCh1 <- struct{}{}
			case 2:
				individualStopCh2 <- struct{}{}
			default:
				fmt.Printf("there is no such option, good luck next time")
				continue
			}
		case 3:
			fmt.Printf("\nselect which routine would you like to stop first(press 1) or second(press 2)\n")
			var subchoice int
			_, err := fmt.Scan(&subchoice)
			if err != nil {
				fmt.Printf("\n\nyour choice might be incorrect this is the problem-> %v\nTry one more time\n\n", err)
				continue
			}
			if subchoice != 1 && subchoice != 2 {
				fmt.Printf("\n\nyour choice might be incorrect this is the problem-> %v\nTry one more time\n\n", err)
				continue
			}
			commonStopCh <- subchoice
		case 4:
			return
		}
	}

}

func withRuntime() {
	fmt.Println("withRuntime routine stopped\n")
	runtime.Goexit()
}

func ctxDeadline(ctx context.Context) {
	start := time.Now()
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err(), "ctxDeadline routine stopped\n")
			return
		default:
			fmt.Printf("time passed:%v \n", time.Since(start))
			time.Sleep(5 * time.Second)
		}
	}
}

func withCondition() {
	i := 0
	for {

		i++
		time.Sleep(time.Second)
		if i == 7 {
			fmt.Println("withCondition routine stopped\n")
			return
		}
	}
}

func withStopCh(stopCh chan struct{}) {
	for {
		seconds := 0
		seconds++
		select {
		case <-stopCh:
			fmt.Printf("withStopCh routine stopped\ntime pssed: %d", seconds)
			close(stopCh)
			return
		default:
			time.Sleep(time.Second)
		}
	}
}

func badDecision() {

	panic("bad Decision routine stopped \nBut I've heard it's not the best way to stop it")
}

func withCtx(ctx context.Context) {

	<-ctx.Done()
	fmt.Println("withCtx routine stopped")

}

func individualStop1(ch chan struct{}) {
	<-ch
	fmt.Println("individualStop1 routine stopped")
	close(ch)
}

func individualStop2(ch chan struct{}) {
	<-ch
	fmt.Println("individualStop2 routine stopped")
	close(ch)
}

func commonStop(id int, ch chan int) {
	for {
		select {
		case number := <-ch:
			if number == id {
				fmt.Printf("commonStop routine #%d stopped", id)
				return
			} else {
				continue
			}
		default:
			//fmt.Printf("commonStop routine #%d doing BzZzzZZ", id)
			time.Sleep(10 * time.Second)
		}

	}
}
