/*
make execution of printIfPrime() concurrent
Print the total number of prime numbers found in the main() function

DO NOT print the prime in the "checkPrime()"
Collect all the prime numbers and print them in the "main()"
*/
package main

import (
	"fmt"
)

func main() {
	for no := 2; no <= 100; no++ {
		checkPrime(no)
	}
	fmt.Println("Done")
}

func checkPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
}
