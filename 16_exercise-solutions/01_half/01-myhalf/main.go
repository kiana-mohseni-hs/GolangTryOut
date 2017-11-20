package main

import "fmt"

func main () {
	fmt.Println(exercise(30000))

}

func exercise (x int) (int, bool) {
	 var y bool
	 if (x%2 == 0) {
	 	y = true
	 }
	return (x/2), y

}



