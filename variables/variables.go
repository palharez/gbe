package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables
	var h, g = "test", false
	fmt.Println(h, g)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// := is shorthand for declaring and initializing a variable
	// This thype of syntax can only be used inside a function
	f := "apple" 
	fmt.Println(f)
}