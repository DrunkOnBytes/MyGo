package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to MySlices!")

	var fruits = []string{"apple", "peach"}
	fmt.Printf("Type of fruits is %T\n", fruits)

	fruits = append(fruits, "Mango", "banana")
	fmt.Println(fruits)

	fruits = append(fruits[1:])
	fmt.Println(fruits)

	fruits = fruits[:2]
	fmt.Println(fruits)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 657
	highScores[2] = 105
	highScores[3] = 239
	// highScores[4] = 345
	highScores = append(highScores, 945, 555)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//remove vslue from slice based on index
	courses := []string{"flutter", "c++", "swift", "python"}
	fmt.Println(courses)
	courses = append(courses[:2], courses[2+1:]...)
	fmt.Println(courses)
}
