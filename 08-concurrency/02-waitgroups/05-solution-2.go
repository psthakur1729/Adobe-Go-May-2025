package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	for no := 2; no <= 100; no++ {
		wg.Add(1)
		go func() {
			printIfPrime(no)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Done")
}

func printIfPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
}
