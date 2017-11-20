package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	println("x is ",x)
	y := factorial(x-1)
	println("y is ",y)
	println("return", x, y)
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
}
