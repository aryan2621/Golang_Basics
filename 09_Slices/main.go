package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to class of Slices!")

	var fruitList = []string{"Apple", "Orange", "Banana", "Grapes"}
	fmt.Println("FruitList is:", fruitList)
	fmt.Printf("Type of fruitList is: %T\n", fruitList)

	fruitList = append(fruitList, "Pineapple", "Mango")
	fmt.Println("FruitList is:", fruitList)

	// fruitList = append(fruitList[:3])
	// include 1, exclude 3

	fmt.Println("FruitList is:", fruitList)

	highScore := make([]int, 5)
	// make means initialize
	highScore[0] = 100
	highScore[1] = 200
	highScore[2] = 300
	highScore[3] = 400
	highScore[4] = 500

	fmt.Println("HighScore is:", highScore)
	fmt.Println("Len of HighScore is:", len(highScore))

	highScore = append(highScore, 600, 900, 1200, 1280)
	// by append we can add more values

	fmt.Println("HighScore is:", highScore)
	fmt.Println("Len of HighScore is:", len(highScore))

	sort.Ints(highScore)
	fmt.Println("Sorted HighScore is:", highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove an element from slice based on index

	var courses = []string{"Docker", "Kubernetes",
		"Puppet", "Terraform", "AWS", "Azure", "GCP"}

	fmt.Println("Courses are:", courses)

	var index int = 3
	courses = append(courses[:index], courses[index+1:]...)

	// :index means from 0 to index-1 and index+1 means from
	// index+1 to end and
	//  ... means add all the values

	fmt.Println("Courses are:", courses)

}
