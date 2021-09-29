package main

import (
	"fmt"
)

func main() {

	// all elements in slice must have the same type

	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	fmt.Println("---------")

	for i := range loons {
		fmt.Println(i)

	}

	fmt.Println("---------")

	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}
	for _, name := range loons {
		fmt.Printf("%s\n", name)
	}
}
