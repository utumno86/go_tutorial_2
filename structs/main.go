package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@party.com",
			zipCode: 90210,
		},
	}

	jim.print()
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerPerson *person) updateName(newFirstName string) {
	(*pointerPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
