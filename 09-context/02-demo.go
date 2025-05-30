package main

import (
	"context"
	"fmt"
	"time"
)

// share memory by communicating

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	ch := genNos(ctx)
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func genNos(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
	NO_LOOP:
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("[genNos] - stop signal received")
				break NO_LOOP
			default:
				time.Sleep(500 * time.Millisecond)
				ch <- i * 10
			}
		}
		close(ch)
	}()
	return ch
}
