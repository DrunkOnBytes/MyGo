package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of GO")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("15:04:05 02-01-2006 Monday"))

	createdDate := time.Date(2000, time.December, 2, 23, 23, 23, 0, time.UTC)
	fmt.Println(createdDate.Format("15:04:05 02-01-2006 Monday"))
}
