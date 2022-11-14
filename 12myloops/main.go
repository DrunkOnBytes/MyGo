package main

import "fmt"

func main() {
	fmt.Println("Welcome to MyLoops !")

	days := []string{"Sun", "Mon", "Tues", "Wed", "Thurs", "Fri", "Sat", "Sun"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for i, d := range days {
	// 	fmt.Printf("index is %v and value is %v\n", i, d)
	// }

	xValue := 1
	fmt.Print("Values are: ")
	for xValue < 10 {

		if xValue == 3 {
			xValue++
			continue
		}
		if xValue == 9 {
			break
		}
		fmt.Printf("%v, ", xValue)
		xValue++

		if xValue == 2 {
			goto aass
		}
		if xValue == 7 {
			goto finish
		}
	}
	fmt.Println()

aass:
	fmt.Println("Jumping to AASS")

finish:
	fmt.Println("Printing FINISH")
}
