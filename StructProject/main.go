package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//name := person{"Asir", "Tamboli"}
	// name := person{firstName: "Asir", lastName: "Tamboli"}

	// var name person

	// name.firstName = "Asir"
	// name.lastName = "Tamboli"

	person1 := person{
		firstName: "ABC",
		lastName:  "DEF",
		contact: contactInfo{
			email:   "ABC@gmail.com",
			zipcode: 123,
		},
	}

	personPointer := &person1

	personPointer.updateName("XYZ")

	person1.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName

}
