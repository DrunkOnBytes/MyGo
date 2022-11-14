package main

import "fmt"

func main() {
	fmt.Println("Welcome to MyPointers")

	var ptr *int
	fmt.Println("Value of ptr = ", ptr)

	myNumber := 25
	var p1 = &myNumber
	fmt.Println("Value of actual pointer p1 = ", p1)
	fmt.Println("Value of actual pointer p1 = ", *p1)

	*p1 = *p1 * 2
	fmt.Println("Value of actual pointer p1 = ", myNumber)
}
