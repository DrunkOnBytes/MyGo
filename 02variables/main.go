package main

import "fmt"

//private variables start with lowercase, public variables start with uppercase
const LoginToken string = "astwt44523sfge523"

func main() {
	fmt.Println(LoginToken)
	fmt.Printf("Type %T \n", LoginToken)

	var user string = "Arjun"
	fmt.Println(user)
	fmt.Printf("Type %T \n", user)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type %T \n", isLoggedIn)

	var smallVal uint8 = 255 //uint8 = byte (alias)
	fmt.Println(smallVal)
	fmt.Printf("Type %T \n", smallVal)

	var smallFloat float32 = 455.567633242312452362345
	fmt.Println(smallFloat)
	fmt.Printf("Type %T \n", smallFloat)

	var bigFloat float64 = 455.5676332423121425234524
	fmt.Println(bigFloat)
	fmt.Printf("Type %T \n", bigFloat)

	//un-initialized variables are auto initz
	var aa int 
	fmt.Println(aa)
	fmt.Printf("Type %T \n", aa)

	var web = "wwwwwwwwwwwwwweeeeeeeeee"
	fmt.Println(web)
	fmt.Printf("Type %T \n", web)

	//walarus operator only allowed inside methods
	//not for global members
	noOfUsers := 399
	fmt.Println(noOfUsers)
	fmt.Printf("Type %T \n", noOfUsers)
}
