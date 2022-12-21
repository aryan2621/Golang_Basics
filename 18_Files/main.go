package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to class of Files!")

	content := "This is a sample content using as an example for files in go"
	file, err := os.Create("./sample.txt")
	// this will create a file if it does not exist

	check(err)
	fmt.Println("Step 1 is done!")

	len, err := io.WriteString(file, content)
	// this will write the content to the file

	check(err)
	fmt.Println("Step 2 is done! as len", len)

	file.Close()
	readFile("./sample.txt")
}

func readFile(filename string) {
	dataType, err := ioutil.ReadFile(filename)
	// _ as reading is not done as string
	// it happens as reading bytes

	check(err)

	fmt.Println("DataType is:", dataType)
	fmt.Println("Readable DataType is:", string(dataType))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
