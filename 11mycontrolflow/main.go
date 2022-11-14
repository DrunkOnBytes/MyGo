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
	if loginCount<10{
		result = "Less"
	} else if loginCount>10{
		result = "More"
	} else {
		result = "Exactly 10"
	}
	fmt.Println(result)

	if num := 3; num<10 {
		fmt.Println("Num less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}
	
	fmt.Println("\n\n\n")
	
	//SWITCH - CASE
	rand.Seed(time.Now().UnixNano())
	diceNo := rand.Intn(6)+1
	fmt.Println("Value of dixe is", diceNo)
}