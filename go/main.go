package main

import (
	"fmt"
)

var t = 44

type myVar string

func main() {
	var jr string = "TTT"
	card := "aceddd"
	fmt.Println(card)
	fmt.Println(card)
	fmt.Println(jr)

	t = 5
	fmt.Println(t)

	r := test()
	fmt.Println(r)
	var arr = []string{}
	arr = append(arr, "aaaa")
	arr = append(arr, "nknkn")

	println(arr)
	for a, k := range arr {
		println(a, k)
	}
	println(testDeck())

	t := myVar("12.3")
	println(t.myVarFunc())

	c := color("Red")

	fmt.Println(c.describe("is an awesome color"))
}

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}

func (t myVar) myVarFunc() myVar {
	return "222.33"
}

func test() int {
	return 13213
}
