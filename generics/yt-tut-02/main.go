package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Weight float64
}

// Set is a user type  to hold a map[comparable]
type Set[T comparable] map[T]struct{}

// NewSet create a set from a variadic args of any(comparable) type
func NewSet[T comparable](values ...T) Set[T] {
	set := make(Set[T], len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}
	return set
}

// method
func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func main() {

	// NewSet
	fmt.Printf("set: %v\n", NewSet[int](1, 2, 3, 4, 5, 5, 3, 2))
	fmt.Printf("set: %v\n", NewSet[string]("a", "b", "c", "d", "c", "a"))

	intSet := NewSet(1, 2, 3, 4, 5, 5, 3, 2)
	fmt.Printf("set has 5 is %v\n", intSet.Has(5))
	fmt.Printf("set has 15 is %v\n", intSet.Has(15))

	fmt.Printf("list has 3 is %v\n", Has([]int{1, 2, 3, 4, 5}, 3))
	fmt.Printf("list has 6 is %v\n", Has([]int{1, 2, 3, 4, 5}, 6))

	fmt.Printf("list has 3.4 is %v\n", Has([]float64{1.2, 2.3, 3.4, 4.5, 5.6}, 3.4))
	fmt.Printf("list has 6.4 is %v\n", Has([]float64{1.2, 2.3, 3.4, 4.5, 5.6}, 6.4))

	fmt.Printf("list has a is %v\n", Has([]string{"a", "b", "c", "d", "e"}, "a"))
	fmt.Printf("list has x is %v\n", Has([]string{"a", "b", "c", "d", "e"}, "x"))

	fmt.Printf("newList[int]: %#v\n", NewEmptyList[int]())
	fmt.Printf("newList[byte]: %#v\n", NewEmptyList[byte]())
	fmt.Printf("newList[Person]: %#v\n", NewEmptyList[Person]())

	// multiple type parameters
	PrintThings(10, 12, "x", 23)
	PrintThings("a", "robert", 10, 48)
}

func Has[T comparable](list []T, value T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// IsEmpty check a list of any type is empty
func IsEmpty[T any](list []T) bool {
	return len(list) == 0
}

// NewEmptyList create a new empty list of any type
func NewEmptyList[T any]() []T {
	return make([]T, 0)
}

// multiple type parameters
func PrintThings[A, B any, C ~int](a1, a2 A, b B, c C) {
	fmt.Printf("PrintThings: %v(%T) | %v(%T) | %v(%T) | %v(%T)\n", a1, a1, a2, a2, b, b, c, c)
}
