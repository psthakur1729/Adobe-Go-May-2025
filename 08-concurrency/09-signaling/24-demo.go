package main

import (
	"fmt"
	"time"
)

// communicate by sharing memory (not advisable)
var stop = false

func main() {
	ch := genNos()
	go func() {
		time.Sleep(5 * time.Second)
		stop = true
	}()
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; ; i++ {
			if stop {
				fmt.Println("[genNos] - stop signal received")
				break
			}
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
		close(ch)
	}()
	return ch
}
