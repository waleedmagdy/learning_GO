package main

import (
	"fmt"
)

func main() {
	//var x float64 // this variable changed from int to float64 to check the mean result
	//var y float64

	x, y := 1.0, 2.0 // you can declere two variables in the same line
	//y := 2.0

	fmt.Printf("x=%v, type of %T\n", x, x) //Printf >> print formated text %v will refer to the variable value
	fmt.Printf("y=%v, type of %T\n", y, y) // while %T will refer to the variable type

	//var mean float64
	mean := (x + y) / 2.0 // calculating the mean
	fmt.Printf("result: %v, type of %T\n", mean, mean)

}
