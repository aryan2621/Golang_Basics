package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to class of MOD in Golang!")

	greeter()
	router := mux.NewRouter()
	// above line is used to create a new router

	router.HandleFunc("/", serveHome).Methods("GET")
	// above line is used to handle the request to the root path of the website

	log.Fatal(http.ListenAndServe(":8080", router))
	// above line is used to listen to the port 8080 and serve the request
}
func greeter() {
	fmt.Println("Hello, Gopher!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {

	// responseWriter is used to write the response back to the client
	// http.Request is a struct that contains information about the request
	// w.write is used to write the response back to the client as a byte slice (byte array)
	// contains the actual HTML that we want to send back to the client

	w.Write([]byte(`
	<h1>Hello, Gopher!</h1>
	<p>Welcome to my website!</p>
	<p>Go is a great language!</p>
	`))
}
