package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string   `json:"courseName"`           // this is called struct tag
	Price     int      `json:"coursePrice"`          // this is called struct tag
	Plateform string   `json:"coursePlateform"`      // this is called struct tag
	Password  string   `json:"-"`                    // this is called struct tag
	Tags      []string `json:"courseTags,omitempty"` // this is called struct tag
}

func main() {
	fmt.Println("Welcome to class of JSON in Golang!")

	// EncodeJson()
	DecodeJson()

}

// encoding to json means converting the data into json format

func EncodeJson() {
	lcoCources := []course{
		{"Go", 100, "LCO", "1234", []string{"Beginner", "Intermediate", "Advance"}},
		{"Python", 200, "LCO", "1234", []string{"Beginner", "Intermediate", "Advance"}},
		{"Django", 300, "LCO", "1234", []string{"Beginner", "Intermediate", "Advance"}},
	}
	for _, v := range lcoCources {
		fmt.Println(v)
	}

	// package this data as JSON data
	// finalJson, err := json.Marshal(lcoCources)
	// marshaling means converting the data into json format

	// marshalIndent is used to format the json data
	finalJson, err := json.MarshalIndent(lcoCources, "", "\t")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", string(finalJson))
}

// consuming the json data
func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"courseName": "Django",
			"coursePrice": 300,
			"coursePlateform": "LCO",
			"courseTags": [
					"Beginner",
					"Intermediate",
					"Advance"
			]
	}`)

	// To check if the json data is valid or not
	var courseData course
	checkValid := json.Valid(jsonDataFromWeb)
	// fmt.Println(checkValid)

	// Type -1

	if checkValid {
		fmt.Println("JSON data is valid")

		json.Unmarshal(jsonDataFromWeb, &courseData)
		// unmarshal means converting the json data into go data
		// it takes two arguments first is the json data and second
		// is the address of the variable where we want to store the data

		fmt.Printf(" %#v\n", courseData)
	} else {
		fmt.Println("JSON data is not valid")
	}

	// some cases where you want to add data to key value pair

	// Type -2

	var myOnlinedata map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlinedata)
	fmt.Printf(" %#v\n", myOnlinedata)

	for k, v := range myOnlinedata {
		fmt.Printf("Key is %v && Value is %v\n", k, v)
	}
}
