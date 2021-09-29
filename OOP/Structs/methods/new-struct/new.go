package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string  // Stock Symbol
	Volume int     // number of shares
	Price  float64 // trade price
	Buy    bool    // true if buy trade, false if sell trade
}

// NewTrade will create a new object trade and will validate inputs
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol can't be empty")
	}
	if volume <= 0 {
		return nil, fmt.Errorf("volume must be > 0 (was %d)", volume)
	}
	if price <= 0.0 {
		return nil, fmt.Errorf("Price must be > 0 (was %f)", price)
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}

	return trade, nil
}

// a receiver of the Trade data pointer to get the value of the trade
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {

	t1, err := NewTrade("MSFT", 10, 99.89, true)
	if err != nil {
		fmt.Printf("error: can't create trade - %s\n", err)
		os.Exit(1)
	}
	fmt.Println(t1.Value())
}
