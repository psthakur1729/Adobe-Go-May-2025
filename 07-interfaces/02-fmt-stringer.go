/*
Write apprpriate methods for ShoppingCart, LineItem & Product & factory functions
to simplify the usage of these types
*/

package main

import (
	"fmt"
	"strings"
)

/* product => id, name, cost */
type Product struct {
	Id   int
	Name string
	Cost float64
}

func NewProduct(id int, name string, cost float64) *Product {
	return &Product{
		Id:   id,
		Name: name,
		Cost: cost,
	}
}

// fmt.Stringer interface implementation
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

type LineItem struct {
	Item  *Product
	Units int
}

func NewLineItem(p *Product, units int) *LineItem {
	return &LineItem{
		Item:  p,
		Units: units,
	}
}

func (li LineItem) String() string {
	return fmt.Sprintf("%s, Units = %d", li.Item, li.Units)
}

func (li LineItem) GetValue() float64 {
	return li.Item.Cost * float64(li.Units)
}

type ShoppingCart struct {
	LineItems []*LineItem
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (sc *ShoppingCart) AddLineItem(li *LineItem) {
	sc.LineItems = append(sc.LineItems, li)
}

func (sc ShoppingCart) Total() float64 {
	var total float64
	for _, li := range sc.LineItems {
		total += li.GetValue()
	}
	return total
}

// fmt.Stringer interface implementation
func (sc ShoppingCart) String() string {
	var sb strings.Builder
	for _, li := range sc.LineItems {
		sb.WriteString(fmt.Sprint(li))
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("Total : %0.2f\n", sc.Total()))
	return sb.String()
}

func main() {
	p1 := NewProduct(100, "Pen", 10)
	p2 := NewProduct(101, "Pencil", 5)
	p3 := NewProduct(102, "Marker", 50)

	cart := NewShoppingCart()
	cart.AddLineItem(NewLineItem(p1, 10))
	cart.AddLineItem(NewLineItem(p2, 50))
	cart.AddLineItem(NewLineItem(p3, 5))

	fmt.Println(cart)
}
