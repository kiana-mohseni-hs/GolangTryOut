package main

import "fmt"

func main() {
	age := 44
	newage := changeMe(&age)
	fmt.Println("age ", age) // 44
	fmt.Println("newage ",newage)
}

func changeMe(z *int) int {
	fmt.Println(z)
	*z = 24
	return (*z)
}

// when changeMe is called on line 8
// the value 44 is being passed as an argument
