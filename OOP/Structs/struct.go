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

func main() {
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)
	fmt.Printf("%+v\n", t1) // +v will give more info of the output and how it looks like
	fmt.Println("-----------")
	fmt.Println(t1.Symbol)

	t2 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)

}

// generated struct from json file
type AutoGenerated struct {
	Glossary struct {
		Title    string `json:"title"`
		GlossDiv struct {
			Title     string `json:"title"`
			GlossList struct {
				GlossEntry struct {
					ID        string `json:"ID"`
					SortAs    string `json:"SortAs"`
					GlossTerm string `json:"GlossTerm"`
					Acronym   string `json:"Acronym"`
					Abbrev    string `json:"Abbrev"`
					GlossDef  struct {
						Para         string   `json:"para"`
						GlossSeeAlso []string `json:"GlossSeeAlso"`
					} `json:"GlossDef"`
					GlossSee string `json:"GlossSee"`
				} `json:"GlossEntry"`
			} `json:"GlossList"`
		} `json:"GlossDiv"`
	} `json:"glossary"`
}
