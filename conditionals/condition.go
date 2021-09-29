package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 5 {
		fmt.Println("X is bigger than 5")
	}

	if x > 100 {
		fmt.Println("X is Very Big")
	} else {
		fmt.Println("X is not that big")
	}

	if x > 5 && x < 15 { // it is AND Logical Operator
		fmt.Println("X is just right")
	}

	if x < 20 || x > 30 { // it is OR logical Operator
		fmt.Println("X is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 { // define a variable in the same line of if statement
		fmt.Println("a is more than half of b")
	}

}
