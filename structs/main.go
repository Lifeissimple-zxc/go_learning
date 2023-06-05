package main

import "fmt"

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
	// Use our newly defined type to create person instances
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// We can also do: alex := person{"Alex", "Anderson"}
	// fmt.Println(alex) // show info in stdout

	// More verbose & step by step struct declaration
	// var alex person
	// fmt.Println(alex) // show info in stdout: empty string (2) { }
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// Nested structs
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000, // comma is needed
		}, // comma is needed too!
	}
	// Update first name of a person
	// jimP := &jim // & before a variable gives us its memory address
	// jimP does not point to jim, but the memory address of jim
	jim.updateName("Jimmy")
	// fmt.Println(jim) // show info in stdout, simple
	jim.printVerbose()

}

func (pToPerson *person) updateName(newFirstName string) {
	// p.firstName = newFirstName // This creates a local copy of p
	// As a result, no modification happens for the p for which we call it

	// Using pointers to update the person
	(*pToPerson).firstName = newFirstName
	// *pToPerson gives us the actual structs sitting in the address
}

// Function with struct receiver
func (p person) printVerbose() {
	fmt.Printf("%+v", p) // show info in stdout, detailed
}
