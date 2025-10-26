package main

import "fmt"

func main() {
	/*
			the switch statement in GO is similiar to the ones in c, c++, java, javascript and PHP.
			The difference is that it only runs the matched case so it does not need a break statement.

	  // SINGLE CASE SWITCH
			switch expression {
		case x:
			// code block
		case y:
			// code block
		default:
			// code block
			}
			
			// MUTLI CASE SWITCH
			switch expression {
		case x,y,z:
			// code block
		case a,b:
			// code block
		default:
			// code block
			}
			*/

	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Thursday")
	default:
		fmt.Println("Sunday")
	}

	day = 5

	switch day {
	case 1,5,3:
		fmt.Println("Monday")
	case 4,2:
		fmt.Println("Thursday")
	default:
		fmt.Println("Sunday")
	}

}
