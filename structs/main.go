package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	//contact contactInfo // you can also write only contactInfo. Go will understand what you mean :)
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"
	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)

	// Every single line must have a comma even if it's the last line
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// if you write contactInfo in person withour contact, you need to point out that the contactInfo is contactInfo.
		//contact: contactInfo{
		//	email: "jim@gmail.com",
		//	zipCode: 94000,
		//},
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	/*
		jimPointer := &jim
		jimPointer.updateName("jimmy")
		jim.print()
	*/
	// A shorter version. Go will understand what we want :)
	jim.updateName("jimmy")
	jim.print()

	/*
		Value types:
			- int
			- float
			- string
			- bool
			- struct
		Reference types:
			- slices
			- maps
			- channels
			- pointers
			- functions
		# Reference type'ları değiştirmek için pointer'a ihtiyaç yok :)
	*/
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
