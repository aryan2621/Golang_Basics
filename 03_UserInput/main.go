package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Welcome to GoLang"
	fmt.Println(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	// comma ok || err ok
	input, err := reader.ReadString('\n')
	fmt.Println("The Name is, ", input)
	fmt.Printf("The TypeName is %T", input)

	fmt.Println("The error  is", err)
}
