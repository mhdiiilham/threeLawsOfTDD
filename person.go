package main

// Person ...
type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func NewPerson(firstname, lastname string, age int) Person {
	person := Person{
		Firstname: firstname,
		Lastname:  lastname,
		Age:       age,
	}

	return person
}
