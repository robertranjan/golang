package main

import "fmt"

// Student ...
type Student struct {
	name  string
	age   uint
	sex   string
	marks [5]int
}

func main() {
	population := map[string]int{
		"California": 10100000,
		"Iowa":       3300300,
		"Texas":      5555445,
	}

	for k, v := range population {
		fmt.Println("\033[31m", k, v)
	}

	s1 := []Student{
		Student{name: "Robert", age: 45, sex: "Male", marks: [5]int{45, 54, 23, 56, 65}},
		Student{name: "Berjina", age: 17, sex: "Female", marks: [5]int{76, 54, 23, 56, 65}},
		Student{name: "Brian", age: 11, sex: "Male", marks: [5]int{45, 54, 67, 56, 65}},
	}
	for k, v := range s1 {
		fmt.Println(k, v)
	}
	fmt.Println("one more...", s1[0].marks[0])
	// fmt.Println("\033[32m", s1)

	// fmt.Println("\033[31m")
	// s2 := "Robert Ranjan"
	// for k, v := range s2 {
	// 	fmt.Printf("%v %v | ", k, string(v))
	// }
	// fmt.Println("\033[0m")
}
