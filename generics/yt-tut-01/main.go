package main

import (
	"fmt"
)

func main() {
	// // int
	// list := []int{1, 2, 3, 4, 5}
	// item, _ := strconv.Atoi(os.Args[1])
	// h := HasInt(list, item)
	// fmt.Printf("int: %v in %v: %v\n", item, list, h)

	// // float
	// list2 := []float64{1, 2, 3, 4, 5}
	// item2, _ := strconv.ParseFloat(os.Args[1], 64)
	// h2 := HasFloat(list2, item2)
	// fmt.Printf("float: %v in %v: %v\n", item2, list2, h2)

	// // float | int
	// list3 := []float64{1, 2, 3, 4, 5}
	// item3, _ := strconv.ParseFloat(os.Args[1], 64)
	// h3 := Has(list3, item3)
	// fmt.Printf("int|float: %v in %v: %v\n", item3, list3, h3)

	// // string
	// list4 := []string{"1", "2", "3", "4", "5", "robert"}
	// // list4 := []int{1, 2, 3}
	// item4 := os.Args[1]
	// h4 := Has(list4, item4)
	// fmt.Printf("string: %v in %v: %v\n", item4, list4, h4)

	fmt.Printf("%v\n", Has([]string{"a", "b", "c"}, "a"))
	fmt.Printf("%v\n", Has([]int{1, 2, 3}, 5))

}

func Has[T comparable](list []T, val T) bool {
	for _, i := range list {
		if i == val {
			return true
		}
	}
	return false
}

// func Has[T int | float64](list []T, val T) bool {
// 	for _, i := range list {
// 		if i == val {
// 			return true
// 		}
// 	}
// 	return false
// }

// func HasInt(list []int, val int) bool {
// 	for _, i := range list {
// 		if i == val {
// 			return true
// 		}
// 	}
// 	return false
// }
// func HasFloat(list []float64, val float64) bool {
// 	for _, i := range list {
// 		if i == val {
// 			return true
// 		}
// 	}
// 	return false
// }
