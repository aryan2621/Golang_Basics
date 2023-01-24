package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of Control Flow!")
	loginCount := -100

	// Three types of control flow

	if loginCount < 10 {
		fmt.Println("Login Count is less than 10")
	} else if loginCount > 50 && loginCount < 100 {
		fmt.Println("Login Count is greater than 50 and less than 100")
	} else {
		fmt.Println("Login Count is greater than 50")
	}

	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	if num := 1000; num < 0 {
		fmt.Println(num, "is negative")
	} else if num > 0 {
		fmt.Println(num, "is positive")
	} else {
		fmt.Println(num, "is zero")
	}
}
