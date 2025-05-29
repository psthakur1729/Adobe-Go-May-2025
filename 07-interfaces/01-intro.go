package main

import (
	"fmt"
	"math"
)

// v1.0
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// v2.0
type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

// v3.0
// utility function

type AreaFinder interface {
	Area() float64
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}
func main() {
	c := Circle{Radius: 16}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)

	r := Rectangle{Length: 14, Breadth: 16}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)
}
