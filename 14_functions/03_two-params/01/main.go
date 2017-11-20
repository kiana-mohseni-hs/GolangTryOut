package main

import "fmt"

func main() {
	greet("Jane", "Doe")
}

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
	fmt.Printf("Hello %v %v!", fname, lname)
}
