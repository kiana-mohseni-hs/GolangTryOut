package main

import "fmt"

func main() {

	a := 1000000
	fmt.Println(42)
	fmt.Printf("%d - %b - %x\n", 42, 42, 42)
	for i := 1000000; i < 1000020; i++ {
		fmt.Printf("%d - %b - %x\n", a , a, a)
		a += 1
	}
}

