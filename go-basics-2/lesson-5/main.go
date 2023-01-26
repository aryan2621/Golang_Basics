package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func sayHello() {
	fmt.Println("Hello World")
}
func getSum(a int, b int) int {
	return a + b
}
func getTwoSum(a int, b int) (int, int, int) {
	return a + 1, a + 2, b + a
}
func getQuotientAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}
func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func getSumOfArray(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func changeF3(f3 *int) {
	*f3 = *f3 + 10
}
func dblArrayValuesByPointer(arr *[10]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
}
func getAverage(nums ...int) float64 {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return float64(sum) / float64(len(nums))
}
func main() {
	sayHello()
	fmt.Print(getSum(100, 116))
	fmt.Println()
	fmt.Println(getTwoSum(100, 116))
	fmt.Println(getQuotientAndRemainder(116, 6))
	fmt.Println(getSum2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(getSumOfArray(array))

	// Pointer and pass by reference
	f3 := 100

	var f4Ptr *int = &f3
	pl("f4Ptr:", f4Ptr)
	pl("f4Ptr value:", *f4Ptr)

	*f4Ptr = 200

	pl("f3 before change:", f3)
	changeF3(&f3)
	pl("f3 after change:", f3)

	pArr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dblArrayValuesByPointer(&pArr)
	pl("pArr:", pArr)

	// slice
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pl("Average of slice:", getAverage(slice...))

	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	iPrime := []int{2, 3, 5, 7, 11, 13}
	var sPrimeArr []string
	for _, i := range iPrime {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, num := range sPrimeArr {
		_, e := f.WriteString(num + "\n")
		if e != nil {
			log.Fatal(e)
		}
	}
	f, er := os.Open("data.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer f.Close()
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		fmt.Println("Prime :", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
}
