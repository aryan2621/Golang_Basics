package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of Arrays!")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[2] = "Banana"
	fruitList[3] = "Grapes"

	fmt.Println("FruitList is:", fruitList)
	fmt.Println("Len of FruitList is:", len(fruitList))

	var vegList = [100]string{"Potato", "Tomato", "Onion","Carrot"}
	fmt.Println("VegList is:", vegList)
	fmt.Println("Len of VegList is:", len(vegList))

}
