package main

import (
	"fmt"
	"regexp"
)

func main() {
	reStr := "The sun was shining brightly on the small town as the birds sang their morning songs. The streets were bustling with activity as people went about their daily routines. Children laughed and played while adults went to work or ran errands. The local cafe was serving up hot coffee and fresh pastries, the aroma wafting through the air. The park was filled with joggers, dog-walkers, and families enjoying a day out together. The community garden was a vibrant green, with rows of colorful flowers and vegetables. Everyone seemed to be in good spirits, basking in the warmth and beauty of the day. It was a perfect day to be alive and enjoy all that life has to offer."
	match, _ := regexp.MatchString("en[^ ]?", reStr)
	fmt.Println(match)

	r, _ := regexp.Compile("[a-z]er")
	fmt.Println(r.MatchString(reStr))
	// return true or false

	fmt.Println(r.FindString(reStr))
	// return the first match

	fmt.Println(r.FindStringIndex(reStr))
	// return the index of the first match

	fmt.Println(r.FindStringSubmatch(reStr))
	// return the first match and the submatch

	fmt.Println(r.FindStringSubmatchIndex(reStr))
	// return the index of the first match and the submatch

	fmt.Println(r.FindAllString(reStr, -1))
	// return all match

	fmt.Println(r.FindAllStringSubmatchIndex(reStr, -1))
	// return all match and the submatch

	fmt.Println(r.FindAllString(reStr, 2))
	// return the first 2 match

	fmt.Println(r.ReplaceAllString(reStr, "Goog"))
}
