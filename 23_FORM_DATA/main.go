// sending the data collected from the form to the server

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welcome to class of POSTFORM request in Golang!")
	performFormPostRequest()
}

func performFormPostRequest() {
	const URL = "https://localhost:8000/postform"

	data := url.Values{}
	data.Add("name", "John")
	data.Add("age", "30")
	data.Add("city", "New York")
	data.Add("country", "USA")

	response, err := http.PostForm(URL, data)
	if err != nil {
		fmt.Println("Error is:", err)
		return
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error is:", err)
		return
	}
	fmt.Println("Content is:", string(content))
}
