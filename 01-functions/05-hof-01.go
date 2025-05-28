package main

import "fmt"

type ArithmaticOperation func(int, int) int

func main() {
	// assign functions as values (data) to variables
	// var oper func(int, int) int
	// var oper ArithmaticOperation

	oper := func(x, y int) int {
		return x + y
	}
	fmt.Println(oper(100, 200))

	oper = func(i1, i2 int) int {
		return i1 * i2
	}
	fmt.Println(oper(100, 200))
}
