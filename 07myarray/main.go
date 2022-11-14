package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays!")

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Pine"
	fruits[3] = "Banana"

	fmt.Println(fruits, len(fruits))

	var vegies = [7]string{"Potato", "Tomato", "Onions"}
	fmt.Println(vegies, len(vegies))
}
