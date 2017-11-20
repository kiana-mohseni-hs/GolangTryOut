package main

import "fmt"

func main () {
	max := highest(1,5,2,3,9,4)
	fmt.Println (max)
	fmt.Println ((false&&true)||(true&&false)||(false&&false))
}

func highest (list ...int) int {
	var m int
	for p, i := range list{
		if i > m {
			m = i
		}
		fmt.Println(p,i,m)

	}
		return m
	}

