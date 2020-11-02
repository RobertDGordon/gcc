package main

import (
	"fmt"
	"strconv"
)

//Person is a person
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

//Greeting method, returns string
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastName = spouseLastName
	}

}

func main() {
	//Init person with struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{"Bob", "Johnson", "DC", "m", 30}

	// fmt.Println(person1)

	person2.age++

	// fmt.Println(person2.age)

	person1.hasBirthday()
	person2.getMarried("Thompson")
	person1.getMarried("Williams")

	fmt.Println(person2.greet())
}
