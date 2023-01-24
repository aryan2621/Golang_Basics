package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of Structs!")
	// no destructor in go
	// no friend function in go
	// no static keyword in go

	Rishabh := User{"Rishabh", 22, "Active", "xyx@gmail.com", 29}
	fmt.Println("User is:", Rishabh)
	// {Rishabh 22 Active xyx@gmail.com}

	fmt.Printf("User details are: %+v\n", Rishabh)
	// %+v is used to print the struct in a better way
	// {Name:Rishabh Age:22 Status:Active Email:xyx@gmail.com}

	fmt.Printf("Name is %v and email is %v\n", Rishabh.Name, Rishabh.Email)

	Rishabh.getStatus()
	Rishabh.NewMail()

	fmt.Printf("User details are: %+v\n", Rishabh)
	// will not change the previous value of email
	// pass by value
}

type User struct {
	Name   string
	Age    int
	Status string
	Email  string
	oneAge int
}

func (user User) getStatus() {
	// func (receiverName ReceiverType) methodName() returnType
	fmt.Println("User Status is:", user.Status)
}

// Getters and Setters in GoLang

func (user User) NewMail() {
	user.Email = "rishabh@gmail.com"
	fmt.Println("New email is:", user.Email)
}
