package main

import "fmt"

func main() {
	/*
		Go Maps

		Maps are used to store data values in key:value pairs

		A map is an unordered and changeable collection that does not allow duplicates.
		Maps hold references to an underlying hash table
		// Create maps using var and :=

		var a = map[keytype]valuetype{key1:value1, key2:value2, key3:value3}
		b := map[keytype]valuetype{key1:value1, key2:value2}
	*/

	var a = map[string]string{"brand": "ford", "model": "Mustang"}
	b := map[string]int{"Osle": 1, "Bergen": 2}

	fmt.Println(a)
	fmt.Println(b)

	// using make function
	var c = make(map[int]int)
	c[1] = 2
	c[2] = 7
	c[3] = 6
	c[4] = 5

	fmt.Println(c)
	fmt.Println(c[1])
	
	// create an empty map
	// var a map[keytype]valuetype
	
	var d map[string]float32
	fmt.Println(d) // map[]
	
	delete(c,2)   // value with key=2 is deleted
	fmt.Println(c)
	
	// Check for specific elements in a Map
	// val, ok := map_name[key]

	val1, ok1 := a["brand"] // checking for existing key and its value
	val2, ok2 := a["price"] // checking for non-existing
	_,ok3 := a["model"] // only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(ok3)

	// Iterate Over Maps

	for k,v := range a{
		fmt.Println(k,v)
	}

}
