package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}
func getSum(a int, b int) int {
	return a + b
}
func getTwoSum(a int, b int) (int, int, int) {
	return a + 1, a + 2, b + a
}
func main() {
	sayHello()
	fmt.Print(getSum(100, 116))
	fmt.Println()
	fmt.Println(getTwoSum(100, 116))
}
