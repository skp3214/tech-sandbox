package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 - 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	// numRating, err := strconv.ParseFloat(input, 64); "4\n" error
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64); 

	if err != nil {
		fmt.Println(err)
		// panic(err);
	}else{
		fmt.Println("Addded 1 to your rating: ", numRating+1)
	}
}
