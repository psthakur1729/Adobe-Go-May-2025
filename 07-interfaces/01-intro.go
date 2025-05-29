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

// v4.0
/*
Add Perimeter() to both Circle (2 * Pi * r) & Rectangle (2 * (L + B))
Write a PrintPermeter() function
*/

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

type PerimeterFinder interface {
	Perimeter() float64
}

func PrintPermeter(x PerimeterFinder) {
	fmt.Println("Perimter :", x.Perimeter())
}

// interface composition
type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShapeStats(x ShapeStatsFinder) {
	PrintArea(x)     // x should implement AreaFinder
	PrintPermeter(x) // x should implement PerimeterFinder
}

func main() {
	c := Circle{Radius: 16}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPermeter(c)
	*/
	PrintShapeStats(c)

	r := Rectangle{Length: 14, Breadth: 16}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPermeter(r)
	*/
	PrintShapeStats(r)
}
