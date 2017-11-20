package main

import (
	"fmt"
)

func main() {
	sl := []int{0,1}
	sls := new([100]int)[0:50]
	sln := make ([]int, 50, 100)
	fmt.Printf("%T \n", sl)
	fmt.Println(sl)
	fmt.Println(sls)
	fmt.Println(sln)
	sln = append (sln, 9)
	fmt.Println(sln)

}
