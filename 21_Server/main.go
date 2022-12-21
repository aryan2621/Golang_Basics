//  GET request in Golang

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to class of GET request in Golang!")

	//get request to the server
	performGetRequest()
}

func performGetRequest() {
	const URL = "https://localhost:8000/get"

	response, err := http.Get(URL)

	if err != nil {
		fmt.Println("Error is:", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("Response is:", response)
	fmt.Println("Response Status is:", response.Status)
	fmt.Println("Response Headers are:", response.Header)

	// read the content of the response
	content, _ := ioutil.ReadAll(response.Body)
	// content is of type []byte

	// Method -1 to convert []byte to string
	fmt.Println("Content is:", string(content))

	// Method -2 to convert []byte to string
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is:", byteCount)
	fmt.Println("Response string is:", responseString.String())
}
