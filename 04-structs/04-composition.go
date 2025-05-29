package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	Product
	// Dummy
	Expiry string
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

	// milk := PerishableProduct{Product{200, "Milk", 48}, "2 Days"}

	fmt.Println(milk.Expiry)
	// fmt.Println(milk.Product.Id, milk.Product.Name, milk.Product.Cost)
	fmt.Println(milk.Id, milk.Name, milk.Cost)
}
