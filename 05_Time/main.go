package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study for Golang")
	presentTime := time.Now()

	fmt.Println("The present time is", presentTime)
	fmt.Println("The present modified time is",
		presentTime.Format("01/02/2006 15:04:05 Monday"))
	createdDate := time.Date(2004, time.February, 10, 10,
		10, 10, 10, time.UTC)

	fmt.Println(`The created date is`,
		createdDate.Format("01/02/2006 15:04:05 Monday"))
}
