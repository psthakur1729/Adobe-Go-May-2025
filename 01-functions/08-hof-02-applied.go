package main

import (
	"fmt"
	"log"
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

	// Updated ver 1.0
	add := wrapLogger(add)
	subtract := wrapLogger(subtract)

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
