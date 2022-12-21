package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of Jumps Statements!")

	days := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday",
		"Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// For loops
	// Type 1: for init; condition; post {}
	// for i := 0; i < len(days); i++ {
	// 	fmt.Printf("Day Number is %v and day is %v\n", i+1, days[i])
	// }

	// Type 2
	// for i := range days {
	// 	fmt.Printf("Day Number is %v and day is %v\n", i+1, days[i])
	// }

	// Type 3
	// for index, day := range days {
	// 	fmt.Printf("Day Number: %v and day: %v\n", index+1, day)
	// 	//  index means the index of the array and day means
	// 	// the value of the array
	// 	// _ is used to ignore the index
	// }

	// Type 4
	// Kind of while loop
	iter := 0
	for iter < 10 {
		iter++
		if iter == 8 {
			goto IdharJumpKaro
		}
		// if iter == 5 {
		// 	break
		// 	// since the condition is true, it will break the loop
		// }
		if iter == 5 {
			continue
			// continue will skip the value 5 and continue the loop
		}
		fmt.Printf("Iter value: %v, ", iter)

	}

	// Go to statement
	// goto <label>
IdharJumpKaro:
	fmt.Println("I am jumping from here")

}
