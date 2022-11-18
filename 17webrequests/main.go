package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to WebRequests!")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//always defer the close as soon as you get response
	defer response.Body.Close()

	fmt.Printf("Type of response: %T\n", response)

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("Content of response", content)
}
