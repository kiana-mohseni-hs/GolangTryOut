package main

import "fmt"

func main () {
	foo(1,2)
	foo(1,2,3)
	aSlice := []int{1,2,3,4}
	foo(aSlice...)
}

func foo (n ...int){
	fmt.Printf("%T",n)
	fmt.Println(n)
}
