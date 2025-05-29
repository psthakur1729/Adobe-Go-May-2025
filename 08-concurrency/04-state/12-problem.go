/*
make execution of printIfPrime() concurrent
Print the total number of prime numbers found in the main() function
*/
package main

import (
	"fmt"
)

func main() {
	for no := 2; no <= 100; no++ {
		printIfPrime(no)
	}
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
