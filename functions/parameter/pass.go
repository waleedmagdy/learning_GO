package main

import "fmt"

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

//func to double the pointer , value of the variable
func doubleptr(n *int) {
	*n *= 2
}

func main() {
	values := []int{1, 2, 6, 4}
	doubleAt(values, 2) // this will pass the refrence
	fmt.Println(values)

	val := 10
	double(val)      // this will pass the value of the variable
	fmt.Println(val) // it will print 10, as this will pass value

	// to change the variable value you need to pass it's position
	// using pointer
	val2 := 10
	doubleptr(&val2)
	fmt.Println(val2) // this will print 20 as you pass the pointer

}
