package example

import (
	"strconv"
)

var Name string = "Rishabh"

func IntArrToStrArr(arr []int) []string {
	var strArr []string
	for _, i := range arr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}

func Into2(arr []int) []int{
	var reponseArr []int
	for _, i := range arr {
		reponseArr = append(reponseArr, i*2)
	}
	return reponseArr
}