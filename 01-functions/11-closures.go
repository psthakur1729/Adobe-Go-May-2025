package main

import "fmt"

func main() {
	increment := getIncrementor()
	for range 10 {
		fmt.Println(increment())
	}

	// x = 200

	for range 10 {
		fmt.Println(increment())
	}
}

func getIncrementor() func() int { // step : 1
	var x int                   // step : 2
	incrementor := func() int { //step : 3
		x += 1 // step : 4
		return x
	}
	return incrementor // step : 5
}

func fx() {
	var x int
}
