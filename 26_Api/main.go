// building api in golang-modals

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses -file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}

// fake database
var courses []Course

// middleware,helper -file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == "" && c.CoursePrice == 0 && c.Author == nil
	return c.CourseName == "" && c.CoursePrice == 0 && c.Author == nil
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// controller -file
// serve home page /route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<h1> Welcome to my API </h1>
		<p> Please use /courses to get the courses </p>
		<a> href="/courses"> /courses </a>
	`))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
	// above line is used to encode the courses to json and send it to the response writer
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// get the id from the url
	params := mux.Vars(r)
	// this line will get the id from the url and store it in the params variable
	// mux.Vars means that it will get the variables from the url

	// loop through the courses and find the one with the id from the url
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No such course found")
}
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Pleasen find some data")
		return
	}

	// what about {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	// this line will decode the json data from
	// the request body and store it in the course variable

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No course data found")
		return
	}
	// generate unique id and convert it to string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000000))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	// this line will encode the course variable to
	//  json and send it to the response writer

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// get the id from the url
	params := mux.Vars(r)

	// loop through the courses and find the one with the id from the url
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// this line will remove the course from the courses slice

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			// this line will decode the json data from the request body and store
			// it in the course variable

			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No such course found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// get the id from the url
	params := mux.Vars(r)

	// loop through the courses and find the one with the id from the url
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// this line will remove the course from the courses slice

			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No such course found")
}

func main() {
	fmt.Println("Welcome to my API")
	router := mux.NewRouter()

	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "Go",
		CoursePrice: 100,
		Author: &Author{
			FullName: "John Doe",
			Website:  "johndoe.com",
		},
	})
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Python",
		CoursePrice: 200,
		Author: &Author{
			FullName: "Jane Doe",
			Website:  "janedoe.com",
		},
	})

	// handle the routes
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")


	// listen to the routes
	log.Fatal(http.ListenAndServe(":8000", router))
}
