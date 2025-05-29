package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling the execution of f1() through the scheduler to be executed in FUTURE
	f2()
	// block the execution of this function so that the scheduler can look for other goroutines that are schedule and execute them

	// DO NOT use the below poor man's synchronization techniques
	time.Sleep(2 * time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
