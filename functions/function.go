package main

import "fmt"

// func to add two numbers
func add(a int, b int) int {
	return a + b
}

//fun to divide two numbers and return value and reminder
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	val := add(5, 6)
	fmt.Println(val)

	div, mod := divmod(10, 6)
	fmt.Printf("div=%d ,mod=%d\n", div, mod)
}
