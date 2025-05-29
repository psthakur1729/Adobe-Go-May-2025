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

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

type LineItem struct {
	Item  *Product
	Units int
}

func NewLineItem() *LineItem {
	return &LineItem{}
}

func (li *LineItem) For(p *Product) *LineItem {
	li.Item = p
	return li
}

func (li *LineItem) Count(units int) *LineItem {
	li.Units = units
	return li
}

func (li LineItem) Format() string {
	return fmt.Sprintf("%s, Units = %d", li.Item.Format(), li.Units)
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

func (sc *ShoppingCart) WithLineItem(li *LineItem) *ShoppingCart {
	sc.LineItems = append(sc.LineItems, li)
	return sc
}

func (sc ShoppingCart) Total() float64 {
	var total float64
	for _, li := range sc.LineItems {
		total += li.GetValue()
	}
	return total
}

func (sc ShoppingCart) Format() string {
	var sb strings.Builder
	for _, li := range sc.LineItems {
		sb.WriteString(li.Format())
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("Total : %0.2f\n", sc.Total()))
	return sb.String()
}

func main() {
	p1 := NewProduct(100, "Pen", 10)
	p2 := NewProduct(101, "Pencil", 5)
	p3 := NewProduct(102, "Marker", 50)

	cart := NewShoppingCart().
		WithLineItem(NewLineItem().For(p1).Count(5)).
		WithLineItem(NewLineItem().For(p2).Count(50)).
		WithLineItem(NewLineItem().For(p3).Count(5))

	fmt.Println(cart.Format())
}
