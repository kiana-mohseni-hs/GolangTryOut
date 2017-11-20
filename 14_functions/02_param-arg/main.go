package main

import "fmt"

func main() {
	greet("Jane")
	greet("John")
	greet("Kiana")
}

func greet(name string) {
	fmt.Printf("Hello ")
	fmt.Println(name)
	fmt.Printf("! Welcome to the dark side!")
}

// greet is declared with a parameter
// when calling greet, pass in an argument
