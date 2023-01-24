package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	fmt.Println("Hello World")
	pl("Hello World")

	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered

	name, err := reader.ReadString('\n')

	if err == nil {
		pl("Hello", name)
	} else {
		pl(err)
	}

	// variables
	// var name type
	// capital letter for public means it can be accessed outside the package

	// Method-1
	var vName string = "Hello World"

	// Method-2
	var v1, v2 = 1.2, 3.4
	var v3 = "helloworld"

	// Method-3
	v4 := 2.4
	v4 = 7.7

	// data types
	// bool
	// string
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64

	pl(vName, v1, v2, v3, v4)

	pl(reflect.TypeOf(25))
	fmt.Printf("%T\n", 25)

	// casting in go
	cv1 := 1.5
	cv2 := int(cv1)

	cv3 := "50000"
	cv4, _ := strconv.Atoi(cv3)
	pl(cv2, cv4)

	cv6 := 889
	cv7 := strconv.Itoa(cv6)
	pl(cv7)

	cv8 := "3.14"
	cv9, _ := strconv.ParseFloat(cv8, 64)
	pl(cv9)

	// "%f" means format the following value as a float
	// "%d" means format the following value as an integer
	// "%s" means format the following value as a string
	// "%t" means format the following value as a boolean
	// "%v" means format the following value as a default format
	// "%T" means format the following value as a type

	// Conditional Statements : > < >= <= == !=
	// && || !

	pl("Enter your age")
	var userAge int
	fmt.Scan(&userAge)
	pl(userAge)
	if userAge > 18 {
		pl("You are eligible to vote")
	} else if userAge == 18 {
		pl("You are not eligible to vote and at the balance age")
	} else {
		pl("You are not eligible to vote")
	}

	sv1 := "A word"
	replace := strings.NewReplacer("A", "The")
	sv2 := replace.Replace(sv1)
	pl(sv2, len(sv2))
	pl("Contains another word", strings.Contains(sv1, "another"))
	pl("o index", strings.Index(sv1, "o"))
	pl("Replace word", strings.Replace(sv1, "word", "sentence", 1))
	// -1 means replace all

	s3 := "\nsomeWords\n"
	sv4 := strings.TrimSpace(s3)
	pl(sv4, "a-b-c-d-e-f-g-h-i-j-k-l-m-n-o-p-q-r-s-t-u-v-w-x-y-z")
	pl("Lowercase", strings.ToLower(sv4))
	pl("Uppercase", strings.ToUpper(sv4))
	pl("Prefix", strings.HasPrefix(sv4, "some"))
	pl("Suffix", strings.HasSuffix(sv4, "some"))
	pl("Split", strings.Split(sv4, " "))
	pl("Join", strings.Join([]string{"a", "b", "c"}, "-"))
}
