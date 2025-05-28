package main

import "fmt"

func main() {
	var x int
	x = 100

	// Data -> Address
	var xPtr *int
	xPtr = &x // &x is "address of x"
	fmt.Println(xPtr)

	// Address -> Data (deferencing)
	var data int
	data = *xPtr
	fmt.Println(data)

	// in other words
	fmt.Println(x == *(&x))

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap()
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func swap( /* ? */ ) /* no return values */ {
	/* ? */
}
