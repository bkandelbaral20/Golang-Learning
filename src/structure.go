package main

import (
	"fmt"
	"strconv" //string converter package
)

// Defining person structure
type Person struct {
	//firstname string
	//lastname  string
	//city      string
	//gender    string
	//age       int

	firstname,lastname, city, gender string
	age int

}

//creating method for greeting(value receiver)
func (p Person) greeting() string {
	return "Hello, my name is " + p.firstname + " " + p.lastname + " and I am " + strconv.Itoa(p.age) + " years old."

}

//creating  hasBirthday method (pointer receiver), it is changing values to no need to return type
func (p *Person) hasBirthday()  {
	p.age++
}

//creating isMarried method (pointer receiver)
func (p *Person) isMarried(spouseLastName string)  {
	if p.gender == "male" {
		return
	} else {
		p.lastname = spouseLastName
	}
}
func main() {

	//	initializing  person using struct
	person1 := Person{ firstname: "Binjita", lastname: "Kandel", city: "San Antonio", gender: "female", age: 24 }

	//alternative way
	person2 := Person{"sanju", "kandel", "Chaitin", "female", 22}
	//person3 := Person{"Bijit", "william", "USA", "male", 30}

	fmt.Println(person1)
	fmt.Println(person2)

	// printing methods here
	person1.hasBirthday()
	person1.isMarried("Baral")

	fmt.Println(person2.greeting())

}
