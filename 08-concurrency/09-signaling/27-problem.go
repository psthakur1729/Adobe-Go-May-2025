package main

import (
	"fmt"
	"time"
)

// share memory by communicating

func main() {

	stopCh := make(chan struct{})
	ch := genNos(stopCh)
	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		close(stopCh)
	}()
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func genNos(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
	NO_LOOP:
		for i := 1; ; i++ {
			select {
			case <-stopCh:
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
