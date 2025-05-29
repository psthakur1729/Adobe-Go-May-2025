package main

import "fmt"

/* product => id, name, cost */

func main() {
	/*
		var p struct {
			Id   int
			Name string
			Cost float64
		}
		p.Id = 100
		p.Name = "Pen"
		p.Cost = 5
	*/

	var p struct {
		Id   int
		Name string
		Cost float64
	} = struct {
		Id   int
		Name string
		Cost float64
	}{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)

	fmt.Println("Before applying discount")
	fmt.Println(Format(p))
	ApplyDiscount( /* ? */ )
	fmt.Println("After applying discount")
	fmt.Println(Format(p))

}

/* Write the following functions
1. Format(product) => "Id = 100, Name = Pen, Cost = 10"
2. ApplyDiscount(product, discountPercentage) => no return result. updates the cost after applying the discount
*/
