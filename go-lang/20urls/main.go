package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://jsonplaceholder.typicode.com/todos/1?limit=5&value=6"

func main() {
	fmt.Println("Welcoming to handling URLs in golang")

	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["limit"])

	for _, val := range qparams {
		fmt.Println("Param is, ", val)
	}

	partsofUrl := &url.URL{
		Scheme: "https",
		Host: "jsonplaceholder.typicode.com",
		Path: "/todos",
		RawQuery: "limit=5",
	}

	anotherUrl:=partsofUrl.String()
	fmt.Println(anotherUrl)
}
