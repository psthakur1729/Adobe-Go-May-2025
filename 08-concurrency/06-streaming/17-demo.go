package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generateNos()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func generateNos() <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- 10
		time.Sleep(1 * time.Second)
		ch <- 20
		time.Sleep(1 * time.Second)
		ch <- 30
		time.Sleep(1 * time.Second)
		ch <- 40
		time.Sleep(1 * time.Second)
		ch <- 50
	}()
	return ch
}
