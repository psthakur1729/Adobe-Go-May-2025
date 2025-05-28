package main

import "fmt"

func main() {
	// generate prime numbers from 2 to 100 and print them
	// main() => printing the generated prime numbers
	// generatePrimes(start, end) => returns the prime numbers
	// isPrime(no) => checks if the given number is a prime number

	primes := generatePrimes(2, 100)
	for _, primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
}

func generatePrimes(start, end int) []int {
	var primes []int
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
