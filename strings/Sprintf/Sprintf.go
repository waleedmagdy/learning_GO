package main

import (
	"fmt"
)

func main() {
	n := 42
	s := fmt.Sprintf("%d", n)

	// fmt.Printf("s = %s (type: %T)\n", s, s)
	// output : s = 42 (type: string)

	fmt.Printf("s = %q (type: %T)\n", s, s) // %q, %v, %s and so on is called verb
	// output: s = "42" (type: string)

}
