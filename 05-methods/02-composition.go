package main

import "fmt"

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

type PerishableProduct struct {
	Product
	// Dummy
	Expiry string
}

// overriding the Format() method of the Product
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}
func main() {
	/*
		milk := PerishableProduct{
			Expiry: "2 Days",
			Product: Product{
				Id:   200,
				Name: "Milk",
				Cost: 48,
			},
		}
	*/
	milk := NewPerishableProduct(200, "Milk", 48, "2 Days")

	fmt.Println(milk.Format())
	milk.ApplyDiscount(20)
	fmt.Println(milk.Format())

}
