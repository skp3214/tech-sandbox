package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "this need to go in a file - LearnCodeOnline.in"
	file, err := os.Create("mynewfile.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Length is: ",length)
	defer file.Close()  // when we use defer, the code will run in last, irrespective of its position
	
	readFile("mynewfile.txt")
	
}

func readFile(filename string){
	databyte,err:=os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is: ",databyte)
	fmt.Println("Text data inside the file is: ",string(databyte))
}
