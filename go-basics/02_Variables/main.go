package main

import "fmt"

var jwtToken = 7000

// jwtToken:=7000
// will not run in global scope

const LoginToken int = 1000

// capital first letter means the variable is public

func main() {
	var username string = "John Doe"
	fmt.Println(username)
	fmt.Printf("Variables is of Type :%T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variables is of Type :%T \n", isLoggedin)

	var smallValue int = 255
	fmt.Println(smallValue)
	fmt.Printf("Variables is of Type :%T \n", smallValue)

	var smallFloat float64 = 255.255782748172041
	fmt.Println(smallFloat)
	fmt.Printf("Variables is of Type :%T \n", smallFloat)

	// default value and some alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variables is of Type :%T \n", anotherVariable)

	//implicit Type
	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Variables is of Type :%T \n", website)
	// website=3

	numberOfUsers := 100
	fmt.Println(numberOfUsers)
	fmt.Printf("Variables is of Type :%T \n", numberOfUsers)

	fmt.Println(jwtToken)
	fmt.Println(LoginToken)
}
