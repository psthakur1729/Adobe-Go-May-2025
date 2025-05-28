package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be zero")

func main() {
	defer func() {
		fmt.Println("		[main] - deferred")
		if err := recover(); err != nil {
			fmt.Println("	[main] shutdown due to panic, error :", err)
		}
	}()
	divisor := 0
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		// q, r := divide(100, divisor)
		q, r, err := divideAdapter(100, divisor)
		if err != nil {
			fmt.Println("error :", err)
			continue
		}
		fmt.Printf("[main] q = %d, r = %d\n", q, r)
		break
	}
}

func divideAdapter(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party library function (not allowed to change the code)
func divide(x, y int) (quotient, remainder int) {
	defer fmt.Println("		[divide] - deferred")
	fmt.Println("[divide] - calculating quotient")
	if y == 0 {
		panic(ErrDivideByZero)
	}

	quotient = x / y
	fmt.Println("[divide] - calculating remainder")
	remainder = x % y
	return
}
