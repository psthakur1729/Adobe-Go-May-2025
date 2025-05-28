package main

import "fmt"

func main() {
	func() {
		fmt.Println("anonymous fn")
	}()

	func(x, y int) {
		fmt.Println(x * y)
	}(100, 200)

	q, r := func(x, y int) (q, r int) {
		q, r = x/y, x%y
		return
	}(100, 7)
	fmt.Println(q, r)
}
