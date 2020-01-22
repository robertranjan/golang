package main

import "fmt"

func main() {
	var s []int

	for i := 1; i < 25; i++ {
		s = append(s, i)
		fmt.Printf("len: %v, cap: %v,  array: %v\n", len(s), cap(s), s)
	}
}
