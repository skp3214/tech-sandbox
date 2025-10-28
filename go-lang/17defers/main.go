package main

import "fmt"

func main() {

	/*
	if defer is used, this function will run in last and if more than two defer is used
	LIFO is used for there execution
	*/
	
	defer fmt.Println("World")
	defer fmt.Println("Hii")
	fmt.Println("Hello")
	myDefer()
}

func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)  // this will reverse the printing
	}
}