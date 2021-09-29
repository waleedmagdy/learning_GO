// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"time"
// )

// var myClient = &http.Client{Timeout: 10 * time.Second}

// func getJson(url string, target interface{}) error {
// 	r, err := myClient.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer r.Body.Close()

// 	return json.NewDecoder(r.Body).Decode(target)
// }

// type Foo struct {
// 	Bar string
// }

// func main() {
// 	foo := map[string]string{} // or &Foo{}
// 	getJson("https://jsoneditoronline.herokuapp.com/v1/docs/4ee96dd68837476ea16b8da920f1c4f8", &foo)
// 	println(&foo)

// 	foo2 := Foo{}
// 	getJson("https://jsoneditoronline.herokuapp.com/v1/docs/4ee96dd68837476ea16b8da920f1c4f8", &foo2)
// 	log.Println(foo2.Bar)
// }

package main

import (
	"fmt"
	"net/http"
)

func getheader(url string, header string) (string, error) {
	resp, err := http.Get(url)

	// if the function can't make a get request
	if err != nil { // to check if there is an error in the get request
		return "", err
	}

	defer resp.Body.Close() // to make sure the response body is closed

	ctype := resp.Header.Get(header) // getting the header value

	if ctype == "" { // if the header not found it will print an error
		fmt.Printf("Error: can't find the Header %s\n", header)
	}

	return ctype, nil // otherwise it will return the header value
}

func main() {
	header, err := getheader("https://golangzxasdasdasdcr.org", "Content-Type")
	if err != nil { // check if there is an error so it will print it
		fmt.Printf("Error: %s\n", err)
	} else { // otherwise it will print the value of the header
		fmt.Println(header)
	}
}
