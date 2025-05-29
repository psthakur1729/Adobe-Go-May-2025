package models

import "fmt"

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
