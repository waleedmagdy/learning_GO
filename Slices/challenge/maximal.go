package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 10, 4, 23, 15}
	max := nums[0]
	// we use slice notation _ to skip the first value
	for _, value := range nums {
		if value > max {
			max = value
		}
	}

	fmt.Println(max)
}
