package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // contactInfo is a struct
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// alex.print()

	// var alex2 person
	// alex2.firstName = "Alex2"
	// alex2.lastName = "Anderson2"
	// fmt.Println(alex2)
	// fmt.Printf("%+v", alex2)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "Email",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

// Update first name
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// receiver function
func (p person) print() {
	fmt.Printf("%+v", p)
}
