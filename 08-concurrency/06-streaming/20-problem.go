/*
Modify the below using "Share memory by communicating"
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	primesCh := make(chan int)
	for no := 2; no <= 100; no++ {
		wg.Add(1)
		go checkPrime(wg, no, primesCh)
	}
	go func() {
		for primeNo := range primesCh {
			fmt.Println("Prime No :", primeNo)
		}
	}()
	wg.Wait()
	close(primesCh)
	fmt.Println("Done")
}

func checkPrime(wg *sync.WaitGroup, no int, ch chan int) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	ch <- no

}
