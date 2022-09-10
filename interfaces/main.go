package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Square struct {
	width float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.width * s.width
}

func (s Square) perimeter() float64 {
	return 4 * s.width
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return (2 * r.width) + (2 * r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	s := Square{width: 3}
	fmt.Println("Square area: ", s.area())
	fmt.Println("Square perimeter: ", s.perimeter())

	r := Rectangle{width: 3, height: 4}
	fmt.Println("Rectangle area: ", r.area())
	fmt.Println("Rectangle perimeter: ", r.perimeter())

	c := Circle{radius: 5}
	fmt.Println("Circle area: ", c.area())
	fmt.Println("Circle perimeter: ", c.perimeter())
}
