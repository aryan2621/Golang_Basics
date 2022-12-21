package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to class of Switch Case!")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	// between 0 and 5

	fmt.Println("Dice Number is:", diceNumber)

	switch diceNumber {
	// fallthrough is used to skip the break statement
	// it continues to the next case
	case 1:
		fmt.Println("You got 1 and now you can open")
	case 2:
		fmt.Println("You got 2 and now you can move 2 steps")
	case 3:
		fmt.Println("You got 3 and now you can move 3 steps")
	case 4:
		fmt.Println("You got 4 and now you can move 4 steps")
		fallthrough
		// You got 4 and now you can move 4 steps
		// You got 5 and now you can move 5 steps
	case 5:
		fmt.Println("You got 5 and now you can move 5 steps")
	case 6:
		fmt.Println("You got 6 and now you can move 6 steps and roll the dice again")
	default:
		fmt.Println("This is not expected from a dice")
	}
}
