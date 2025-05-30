package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := generateNos()
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(1 * time.Second)
			fmt.Println(data)
			continue
		}
		fmt.Println("[main] channel closed!")
		break
	}

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
