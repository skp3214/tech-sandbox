package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Pointers")

    var myptr *int 
	fmt.Println("Value of pointer is,",myptr)

	myNumber:=23

	var ptr = &myNumber

	fmt.Println("Value of pointr is,",ptr)  // stores address of myNumber variable
	fmt.Println("Value of pointr is,",*ptr) // * is used to see the actual value of pointer

	*ptr = *ptr*2
	fmt.Println("new value is,",myNumber)	
}