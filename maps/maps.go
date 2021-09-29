package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61, // you must have a trailing comma in multi line
	}

	// print the number of the items
	fmt.Println(len(stocks))

	// Get a Value
	fmt.Println(stocks["MSFT"])

	// Get a Zero value if not found
	fmt.Println(stocks["TSLA"]) // 0

	// Use two value to see if found
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	// Set a value
	stocks["TSLA"] = 322.12
	fmt.Println(stocks)

	// delete a value
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// Single value "for" is on keys
	for key := range stocks {
		fmt.Println(key)
	}

	// Double value "for" is key, value
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}
