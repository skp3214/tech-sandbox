package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web request")

	response,err:=http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil{
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)

	databytes,err:=io.ReadAll(response.Body)

	if err !=nil{
		panic(err)
	}

	content:=string(databytes)
	fmt.Println(content)

	defer response.Body.Close()
}