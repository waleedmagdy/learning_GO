package main

import "fmt"

// Defer
// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

func clenup(name string) {
	fmt.Printf("Cleaning UP %s\n", name)
}

func worker() {
	defer clenup("A")
	defer clenup("B")

	fmt.Println("Worker")
}

func main() {
	worker()
}
