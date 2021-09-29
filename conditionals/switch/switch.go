package main

import "fmt"

func main() {

	x := 20

	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("many")
	}

	switch {
	case x > 100:
		fmt.Println("X is very Big")
	case x > 10:
		fmt.Println("X is big")
	default:
		fmt.Println("X is small")

	}
}
