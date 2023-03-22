package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to webserver demo")
	// PerformGetRequest("http://localhost:8000/get")

	// //fake json payload
	// requestBody := strings.NewReader(`
	// 	{
	// 		"name":"arjun",
	// 		"age":22
	// 	}
	// `)
	// PerformPostJsonRequest("http://localhost:8000/post", requestBody)

	//fake form payload
	data := url.Values{}
	data.Add("firstName", "arjun")
	data.Add("lastName", "sinha")
	data.Add("userName", "DrunkOnBytes")

	PerformPostFormRequest("http://localhost:8000/postform", data)
}

func PerformGetRequest(myUrl string) {
	res, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Status Code: ", res.StatusCode)
	fmt.Println("Content length: ", res.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(content))
	byteCount, _ := responseString.Write(content)
	fmt.Println("Bytecount is: ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest(myUrl string, requestBody io.Reader) {

	res, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest(myUrl string, data url.Values) {
	res, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
