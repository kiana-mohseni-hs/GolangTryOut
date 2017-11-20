package main

import (
	"fmt"
)

func main() {
	student1 := []string{"kiana", "mohseni", "female", "180 W 16th Ave"}
	student2 := []string{"claudia", "molina", "female", "180 W 16th Ave"}
	student3 := make([]string, 3)
	student3[0] = "steve"
	student3[1] = "tate"
	students := [][]string{student1}
	students = append (students, student2)
	students = append (students, student3)
	fmt.Println(student1)
	fmt.Println(students)
	fmt.Println(student1 == nil)
}
