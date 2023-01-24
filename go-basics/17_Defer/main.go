package main

import "fmt"

func main() {
	// defer is used to execute a function 
	// after the main function is executed

	// it is used to close the file, 
	// database connection, etc

	defer fmt.Println("This is the first defer statement")
	defer fmt.Println("This is the second defer statement")
	defer fmt.Println("This is the third defer statement")

	// in the above example, the statements 
	// will be executed in the reverse order
	// last in first out

	fmt.Println("Welcome to class of Defer!")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i + 1)
	}
}
