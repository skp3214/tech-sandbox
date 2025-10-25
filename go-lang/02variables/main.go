package main

import "fmt"

// numberOfPeople:= 30000 not allowed (no type declaration)

var numberOfPeople int = 300 // allowed

const LoginToken string = "aasdfghj" // Public


func main() {
	var username string = "sachin";
	    fmt.Println(username);
	fmt.Printf("variable id of type: %T \n ", username);

	var isLoggedin bool = true
	fmt.Println(isLoggedin);
	fmt.Printf("Variable is of type: %T \n ", isLoggedin);

	var smallValue uint8 = 255
	fmt.Println(smallValue);
	fmt.Printf("Variable is of type: %T \n",smallValue);

	var smallFloat float32 = 255.45675432113365432
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type: %T \n",smallFloat);

	var largeFloat float64 = 255.45675432113365432
	fmt.Println(largeFloat);
	fmt.Printf("Variable is of type: %T \n",largeFloat);
	
	var intVariable int
	fmt.Println(intVariable);
	fmt.Printf("Variable is of type: %T \n",intVariable);
	
	// implicit type

	var website = "user.com"
	fmt.Println(website);
	fmt.Printf("Variable is of type: %T \n",website);

	// website = 3 // not allowed to assign other type

	// no var type

	numberOfUser:= 30000
	fmt.Println(numberOfUser);

}