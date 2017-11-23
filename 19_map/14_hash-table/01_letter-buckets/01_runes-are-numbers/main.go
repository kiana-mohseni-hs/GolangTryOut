package main

import "fmt"

func main() {
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)

	var this_is_a_rune rune
	this_is_a_rune = 'K'
	fmt.Println(this_is_a_rune)
	fmt.Println( string(this_is_a_rune))
	fmt.Println( rune(string(this_is_a_rune)[0]))

}
