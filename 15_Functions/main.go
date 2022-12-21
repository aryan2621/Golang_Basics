package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions in GoLang")

	// main function is the entry point of the program
	// not allowed to write function inside a function
	// function can be called from anywhere in the program

	reader()

	result := adder(3, 5)
	fmt.Println("Result:", result)

	result2, responseString := adderModified(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("Result2: %v and string: %v\n", result2, responseString)
}

func adder(a int, b int) int {
	// func function_name (parameter1 type,
	// parameter2 type) return_type
	return a + b
}

func adderModified(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "The result is here"
}

func reader() {
	fmt.Println("Namastey from Rishabh Verma")
}
