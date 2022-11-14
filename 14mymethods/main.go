package main

import "fmt"

func main() {
	fmt.Println("Welcome to MyMethods !")
	arjun := User{"Arjun", "asinha@commvault.com", true, 21}
	fmt.Printf("More details of arjun: %+v\n", arjun)
	fmt.Println("Age of Arjun is", arjun.Age)
	arjun.GetStatus()
	arjun.NewMail("iamboss@arjun.com")
	fmt.Printf("More details of arjun: %+v\n", arjun)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail(em string) {
	u.Email = em
	fmt.Println("New Email:", u.Email)
}
