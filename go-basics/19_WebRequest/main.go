package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://leetcode.com/aryan2621"

func main() {
	fmt.Println("Welcome to class of WebRequest!")

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of Type: %T\n", response)
	// Response is of Type: *http.Response which is a pointer to http.Response

	defer response.Body.Close()
	// this is to close the connection

	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data from the %v is: %v", URL, string(databyte))
}
