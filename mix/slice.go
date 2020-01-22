package mix

import (
	"fmt"
	"math/rand"
)

var slice_1 []int

func printSlice() {
	fmt.Println("Slice_1: ", slice_1, "\t|\tLength: ", len(slice_1), "\t|\tCapacity of slice: ", cap(slice_1))
}

func appendSlice(tmp_slice []int) {
	x := rand.Intn(10)
	tmp_slice = append(tmp_slice, x)
}

func del2ndElementFromSlice(tmp_slice []int) {
	var new_slice []int
	new_slice = append(new_slice, tmp_slice[2:]...)
	new_slice = append(new_slice, tmp_slice[:1]...)
}

func slice() {
	slice_1 = []int{1, 2, 3, 4, 5}
	printSlice()
	slice_1 = append(slice_1, 100)
	printSlice()

	appendSlice(slice_1)
	printSlice()

	del2ndElementFromSlice(slice_1)
	printSlice()
}
