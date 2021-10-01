package main

import (
	"encoding/json"
	"fmt"
)

/**
Write a  program to convert JSON data to Python object.
**/
func main() {
	jsonStr := "{\"make\": \"Nokia\", \"model\": 216, \"color\": \"Black\"}"
	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	println(dat["make"].(string))

	println(dat["model"].(float64))
	println(int(dat["model"].(float64)))

	var i interface{}
	var count int = 33
	i = count
	fmt.Println(i)
	i = "Hello World!!"
	fmt.Println(i)
}
