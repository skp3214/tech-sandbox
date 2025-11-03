package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web request in golang")
	// performGET()
	performPost()
}

func performGET() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)
	content, err := io.ReadAll(response.Body)
	fmt.Println(string(content))

	var responseString strings.Builder
	contentnew, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(contentnew)
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

}

func performPost() {
	const myurl = "http://localhost:3000/post"

	// fake json payload

	requestBody := strings.NewReader(`
	{
	  "coursename":"Let's go with golang",
	  "price":0,
	  "platform":"learnCodeOnline.in"
	}
	`)

	response,err:=http.Post(myurl,"application/json",requestBody)

	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()

	content,_:=io.ReadAll(response.Body)

	fmt.Println(string(content))
}
