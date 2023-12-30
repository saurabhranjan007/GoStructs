package main

import "fmt"

// contactInfo => will be embedded with struct of type person
type contactInfo struct {
	email   string
	zipCode int
}

// Define Person struct (creates a blue print of Person Object - in context of OOPS)
type person struct {
	firstName string
	lastName  string
	contact   contactInfo // here we can also avoid using ▶️ "contact contactInfo and just use contactInfo"
}

// Note📝: Struct embedding, associating one struct with another (Class Composition in context of OOPS)
// 		Struct embedding - embedding one struct into another.

func main() {
	// One way to creating new struct of type person 1️⃣
	// saurabh := person{"Saurabh", "Ranjan"}

	// Another way of creating new struct of type person 2️⃣
	// saurabh := person{firstName: "Saurabh", lastName: "Ranjan"}

	// Another way of creating new struct of type person (with no data assignment) 3️⃣
	// var saurabh person
	// saurabh.firstName = "Saurabh"
	// saurabh.lastName = "Ranjan"
	// fmt.Println("Initial:", saurabh)
	// // New way of printing each data param of a struct
	// fmt.Printf("%+v", saurabh)

	sana := person{
		firstName: "Sana",
		lastName:  "Singh",
		contact: contactInfo{
			email:   "sana@singh.com",
			zipCode: 110092,
		},
	}

	// fmt.Printf("%+v", sana)
	// fmt.Println("\n", sana)
	// Calling the custom print statement
	sana.print()

	// ▶️ Updating the person data; we require to create a Pointer in order to update the struct of type person (here)
	sanaPointer := &sana
	sanaPointer.updatePersonName("Suryavanshi", "Sana")

	sana.print()

}

// Creating a custom print function for the structs of type person
func (p person) print() {
	fmt.Println("Custom Print Statement Called!!")
	fmt.Printf("%+v", p)
}

// Custom function to update the firstName and lastName
// func (p person) updatePersonName() {
// 	fmt.Println("Function called to update first and last name.")
// }
func (pointerToPerson *person) updatePersonName(newFirstName string, newLastName string) {
	fmt.Println("Custom function to update the first and last name of the struct data of type person")
	(*pointerToPerson).firstName = newFirstName
	(*pointerToPerson).lastName = newLastName
}
