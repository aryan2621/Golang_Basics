package main

import "fmt"

func main() {
	// "%d":Integer
	// "%f":Float
	// "%t":Boolean
	// "%s":String
	// "%c":Character
	// "%v":Value
	// "%T":Type
	// "%b":Binary
	// "%x":Hexadecimal
	// "%o":Octal
	// "%U":Unicode
	// "%e":Scientific Notation
	// "%E":Scientific Notation
	// "%p":Pointer
	// "%q":Quoted String
	// "%b":Base 2
	// "\n"	:New Line
	// "\t"	:Tab
	// "\r"	:Carriage Return

	fmt.Printf("Hello, World!")
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.13333334)
	fmt.Printf("%9.f\n", 3.1333333333333334)

	sp1 := fmt.Sprintf("%9.f", 3.141333333333333333593)
	fmt.Println(sp1)
}
