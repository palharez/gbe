package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func sum_and_product(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	sum, product := sum_and_product(a, b)
	fmt.Println(sum)
	fmt.Println(product)
}