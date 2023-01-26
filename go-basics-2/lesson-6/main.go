package main

import (
	"fmt"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	/*
		Eaxctly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.

		O_RDONLY - Open the file read-only.
		O_WRONLY - Open the file write-only.
		O_RDWR - Open the file read-write.

		The remaining values may be or'ed in to control behavior.

		O_APPEND - Append data to the file when writing.
		O_CREATE - Create a new file if none exists.
		O_EXCL - Used with O_CREATE, file must not exist.
		O_SYNC - Open for synchronous I/O.
		O_TRUNC - If possible, truncate file when opened.

	*/

	_, err := os.Stat("test.txt")
	if os.IsNotExist(err) {
		pl("File does not exist,creating it now")
		file, err := os.Create("test.txt")
		if err != nil {
			return
		}
		defer file.Close()
	} else {
		f, e := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if e != nil {
			pl(e)
		}
		defer f.Close()
		if _, err := f.WriteString("appended some data"); err != nil {
			pl(err)
		}
	}
	pl("os.Args", os.Args)
	args := os.Args[1:]
	pl("args", args)
	var iArg = []int{}
	for _, arg := range args {
		val, err := strconv.Atoi(arg)
		if err != nil {
			pl("Error", err)
		}
		iArg = append(iArg, val)
	}
	max := 0
	for _, val := range iArg {
		if val > max {
			max = val
		}
	}
	pl("Max", max)
}
