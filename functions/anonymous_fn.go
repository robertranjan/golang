package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeFunc() func(int64, int64) int64 {
	return func(a, b int64) int64 {
		if a-b >= 0 {
			return a
		}
		return b
	}
}

func findMax(a, b int64, callback func(int64, int64) int64) int64 {
	return callback(a, b)
}

func main() {
	max := makeFunc()
	a := rand.Int63n(time.Now().UnixNano())
	b := rand.Int63n(time.Now().UnixNano())
	fmt.Printf("\ta: %v \tb: %v Max: %v\n", a, b, max(a, b))
	fmt.Printf("\ta: %v\tb: %v \tMax: %v \ttime.Now().UnixNano(): %v\n", a, b, findMax(a, b, max), time.Now().UnixNano())
}
