package main

import (
	"fmt"
)

func main() {
	book := "The Colour of the Universe"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type: %T)\n", book[0], book[0])

	// Strings in GO are immutable
	// book[0] = 116  // you will get an error, you can not assign to book[0]

	// Slice (start, end) , 0 based, half emty range
	fmt.Println(book[4:11])
	// Slice no end
	fmt.Println(book[4:])
	// Slice no start
	fmt.Println(book[:4])

	// Use + to concatenate Strings
	fmt.Println("t" + book[1:])

	// Unicode
	fmt.Println("It was ¹/² Price")

	// Multi line
	poem := `
	The road goes over on
	Down frpm the door where it began writing
	...
	`
	fmt.Println(poem)
}
