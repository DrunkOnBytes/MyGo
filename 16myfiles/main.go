package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to MyFiles!")
	content := "This needs to go in a file - ArjunSinha"

	file, err := os.Create("./arjunsinha.txt")
	checkNilErr(err)
	defer file.Close()

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)

	readFile("./arjunsinha.txt")
}

func readFile(filePath string) {
	dataByte, err := ioutil.ReadFile(filePath)
	checkNilErr(err)
	fmt.Println("Text data in File is:")
	fmt.Println(string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
