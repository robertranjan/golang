package mix

import (
	// "bytes"
	"crypto/rand"
	"fmt"
)

// Random() ; exporting this one
func Random() int {
	// c := 10
	b := make([]byte, 10)
	// b := [...]byte{0,0}
	fmt.Printf("%v\n", b)
	_, err := rand.Read(b)
	fmt.Printf("%v\n", b)
	if err != nil {
		fmt.Println("error:", err)
		return 1
	}
	// The slice should now contain random bytes instead of only zeroes.
	fmt.Printf("%v\n", b)
	// d := make([]int, 5)
	// _, e2 := rand.Int(d, 1000)
	// if e2 != nil {
	// 	fmt.Println("Error: ", e2)
	// 	return
	// }
	// fmt.Println(d)

	sum := func(x []byte) int { return int(x[0]) + int(x[1]) }(b)
	// sum := func(a, b int) int { return a+b } (3, 4)
	fmt.Println("sum: ", sum)
	return sum
}
