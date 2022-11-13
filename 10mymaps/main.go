package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to MyMaps")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["TS"] = "TypeScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println("JS short for", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	//loop through map
	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}