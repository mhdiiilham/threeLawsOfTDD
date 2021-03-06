package main

import (
	"testing"
)

func TestCreateNewPerson(t *testing.T) {
	Person := NewPerson("Muhammad", "Ilham", 19)

	if Person.Firstname != "Muhammad" {
		t.Errorf("Expected firstname of Person is Muhammad, but got:%v ", Person.Firstname)
	}

	if Person.Lastname != "Ilham" {
		t.Errorf("Expected lastname of Person is Ilham, but got:%v ", Person.Lastname)
	}

	if Person.Age != 19 {
		t.Errorf("Expected Age of Person is 19 but got: %v", Person.Age)
	}
}
