package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

func main() {
	// fmt.Println("Increasing")
	// for x := 0; x < 10; x++ {
	// 	fmt.Println(x)
	// }
	// fmt.Println("Decreasing")
	// for x := 10; x > 0; x-- {
	// 	fmt.Println(x)
	// }

	// fmt.Println("For without condition")
	// fx := 0
	// for fx < 10 {
	// 	fmt.Println(fx)
	// 	fx++
	// }

	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	// randomNum := rand.Intn(50) + 1
	// fmt.Println("randomNum", randomNum)
	// for {
	// 	fmt.Println("Guess the number")
	// 	var guess int
	// 	fmt.Scan(&guess)
	// 	if guess > randomNum {
	// 		fmt.Println("Too high")
	// 	} else if guess < randomNum {
	// 		fmt.Println("Too low")
	// 	} else {
	// 		fmt.Println("You got it")
	// 		break
	// 	}
	// }

	aNum := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range aNum {
		// i is the index and num is the value
		// _ is used to ignore the index
		fmt.Println(num)
	}

	// default values in array for int is 0 and for
	// string is "" (empty string) and for bool is false

	var arr1 [5]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5
	fmt.Println(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	for _, num := range arr2 {
		fmt.Println(num)
	}

	arr3 := [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println(arr3)
	for _, arr := range arr3 {
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}

	aStr1 := "Hello"
	for _, r := range aStr1 {
		fmt.Printf("Rune Array: %c\n", r)
	}

	byteAtt := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	bStr := string(byteAtt)
	fmt.Println(bStr)

	// slices
	sl1 := make([]string, 3)
	fmt.Println(sl1)
	sl1[0] = "a"
	sl1[1] = "b"
	sl1[2] = "c"
	fmt.Println(sl1)

	for i := 0; i < len(sl1); i++ {
		fmt.Print(sl1[i])
	}
	for _, s := range sl1 {
		fmt.Print(s)
	}

	sArr := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Println("First three elements of sArr", sArr[:3])
	fmt.Println("Last three elements of sArr", sArr[2:])
	fmt.Println("All elements of sArr", sArr[:])
	fmt.Println("from 2nd to 4th elements of sArr", sArr[1:4])
	// [1:4] means from 1 to 3 (4-1) where 1 is inclusive and 4 is exclusive
	fmt.Println("from 3rd to 7th elements of sArr", sArr[2:7])

	sl3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl3 = append(sl3, 11, 12, 13, 14, 15)
	fmt.Println(sl3)
}
