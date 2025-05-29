/*
make execution of printIfPrime() concurrent
Print the total number of prime numbers found in the main() function

DO NOT print the prime in the "checkPrime()"
Collect all the prime numbers and print them in the "main()"
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
		go func() {
			defer wg.Done()
			checkPrime(no)
		}()
	}
	wg.Wait()
	fmt.Println("Done")
	for _, primeNo := range primes {
		fmt.Println("Prime No :", primeNo)
	}
	fmt.Println("Prime Count :", primeCount)
}

func checkPrime(no int) {
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
