package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to class of Maps!")
	language := make(map[string]string)
	// map[key]value

	language["js"] = "JavaScript"
	language["py"] = "Python"
	language["rb"] = "Ruby"
	language["go"] = "GoLang"
	language["cpp"] = "C++"

	fmt.Println("Language is:", language)
	fmt.Println("Length of Language is:", len(language))

	fmt.Println("Value at key js is:", language["js"])

	// loop in map
	for key, value := range language {
		fmt.Println("Key is:", key, ",Value is:", value)
		// if you want to use comma ok use _, value := range language
		// _ means ignore the value
	}
	// delete a key value pair
	delete(language, "rb")
}
