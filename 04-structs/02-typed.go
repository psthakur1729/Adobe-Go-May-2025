package main

import "fmt"

/* product => id, name, cost */
type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	/*
		var p Product
		p.Id = 100
		p.Name = "Pen"
		p.Cost = 5
	*/

	var p Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	// Below is not advisable
	// var p Product = Product{100, "Pen", 5}

	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)

	// struct instances are values
	var p1 = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	var p2 = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	fmt.Println("p1 == p2 :", p1 == p2)

	fmt.Println("Before applying discount")
	fmt.Println(Format(p))
	ApplyDiscount(&p, 10)
	fmt.Println("After applying discount")
	fmt.Println(Format(p))

}

/* Write the following functions
1. Format(product) => "Id = 100, Name = Pen, Cost = 10"
2. ApplyDiscount(product, discountPercentage) => no return result. updates the cost after applying the discount
*/

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *Product, discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
