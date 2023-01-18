package main

import 
	f "fmt"


type Person struct {
	firstName string
	lastName string
	age int
}

type People []Person

var people = People{
	Person {
		firstName: "John",
		lastName: "Doe",
		age: 61,
	},
	Person {
		firstName: "Mary",
		lastName: "Hary",
		age: 32,
	},
}


func main() {
	showPeopleData()
}


func showPeopleData() {
	for i := 0; i < len(people); i++ {
		var fullName = people[i].firstName +" "+ people[i].lastName
		f.Printf("person %s has %d years old\n", fullName, people[i].age)
	}
}