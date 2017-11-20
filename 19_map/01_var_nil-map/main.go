package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"claudia": "hola",
		"steve": "hello",
	}
	var MGreeting = make(map[string]string)

	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
	MGreeting["kiana"] = "salaam"
	fmt.Println(MGreeting)
	fmt.Println(myGreeting)

}

// add these lines:
/*
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
*/
// and you will get this:
// panic: assignment to entry in nil map
