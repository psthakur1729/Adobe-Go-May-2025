package models

import (
	"fmt"
	"strings"
)

type ShoppingCart struct {
	lineItems []*LineItem
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (sc *ShoppingCart) WithLineItem(li *LineItem) *ShoppingCart {
	sc.lineItems = append(sc.lineItems, li)
	return sc
}

func (sc ShoppingCart) Total() float64 {
	var total float64
	for _, li := range sc.lineItems {
		total += li.GetValue()
	}
	return total
}

func (sc ShoppingCart) Format() string {
	var sb strings.Builder
	for _, li := range sc.lineItems {
		sb.WriteString(li.Format())
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("Total : %0.2f\n", sc.Total()))
	return sb.String()
}
