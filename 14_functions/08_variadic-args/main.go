package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	var total float64
	for i, v := range sf {
		total += v
		fmt.Println ("i = %v, v = %v", i, v, total, len(sf), sf)
	}
	return total / float64(len(sf))
}
