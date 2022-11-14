package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to ControlFlow!")

	//IF - ELSE
	loginCount := 10
	var result string
	if loginCount < 10 {
		result = "Less"
	} else if loginCount > 10 {
		result = "More"
	} else {
		result = "Exactly 10"
	}
	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

	fmt.Println("\n")

	//SWITCH - CASE (LUDO)
	rand.Seed(time.Now().UnixNano())
	diceNo := rand.Intn(6) + 1
	fmt.Println("Value of dixe is", diceNo)

	//switch works with strings too in Golang
	switch diceNo {
	case 1:
		fmt.Println("Dice value is 1, OPEN UP")
	case 2:
		fmt.Println("MOVE UP 2")
	case 3:
		fmt.Println("MOVE UP 3")
	case 4:
		fmt.Println("MOVE UP 4")
	case 5:
		fmt.Println("MOVE UP 5")
		fallthrough //needs to be explicity mentioned to fallthrough, but doesn't go to "default"
	case 6:
		fmt.Println("MOVE UP 6 and RE-ROLL")
	default:
		fmt.Println("what was that?!")
	}
}
