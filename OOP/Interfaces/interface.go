package main

import (
	"fmt"
	"math"
)

//Square type with Center points and Length
type Square struct {
	Length float64
}

//Calculat the Are of the square
func (s *Square) Area() float64 {
	return s.Length * s.Length
}

//Circle type struct
type Circle struct {
	Radius float64
}

//Area for the circle
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func sumArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

type Shape interface {
	Area() float64
}

func main() {
	sq := Square{
		Length: 20.50,
	}

	cr := Circle{
		Radius: 30.50,
	}
	fmt.Println(sq.Area())
	fmt.Println(cr.Area())

}
