package main

import (
	"fmt"
	"time"
)

// share memory by communicating

func main() {

	stopCh := time.After(5 * time.Second)
	ch := genNos(stopCh)
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func genNos(stopCh <-chan time.Time) <-chan int {
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
