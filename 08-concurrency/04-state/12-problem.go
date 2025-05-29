/*
make execution of printIfPrime() concurrent
Print the total number of prime numbers found in the main() function
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var primeCount int64

func main() {
	wg := &sync.WaitGroup{}
	for no := 2; no <= 100; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			printIfPrime(no)
		}()
	}
	wg.Wait()
	fmt.Println("Done")
	fmt.Println("Prime Count : ", primeCount)
}

func printIfPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	atomic.AddInt64(&primeCount, 1)
	fmt.Println("Prime No :", no)
}
