package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mayPanic()
	fmt.Println("After mayPanic()")
}
