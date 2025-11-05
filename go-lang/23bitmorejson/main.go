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
	EncodeJSON()
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
