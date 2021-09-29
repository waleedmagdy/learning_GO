package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := &Point{1, 2}
	p.Move(2, 3)
	fmt.Printf("%+v\n", p)
}

// func zero(xPtr *int) {
// 	*xPtr = 0
// }
// func main() {
// 	x := 5
// 	zero(&x)
// 	fmt.Println(x) // x is 0
// }
