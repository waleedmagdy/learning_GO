package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("---------")

	for i := 0; i < 3; i++ {
		if i > 1 {
			break // it will end the iteration and exit the loop
		}
		fmt.Println(i)
	}

	fmt.Println("---------")

	for i := 0; i < 3; i++ {

		if i < 1 {
			continue // it will move to the next iteration without executing the next code
		}
		fmt.Println(i)
	}
	fmt.Println("---------")

	a := 0 //you can use for loop that way if you like to iterate over variable
	for a < 3 {
		fmt.Println(a)
		a++
	}

	fmt.Println("---------")

	b := 0
	for { // using for without a condition
		if b > 2 { // you need to increment the value used in iteration and exit condition
			break
		}
		fmt.Println(b)
		b++
	}
}
