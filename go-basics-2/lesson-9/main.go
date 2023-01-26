package main

import (
	"fmt"
)

var pl = fmt.Println

type Animal interface {
	Angry()
	Happy()
}
type Cat string

func (c Cat) Angry() {
	pl("Cat is angry")
}

func (c Cat) Happy() {
	pl("Cat is happy")
}

func (c Cat) Attack() {
	pl("Cat is attacking")
}

func (c Cat) Name() string {
	return string(c)
}

func PrintTo15() {
	for i := 0; i < 15; i++ {
		pl("PrintTo15", i)
	}
}

func PrintTo20() {
	for i := 0; i < 20; i++ {
		pl("PrintTo20", i)
	}
}

func Nums1(channel chan int) {
	for i := 0; i < 4; i++ {
		channel <- i
	}
	close(channel)
}
func Nums2(channel chan int) {
	for i := 3; i < 7; i++ {
		channel <- i
	}
	close(channel)
}
func useFunc(f func(int, int) int, x int, y int) {
	// f func(int int) int is a function that takes 2 int and returns an int
	// x int is an int y int is an int too (x and y are parameters)
	pl(f(1, 2))
}

func sumValues(a int, b int) int {
	return a + b
}
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var animal Animal = Cat("Tom")
	animal.Angry()
	animal.Happy()
	// animal.Attack() // error
	// animal.Name() // error

	// Threads and Routine
	// go PrintTo15()
	// go PrintTo20()
	// Nothing will be printed
	// because the main thread is finished
	// before the other threads

	// To fix this, we need to wait for the threads to finish
	// go PrintTo15()
	// go PrintTo20()
	// time.Sleep(2 * time.Second)

	// Channels
	channel := make(chan int)
	channel2 := make(chan int)
	go Nums1(channel)
	go Nums2(channel2)
	pl("Channel 1", <-channel)
	pl("Channel 1", <-channel)
	pl("Channel 1", <-channel)
	pl("Channel 2", <-channel2)
	pl("Channel 2", <-channel2)
	pl("Channel 2", <-channel2)

	// Channel is used to communicate between threads and routines
	// Channel is a blocking queue (FIFO)
	// Channel is a type of data structure that can be used to communicate between threads and routines
	// They helps to synchronize the threads and routines

	// Clousures
	// A clousure is a function that can access the variables of the outer function
	intSum := func(a int, b int) int {
		return a + b
	}
	pl(intSum(1, 2))

	samp1 := 1
	pl("Before :", samp1)
	changeSamp1 := func() {
		samp1++
	}
	changeSamp1()
	pl("After :", samp1)

	useFunc(sumValues, 1, 2)

	pl(factorial(10))
}
