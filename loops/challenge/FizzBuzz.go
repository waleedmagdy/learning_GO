package main

import (
	"fmt"
)

// Challenge
// print number from 1 to 20 but if the number is divisible by 3 print fizz and if the number is divisible by 5 print buzz
// if the number is divisible by both print fizz buzz

// this is my of finding the solution

// func main() {
// 	for i := 1; i < 21; i++ {

// 		if i%3 == 0 && i%5 == 0 {
// 			fmt.Println("fizz buzz")
// 			continue
// 		}

// 		if i%3 == 0 {
// 			fmt.Println("fizz")
// 			continue
// 		}

// 		if i%5 == 0 {
// 			fmt.Println("buzz")
// 			continue
// 		}

// 		fmt.Println(i)

// 	}
// }

// Another Way from the training

func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
