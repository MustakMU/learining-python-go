package main

import (
	"fmt"
	"sort"
)

/**
Write a program to sort a list of dictionaries using Lambda.
Original list of dictionaries :
[{'make': 'Nokia', 'model': 216, 'color': 'Black'}, {'make': 'Mi Max', 'model': '2', 'color': 'Gold'}, {'make': 'Samsung', 'model': 7, 'color': 'Blue'}]

**/

func main() {

	var phoneModels = []struct {
		make  string
		model int
		color string
	}{
		{"Nokia", 216, "Black"},
		{"Sony", 222, "White"},
		{"Motorola", 200, "Gray"},
		{"Samsung", 220, "Blue"},
	}
	sort.SliceStable(phoneModels, func(i, j int) bool {
		return phoneModels[i].model < phoneModels[j].model
	})
	fmt.Println(phoneModels)
}
