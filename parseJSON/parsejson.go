package main

import (
	"encoding/json"
	"fmt"
)

//AboutMe is a struct which contains details of a User
type AboutMe struct {
	Name string `json:"name"`

	Age int `json:"age"`

	Country string `json:"country"`

	Skills []string `json:"skills"`
}

func main() {

	myJSONString := `
	
	{
		"name": "Shreyash Bukkawar",
		"age": 23,
		"country": "India",
		"skills": ["Python","Golang","jenkins","Web Developemnt"]

	}
	
	
	`
	var aboutMe AboutMe

	json.Unmarshal([]byte(myJSONString), &aboutMe)

	fmt.Println("%+v", aboutMe)

}
