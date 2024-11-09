package main

import (
	"fmt"
)

func main() {
	var a bool = true    //Boolean (true/false)
	var b int = 5        //Integer
	var c float32 = 3.14 //Floating
	var d string = "Hi!" //String
	e := `String
				multilinha`

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float: ", c)
	fmt.Println("String: ", d)
	fmt.Println("String: ", e)
}
