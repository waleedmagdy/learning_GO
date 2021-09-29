package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value (%0.2f)", n)
	} else {
		return math.Sqrt(n), nil
	}
}

func main() {
	s1, err := sqrt(10)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("sqrt for s1 is: %0.2f\n", s1)
	}

	s2, err := sqrt(-10)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("sqrt for s1 is: %0.2f\n", s2)
	}
}
