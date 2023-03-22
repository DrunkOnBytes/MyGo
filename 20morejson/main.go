package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Topic    string   `json:"topic"`
	password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to more JSONs")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	someCourses := []course{
		{" Learn React", 399, "ReactJS", "abc123", []string{"web-dev", "js"}},
		{" Learn Flutter", 599, "Flutter", "abc123", []string{"mob-dev", "Dart", "Flutter"}},
		{" Learn Angular", 299, "AngularJS", "abc123", nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(someCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Learn React",
		"price": 399,
		"topic": "ReactJS",
		"tags": [
				"web-dev",
				"js"
		]
	}
	`)

	var courses course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Println("JSON not valid")
	}

	//just access data from JSON as key-value pairs
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
