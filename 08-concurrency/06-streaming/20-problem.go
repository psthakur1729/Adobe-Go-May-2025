/*
Modify the below using "Share memory by communicating"
*/

package main

import (
	"fmt"
	"sync"
)

// Communicate by sharing memory

var primeCount int64
var primes []int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for no := 2; no <= 100; no++ {
		wg.Add(1)
		go checkPrime(wg, no)
	}
	wg.Wait()
	fmt.Println("Done")
	for _, primeNo := range primes {
		fmt.Println("Prime No :", primeNo)
	}
	fmt.Println("Prime Count :", primeCount)
}

func checkPrime(wg *sync.WaitGroup, no int) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	mutex.Lock()
	{
		primeCount += 1
		primes = append(primes, no)
	}
	mutex.Unlock()
}
