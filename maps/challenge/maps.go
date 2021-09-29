package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles Needles Needles Needles Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	words := strings.Fields(text) // this will ignore the spaces and create a slice with only each word
	counts := map[string]int{}    // Empty map
	//fmt.Println(counts[""])
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)
}
