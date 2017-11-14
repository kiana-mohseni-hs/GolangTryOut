package main

import "fmt"

func main() {

	a := 0
	fmt.Println(42)
	fmt.Printf("%d - %b - %x\n", 42, 42, 42)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d - %b - %x\n", a , a, a)
		a += 1
	}
}

