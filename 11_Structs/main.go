package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of Structs!")
	// no inheritance in go
	// no method overloading in go
	// no operator overloading in go
	// no default constructor in go
	// no destructor in go
	// no friend function in go
	// no static keyword in go
	Rishabh := User{"Rishabh", 22, "Active", "xyx@gmail.com"}
	fmt.Println("User is:", Rishabh)
	// {Rishabh 22 Active xyx@gmail.com}

	fmt.Printf("User details are: %+v\n", Rishabh)
	// %+v is used to print the struct in a better way
	// {Name:Rishabh Age:22 Status:Active Email:xyx@gmail.com}

	fmt.Printf("Name is %v and email is %v", Rishabh.Name, Rishabh.Email)
}

type User struct {
	Name   string
	Age    int
	Status string
	Email  string
}
