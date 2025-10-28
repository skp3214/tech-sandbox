package main

import "fmt"

func main() {
	/*
		Create a function

		func functionName(){

		}

		call a function
		functionName();

		Parameter and arguments

		func functionName(param1 type, param2 type, param3 type){

		}

		Return Types

		func functionName(param1 type, param2 type, param3 type) type {

		}

		Named Return Values

		func functionName(param1 type, param2 type, param3 type) (variableName type) {

		}

		Multi return values

		func functionName(param1 type, param2 type, param3 type) (variableName1 type, variableName2 type) {

		}

		Recursion Functions


	*/

	familyName("Sachin")
	myFunction(5, 6)
	addTwoNum(6, 7)

	_, b := returnTwoValues(6, "Hello ")
	fmt.Println(b)

	result:=factorial_recursion(5)
	fmt.Println(result)

	// IIFE

	func (){
		fmt.Println("Immediately Invoked Function Expression")
	}()

	// multiple values

	valueProAdder:=proAdder(5,6,7,8);
	fmt.Println(valueProAdder)
}

func familyName(name string) {
	fmt.Println("Your name is: ", name)
}

func myFunction(x int, y int) int {
	return x + y
}

func addTwoNum(x int, y int) (result int) {
	result = x + y
	return result
}

func returnTwoValues(x int, y string) (result int, text1 string) {
	result = x + 5
	text1 = y + "World!"
	return
}

func factorial_recursion(x int) (y int) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	}else{
		y=1
	}
	return
}


func proAdder(values ...int)int{
	total:=0
	for _,val:=range values{
		total+=val
	}
	return total
}