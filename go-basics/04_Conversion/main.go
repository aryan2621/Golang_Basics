package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza Shop")
	fmt.Println("Rate our Pizza from 1 to 5:")

	reader := bufio.NewReader(os.Stdin)

	response, _ := reader.ReadString('\n')
	fmt.Println("The response is", response)
	responseTemp := strings.TrimSpace(response)

	numRating, err := strconv.ParseFloat(responseTemp, 64)

	if err != nil {
		fmt.Println("The error is", err)
	} else {
		fmt.Println("The rating is", numRating+1)
	}
}
