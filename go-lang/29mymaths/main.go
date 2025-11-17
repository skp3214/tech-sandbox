package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.5

	// fmt.Println("The sum is: ", myNumberOne + int(myNumberTwo))

	// randon number

	// fmt.Println(rand.Intn(5))

	// randon from crypto

	myRandonNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandonNum)

}
