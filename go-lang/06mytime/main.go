package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang");

	presentTime:=time.Now()
	fmt.Println(presentTime);
	fmt.Println(presentTime.Format("01-02-2006 14:04:05 Monday"))
	fmt.Println(presentTime.Date())
	fmt.Println(presentTime.Day())

	createdDate:=time.Date(2025,time.October,26,23,23,0,0,time.UTC);
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}