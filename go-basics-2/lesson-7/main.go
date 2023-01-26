package main

// Importing the package by first module name and then package name
// in this case moduleName=lesson-7 and packageName=example

import (
	"fmt"
	example "lesson-7/example"
)

type MyConstraint interface {
	int | float64
}

func getSum[T MyConstraint](a T, b T) (T, T) {
	return a + b, a - b
}

type Customer struct {
	name    string
	address string
	bal     int
	age     int
}

func getCustomerInfo(c Customer) {
	fmt.Printf("Name is %s\n and address is %s\n and balance is %d\n and age is %d\n", c.name, c.address, c.bal, c.age)
}
func newCustomerAdd(c *Customer, address string) {
	c.address = address
	fmt.Println(*c)
}

type Rectangle struct {
	length  int
	breadth int
}

func area(r Rectangle) int {
	return r.length * r.breadth
}
func (r Rectangle) parameter() int {
	return 2 * (r.length + r.breadth)
}

func main() {
	fmt.Println("Hello, World!", example.Name)
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(example.IntArrToStrArr(intArr))
	fmt.Println(example.Into2(intArr))

	// Map in Go
	// map[keyType]valueType

	// var heroes1 map[string]string
	heroes1 := make(map[string]string)
	// OR
	heroes2 := make(map[string]string)

	heroes1["Batman"] = "Bruce Wayne"
	heroes1["Superman"] = "Clark Kent"
	heroes1["Ben 10"] = "Ben Tennyson"
	heroes1["WonderWomen"] = "Diana Prince"

	heroes2["Thanoes"] = "Thanos"
	heroes2["Iron Man"] = "Tony Stark"
	heroes2["Captain America"] = "Steve Rogers"
	heroes2["Hulk"] = "Bruce Banner"

	fmt.Println(heroes1)
	fmt.Println(heroes2)

	superPets := map[int]string{
		1: "Krypto",
		2: "Ace",
		3: "Duke",
		4: "Pluto",
	}
	fmt.Println(superPets)

	fmt.Println(heroes1["Batman"])
	fmt.Println(heroes1["Superman"])
	fmt.Println(heroes1["Ben 10"])

	// Looping over map
	for key, value := range heroes1 {
		fmt.Println(key, value)
	}

	// Delete from map
	delete(heroes1, "Batman") // delete(mapName, key)
	fmt.Println(heroes1)

	// Check if key exists in map
	_, ok := heroes1["Batman"]
	fmt.Println(ok)

	// generics in Go
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(1.2, 2.3))

	// Structs in Go
	var customer1 Customer
	customer1.name = "John"
	customer1.address = "New York"
	customer1.bal = 1000.00
	customer1.age = 30

	// fmt.Println(customer1)
	getCustomerInfo(customer1)
	newCustomerAdd(&customer1, "New York")

	customer2 := Customer{
		name:    "John",
		address: "New York",
		bal:     1000.00,
		age:     30,
	}
	getCustomerInfo(customer2)

	rect1 := Rectangle{
		length:  10,
		breadth: 20,
	}
	fmt.Println(area(rect1))
	fmt.Println(rect1.parameter())
}
