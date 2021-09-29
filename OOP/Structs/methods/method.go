package main

import "fmt"

// Define a Trade of a trade for the stocks
// this is a new data structure type
type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

// Value returns the trade value
func (t *Trade) Value() float64 {
	//using * for pointer receiver to let the method edit the value
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}

func main() {
	t1 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.89,
		Buy:    true,
	}

	fmt.Println(t1.Value())
}
