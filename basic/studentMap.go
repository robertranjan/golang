package main

import "fmt"

func main ( ) {
	var grades = map[string] []int {
		"Robert": []int{ 45, 56, 65, 72},
		"Brian1": []int{ 65, 76, 75, 82},
		"Berjina1": []int{69, 79, 79, 89},
	}
	fmt.Println(grades["Robert"][0])	
}