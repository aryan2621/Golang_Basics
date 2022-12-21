// POST request in Golang

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to class of POST request in Golang!")

	// post request to the server
	performJsonPostRequest()
}

func performJsonPostRequest() {
	const URL = "https://localhost:8000/post"

	// fake jsonPayload
	requestBody := strings.NewReader(`{"name":"John", "age":30, "city":"New York"}`)

	response, err := http.Post(URL, "application/json", requestBody)

	if err != nil {
		fmt.Println("Error is:", err)
		return
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content is:", string(content))
}
