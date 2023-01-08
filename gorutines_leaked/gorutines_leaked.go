package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// why goroutines leaked? find problem and fix it
func main() {
	ctx, _ := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		go work(ctx)
	}
	fmt.Printf("Goroutines running: %d\n", runtime.NumGoroutine())

	//cancel()

	// server doesn't die. Imagine, it's doing useful work.
	for {
		fmt.Printf("Goroutines leaks: %d\n", runtime.NumGoroutine()-1)
		//fmt.Printf("i do some useful work, print num: %d\n", rand.Int())
		time.Sleep(time.Second)
	}
}

func work(ctx context.Context) {
	ch := make(chan int)

	go func() {
		select {
		case ch <- rpcCall():
			fmt.Println()
			return
		case <-ctx.Done():
			return
		}
	}()

	for {
		select {
		case value := <-ch:
			fmt.Printf("result of rpcCall: %d\n", value)
			return
		case <-ctx.Done():
			return
		}
	}
}

func rpcCall() int {
	<-time.After(time.Second * 10)
	return rand.Int()
}
