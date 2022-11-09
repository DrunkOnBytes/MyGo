package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input:"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our APP:")

	//comma ok || comma error syntax 
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ",input)
	fmt.Printf("Type of rating: %T",input)

}