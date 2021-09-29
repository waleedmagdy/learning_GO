package main

import (
	"fmt"
)

func main() {
	// count is 0
	count := 0

	// iterate over numbers between 1000 to 9999 for the first and second 4 digit
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ { // don't count twice so we use b := a to start iterate from a value
			n := a * b // multiply two digits with

			// check if even ended found
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				// count = count +1
				count++
			}
		}
	}

	// print the numbers
	fmt.Println(count)
}
