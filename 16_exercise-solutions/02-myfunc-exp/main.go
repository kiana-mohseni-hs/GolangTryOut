package main

import "fmt"

func main () {
//	var myint int
//	var mybool bool
//	myint, mybool = exercise (16)
//	fmt.Println(exercise(3))
//fmt.Println(myint, mybool)


	half := func (n int) (int, bool) {
	return (n/2), n%2 == 0

	}

	fmt.Println(half(9))
	}

