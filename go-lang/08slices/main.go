package main

import "fmt"

func main() {
	/*
		GO SLICES: Slices are similiar to arrays, but are more powerful and flexible.
		unlike arrays, the length of a slice can grow and shrink as you see fit.

		slice_name := []datatypes{values}  here [] is empty unlike array [length] or [...]
	*/

	childrens := []string{"sachin", "sameer", "siddhi"}
	fmt.Println(childrens)

	/*

	    Create Slice from an array
		var myarray = [length]datatypes{values}
		myslice := myarray[start:end]
	*/

	arr1:=[6]int{0,1,2,3,4,5,}
	myslice:=arr1[2:5];
	fmt.Println(myslice)

	fmt.Println(len(myslice));  // 3
	fmt.Println(cap(myslice));  // 4  cap = len(arr1) - start

	/*
	Create a Slice with the make() function
	slice_name := make([]type, length, capacity)

	if the capacity parameter is not defined, it will be equal to length
	*/

	myslice1 := make([]bool,5,6)
	fmt.Println(myslice1)

	/*
	Access, Change, Append and Copy Slices
	*/

	fmt.Println(myslice[0])  // 2   e.g access
	myslice[0]=5
	fmt.Println(myslice[0])  // 5   e.g change
	myslice = append(myslice,6,7,8,9,10,11,12,13,14,15) // append
	fmt.Println(myslice)

	newslice:=arr1[3:5];
	newappended := append(myslice,newslice...)
	fmt.Println(newappended)

	/*
	Memory Efficiency
	if the array is large and you need only a few elements, it is better to copy those elements
	using the copy() function.
	*/

	numbers:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,15,15}

	neededNumbers := numbers[0:6]
	numbersCopy:= make([]int,5)
	copy(numbersCopy,neededNumbers)
	fmt.Println(numbersCopy) // becaue of length 5 only 5 elements will be copied

}
