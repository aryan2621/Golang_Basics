package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type TodoList struct {
	TodoCount int
	Todos     []string
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func getStrings(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil
	}
	errorCheck(err)
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	errorCheck(scan.Err())
	return lines
}

func write(w http.ResponseWriter, message string) {
	_, err := w.Write([]byte(message))
	errorCheck(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	write(w, "Hello Internet!")
}
func englishHandler(w http.ResponseWriter, r *http.Request) {
	write(w, "Hello World!")
}

func interactHandler(w http.ResponseWriter, r *http.Request) {
	todoVals := getStrings("todos.txt")
	tmpl, err := template.ParseFiles("index.html")
	errorCheck(err)
	todos := TodoList{
		TodoCount: len(todoVals),
		Todos:     todoVals,
	}
	err = tmpl.Execute(w, todos)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("new.html")
	errorCheck(err)
	err = tmpl.Execute(w, nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	todo := r.FormValue("todos")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("todos.txt", options, os.FileMode(0600))
	errorCheck(err)
	_, err = fmt.Fprintln(file, todo)
	errorCheck(err)
	errorCheck(file.Close())
	http.Redirect(w, r, "/interact", http.StatusFound)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/interact", interactHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)
	http.ListenAndServe(":8080", nil)

}
