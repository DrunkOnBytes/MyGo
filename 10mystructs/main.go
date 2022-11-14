package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs!")
	//there is no inheritance in GOLANG; So no super or parent

	arjun := User{"Arjun", "asinha@commvault.com", true, 21}
	fmt.Println(arjun)
	fmt.Printf("More details of arjun: %+v\n", arjun)
	fmt.Println("Age of Arjun is", arjun.Age)
}

//User struct
type User struct{
	Name string
	Email string
	Status bool
	Age int
}