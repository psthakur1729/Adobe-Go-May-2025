package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := generateNos()
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("[main] channel closed!")

}

func generateNos() <-chan int {
	ch := make(chan int)
	count := rand.Intn(20)
	fmt.Println("[generateNos] count = ", count)
	go func() {
		for i := range count {
			ch <- (i + 1) * 10
		}
		close(ch)
	}()
	return ch
}
