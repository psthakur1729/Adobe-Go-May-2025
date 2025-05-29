package main

import (
	"fmt"

	"github.com/tkmagesh/Adobe-Go-May-2025/06-shoppingcart-app/models"
)

func main() {
	p1 := models.NewProduct(100, "Pen", 10)
	p2 := models.NewProduct(101, "Pencil", 5)
	p3 := models.NewProduct(102, "Marker", 50)

	cart := models.NewShoppingCart().
		WithLineItem(models.NewLineItem().For(p1).Count(5)).
		WithLineItem(models.NewLineItem().For(p2).Count(50)).
		WithLineItem(models.NewLineItem().For(p3).Count(5))

	fmt.Println(cart.Format())
}
