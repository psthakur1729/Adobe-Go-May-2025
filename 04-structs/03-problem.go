/*
Introduce a ShoppingCart
	Shopping Cart
		has collection of LineItem
	LineItem
		Product
		Units

Write the following functions
	PrintCart()
	CalculateTotal()
*/

package main

import "fmt"

/* product => id, name, cost */
type Product struct {
	Id   int
	Name string
	Cost float64
}

type LineItem struct {
	Item  Product
	Units int
}

type ShoppingCart struct {
	LineItems []LineItem
}

func main() {
	p1 := Product{Id: 100, Name: "Pen", Cost: 10}
	p2 := Product{Id: 101, Name: "Pencil", Cost: 5}
	p3 := Product{Id: 102, Name: "Marker", Cost: 50}
	l1 := LineItem{Item: p1, Units: 20}
	l2 := LineItem{Item: p2, Units: 50}
	l3 := LineItem{Item: p3, Units: 5}
	cart := ShoppingCart{}
	cart.LineItems = append(cart.LineItems, l1)
	cart.LineItems = append(cart.LineItems, l2)
	cart.LineItems = append(cart.LineItems, l3)

	PrintCart(cart)
	total := CalculateTotal(cart)
	fmt.Printf("Total : %0.2f\n", total)
}

func FormatProduct(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func FormatLineItem(li LineItem) string {
	return fmt.Sprintf("%s, Units = %d", FormatProduct(li.Item), li.Units)
}

func PrintCart(cart ShoppingCart) {
	for _, li := range cart.LineItems {
		fmt.Println(FormatLineItem(li))
	}
}

func CalculateTotal(cart ShoppingCart) float64 {
	var total float64
	for _, li := range cart.LineItems {
		total += float64(li.Units) * li.Item.Cost
	}
	return total
}
