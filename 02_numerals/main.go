package main

import "fmt"

func main() {

	a := 60
	fmt.Println(42)
	fmt.Printf("%d - %b - %x\n", 42, 42, 42)
	for i := 1; i < 70; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", a , a, a, a)
		a += 1
	}
}

