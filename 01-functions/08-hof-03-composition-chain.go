package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// ver 1.0
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	// ver 2.0
	/*
		logAdd := wrapLogger(add)
		logAdd(100, 200)

		logSubtract := wrapLogger(subtract)
		logSubtract(100, 200)
	*/
	/*
		logAdd := wrapLogger(add)
		profileAdd := wrapProfiler(logAdd)

		logSubtract := wrapLogger(subtract)
		profileSubtract := wrapProfiler(logSubtract)

		profileAdd(100, 200)
		profileSubtract(100, 200)
	*/

	/*
		add := wrapProfiler(wrapLogger(add))
		subtract := wrapProfiler(wrapLogger(subtract))
	*/

	add := compose(add, wrapProfiler, wrapLogger)
	subtract := compose(subtract, wrapProfiler, wrapLogger)

	add(100, 200)
	subtract(100, 200)
}

// ver 1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

// ver 2.0

func wrapLogger(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started")
		oper(x, y)
		log.Println("Operation completed")
	}
}

// ver 3.0
func wrapProfiler(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Since(start)
		fmt.Printf("Time taken : %v\n", elapsed)
	}
}

/*
func compose(fn func(int, int), compsers ...func(func(int, int)) func(int, int)) func(int, int) {
	for i := len(compsers) - 1; i >= 0; i-- {
		fn = compsers[i](fn)
	}
	return fn
}
*/

type Operation func(int, int)
type WrapOperation func(Operation) Operation

func compose(fn Operation, compsers ...WrapOperation) Operation {
	for i := len(compsers) - 1; i >= 0; i-- {
		fn = compsers[i](fn)
	}
	return fn
}
