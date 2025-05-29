package main

import "fmt"

/* product => id, name, cost */
type Product struct {
	Id   int
	Name string
	Cost float64
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func main() {

	var p Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	fmt.Printf("%+v\n", p)

	fmt.Println("Before applying discount")
	// fmt.Println(Format(p))
	fmt.Println(p.Format())

	// ApplyDiscount(&p, 10)
	p.ApplyDiscount(10)
	fmt.Println("After applying discount")
	// fmt.Println(Format(p))
	fmt.Println(p.Format())

}

/* Write the following functions
1. Format(product) => "Id = 100, Name = Pen, Cost = 10"
2. ApplyDiscount(product, discountPercentage) => no return result. updates the cost after applying the discount
*/
