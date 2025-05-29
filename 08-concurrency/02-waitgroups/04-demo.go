package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	wg := &sync.WaitGroup{}
	flag.IntVar(&count, "count", 0, "Number of goroutines to spin up!")
	flag.Parse()
	fmt.Printf("Spinning up %d goroutines.... Hit ENTER to start!\n", count)
	fmt.Scanln()
	for id := range count {
		wg.Add(1)       // increment the counter by 1
		go fn(id+1, wg) // scheduling the execution of f1() through the scheduler to be executed in FUTURE
	}
	wg.Wait() // block until the counter becomes 0 (default = 0)
	fmt.Println("Done!")
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)

}
