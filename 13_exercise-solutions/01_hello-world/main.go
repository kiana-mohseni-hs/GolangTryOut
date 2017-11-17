package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	for i := 5000; i <= 10100; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
