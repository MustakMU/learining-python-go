package main

import (
	"fmt"
)

type contact struct {
	num   string
	email string
}
type person struct {
	first string
	last  string
	contact
}

func main() {
	m := person{first: "as",
		last: "sd",
		contact: contact{
			num:   "efsdf",
			email: "fsdf",
		},
	}
	fmt.Println(m)

	t := person{}
	m.update("fd")
	t.print()
	m.print()
	m.update1("asd")
	m.print()
}

func (p person) print() {
	fmt.Println("ss ", p)
}

func (p *person) update(l string) {
	(*p).last = "00"
	fmt.Println("S ", p)
}

func (p person) update1(l string) {
	p.last = "99"
	fmt.Println("T ", p)
}
