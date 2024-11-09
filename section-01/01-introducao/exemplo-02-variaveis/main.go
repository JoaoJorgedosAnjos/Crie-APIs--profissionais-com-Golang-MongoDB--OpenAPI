package main

import (
	"fmt"
)

var value = 1 // Global
var v2 int    // Global

func main() {
	// Local
	var code int
	test := true
	var name string
	other := "John Doe"
	fmt.Println(value)
	fmt.Println(v2)
	fmt.Println(code)
	fmt.Println(test)
	fmt.Println(name)
	fmt.Println(other)
}
