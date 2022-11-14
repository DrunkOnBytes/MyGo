package main

import "fmt"

func main() {
	fmt.Println("Welcome to MyFunctions")
	greet()

	result := adder(3, 4)
	fmt.Println("Result is:", result)
	res, nm := proAdder(1, 2, 3, 4, 5)
	fmt.Println("ProResult is:", res, nm)
}

func greet() {
	fmt.Println("Namaste Golang")
}

func adder(v1 int, v2 int) int { //function signature
	return v1 + v2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hero"
}
