package main

import "fmt"

func main() {
	/*
		for loop is the only loop available in GO

		for statement1;statement2;statement3{
		// code to be executed for each iteration
		}

		break and continue works the same, usually they are used with conditions

		THE RANGE KEYWORD

		for index,value:=range array|slice|map{
		// code to be executed 
		}
	*/

	for i := 0; i < 6; i++ {
		fmt.Println(i)
	}

	fruits := [3]string{"apple","orange","banana"}

	for index, val:=range fruits{
		fmt.Println(val,index)
	}
}