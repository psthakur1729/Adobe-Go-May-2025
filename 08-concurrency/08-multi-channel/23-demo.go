package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		ch2 <- 200
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		d3 := <-ch3
		fmt.Println("Data from ch3 :", d3)
	}()

	for range 3 {
		select {
		case d1 := <-ch1:
			fmt.Println(d1)
		case d2 := <-ch2:
			fmt.Println(d2)
		case ch3 <- 300:
			fmt.Println("[select-case] data sent to ch3")
			/* default:
			fmt.Println("No channel operations where successful!") */
		}

	}
}
