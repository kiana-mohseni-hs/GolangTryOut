package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	myslice := make([]string, 100)
	myslice[0] = "kiana"
	myslice[99] = "the end"
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
	fmt.Println(myslice)
}
