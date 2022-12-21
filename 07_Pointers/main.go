package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of Pointers!")

	var ptrForInteger *int
	fmt.Println("ptrForInteger:", ptrForInteger)

	myNumber := 10
	ptrForInteger = &myNumber

	// & for address
	// * for value
	

	fmt.Println("ptrForMyNumber:", ptrForInteger)
	fmt.Println("ptrForInteger Value:", *ptrForInteger)
	fmt.Println("myNumber:", myNumber)

	*ptrForInteger = *ptrForInteger + 10

	fmt.Println("Changed ptrForInteger Value:", *ptrForInteger)
	fmt.Println("Changed myNumber:", myNumber)

}
// Golang
// Blockchain
// C#
// Slevete
// Serverless