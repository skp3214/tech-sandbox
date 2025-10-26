package main

import "fmt"

func main() {
	/*
		IF Statement

		if condition {
		// code to be executed
		}
	*/

	if 8 > 5 {
		fmt.Println("8 is greated than 5")
	}

	// IF ELSE statement
	if true {
		fmt.Println("You are right")
	}else{
		fmt.Println("You are wrong")
	}

	// ELSE IF statement

	if 8<5 {
		fmt.Println("Execute if part")
	} else if 8>5{
		fmt.Println("Execute else if part")
	}else{
		fmt.Println("Otherwise execute else part")
	}

	// Nested if else are also allowed
}
