package main

import (
	"fmt"
	"log"
	"mongo/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB CRUD Operations")

	r := router.Router()
	
	fmt.Println("Starting the application...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Application started successfully")
}
