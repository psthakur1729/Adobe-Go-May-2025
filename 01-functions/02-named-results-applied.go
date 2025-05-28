package main

import (
	"errors"
	"fmt"
)

func main() {
	divisor := 0
	/*
		q, r, err := divide(100, divisor)
		if err != nil {
			fmt.Println("error :", err)
			return
		}
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	*/

	if q, r, err := divide(100, divisor); err != nil {
		fmt.Println("error :", err)
		return
	} else {
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	}
}

/*
func divide(multiplier, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, errors.New("divisor cannot be zero")
	}
	quotient := multiplier / divisor
	remainder := multiplier % divisor
	return quotient, remainder, nil
}
*/

func divide(multiplier, divisor int) (quotient, remainder int, err error) {
	if divisor == 0 {
		err = errors.New("divisor cannot be zero")
		return
	}
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return
}
