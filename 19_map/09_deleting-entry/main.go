package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"zero":  "Good morning!",
		"one":   "Bonjour!",
		"two":   "Buenos dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(myGreeting)
	fmt.Println(len(myGreeting))
	delete(myGreeting, "two")
	fmt.Println(myGreeting)
	fmt.Println(len(myGreeting))
	myGreeting["harold"]="what?"
	fmt.Println(myGreeting)
	fmt.Println(len(myGreeting))
	fmt.Println((myGreeting["two"]))

}
