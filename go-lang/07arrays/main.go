package main

import "fmt"

func main() {

	// WITH VAR KEYWORD
	// var array_name = [length]datatypes{values} // here length is defined
	var arr1 = [5]int{1, 2, 3, 4} // partially initialzed
	fmt.Println(arr1)             // [1,2,3,4,5]
	fmt.Println(len(arr1))
	fmt.Println(arr1[0])
	fmt.Println(cap(arr1))

	// or
	// var array_name = [...]datatypes{values} // here length is inferred
	var arr2 = [...]bool{true, true, false, false, true} // fully initialized
	fmt.Println(arr2)

	// WITH := SIGN
	// array_name := [length] datatypes {values}
	scores := [5]string{"56", "55", "67", "34", "210"}
	fmt.Println(scores)
	scores[2] = "57"
	fmt.Println(scores)

	// or
	// array_name:=[...]datatypes{values}
	marks := [...]float32{23.4, 45.6, 67, 36, 33}
	fmt.Println(marks)

	// initialize only specific elements

	arr4 := [5]int{1: 20, 3: 40}
	fmt.Println(arr4)  // [0,20,0,40,0]
}
