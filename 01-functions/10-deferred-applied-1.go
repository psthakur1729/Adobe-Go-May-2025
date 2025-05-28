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
	q, r := divide(100, divisor)
	fmt.Printf("[main] q = %d, r = %d\n", q, r)
}

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
