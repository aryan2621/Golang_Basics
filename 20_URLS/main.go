package main

import (
	"fmt"
	"net/url"
)

const URL = `https://lco.dev:3000/learn?course=web-dev&batch=2021&teacher=prince&paymentid=1234567890`

func main() {
	fmt.Println("Welcome to class of URLS!")

	fmt.Println(URL)

	// parse the URL
	result, _ := url.Parse(URL)
	fmt.Printf("Result is of Type: %T and value is: %v\n", result, result)
	fmt.Println("Scheme is:", result.Scheme, result.Host,
		result.Path, result.RawQuery, result.Port())

	// get the query parameters
	qparams := result.Query()
	fmt.Println("Query Parameters are:", qparams)
	// in the above line, qparams is of type map[string][]string

	for _, v := range qparams {
		fmt.Println(v)
	}

	partsOfUrls := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "course=web-dev&batch=2021&teacher=prince&paymentid=1234567890",
	}
	fmt.Println(partsOfUrls)
	fmt.Println(partsOfUrls.String())
}
