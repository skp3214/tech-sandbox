package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	/*
		STRUCTS

		while arrays are used to store multiple values of the same data types into a
		single variable,
		structs are used to store multiple values of different data types into a single variable.

		DECLARE a STRUCT

		type struct_name struct{
		member1 datatype;
		member2 datatype;
		member3 datatype;
		}

	*/

	// Access Struct Members
	var first_person Person
	var second_person Person

	first_person.name = "Hege"
	first_person.age = 45
	first_person.job = "Teacher"
	first_person.salary = 6000

	second_person.name = "Cecille"
	second_person.age = 29
	second_person.job = "Engineer"
	second_person.salary = 70000

	fmt.Println(first_person) // {Hege 45 Teacher 6000}

	// Pass struct as function argument
	second_person.showPersonDetail()
}

func (person Person) showPersonDetail() {
	fmt.Println("Name: ", person.name)
	fmt.Println("Age: ", person.age)
	fmt.Println("Job: ", person.job)
	fmt.Println("Salary: ", person.salary)
}
