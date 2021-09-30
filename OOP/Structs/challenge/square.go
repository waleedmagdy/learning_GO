package main

import (
	"fmt"
	"log"
)

//Point Type
type Point struct {
	X int
	Y int
}

//Square type with Center points and Length
type Square struct {
	Center Point
	Length int
}

//NewSquare to create new square object and validate Length
func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0 (was %d)", length)
	}
	s := &Square{
		Center: Point{x, y},
		Length: length,
	}
	return s, nil
}

//Move the Square points
func (s *Point) Move(dx int, dy int) {
	s.X += dx
	s.Y += dy
}

//Calculat the Are of the square
func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("Error: can't create square")
	}

	s.Center.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())

}
