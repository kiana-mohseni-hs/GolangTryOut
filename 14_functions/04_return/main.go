package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
	fmt.Println(byebye("Kiana", "Mohseni"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
func byebye(fname, lname string) string {
	return fmt.Sprintf("Please don't go, %v lovely %v", fname, lname)
}