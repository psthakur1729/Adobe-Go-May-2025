package models

import "testing"

func Test_ShoppingCart(t *testing.T) {
	p1 := NewProduct(100, "Pen", 10)
	p2 := NewProduct(101, "Pencil", 5)
	p3 := NewProduct(102, "Marker", 50)

	cart := NewShoppingCart().
		WithLineItem(NewLineItem().For(p1).Count(5)).
		WithLineItem(NewLineItem().For(p2).Count(50)).
		WithLineItem(NewLineItem().For(p3).Count(5))
	cartTotal := cart.Total()
	if cartTotal != 550 {
		t.Errorf("Expected : 550, actual : %v", cartTotal)
	}

}
