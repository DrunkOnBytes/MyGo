package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=sfn485eek4t"

func main() {
	fmt.Println("Welcome to MyUrls!")
	// fmt.Println(myurl)

	//parsing url
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of qparams is: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	for key, val := range qparams {
		fmt.Printf("Value for %v is %v\n", key, val)
	}
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
