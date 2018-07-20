package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Blah",
		contactInfo: contactInfo{
			email:   "alex@blah.com",
			zipCode: 3000,
		},
	}
	alex.updateName("alexander")
	alex.print()

}
func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
