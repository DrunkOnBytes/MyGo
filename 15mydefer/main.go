package main

import "fmt"

func main() {
	fmt.Println("Welcome to MyDefer !")

	//"defer" works in LIFO form
	//"defer" code lines execute at the end of the current function
	defer fmt.Println("One")
	defer fmt.Println("Two")
	deferFunc()
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func deferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
