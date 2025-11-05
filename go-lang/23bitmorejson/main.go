package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	DecodeJSON()
}

func EncodeJSON() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"AngularJS Bootcamp", 299, "LearnCodeOnline.in", "abc1234", []string{"angular", "js"}},
		{"JavaScriptJS Bootcamp", 299, "LearnCodeOnline.in", "abc1235", []string{"google", "js"}},
		{"NodeJS Bootcamp", 299, "LearnCodeOnline.in", "abc1236", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
	{
        "coursename": "JavaScriptJS Bootcamp",
        "Price": 299,
        "website": "LearnCodeOnline.in",
        "tags": ["google","js"]
        }
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k,v:=range myOnlineData{
		fmt.Printf("Key is %v and value is %v and Type is:%T\n",k,v,v)
	}
}
