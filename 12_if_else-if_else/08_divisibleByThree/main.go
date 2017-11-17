package main

import "fmt"

var b func() = main

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
			a := 3
			fmt.Printf("%v", a)
		}
	}
}
