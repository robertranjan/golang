package main

import "fmt"

// Person ...
type Person struct {
	fname string
	lname string
}

func (p Person) speak() int {
	fmt.Printf("My name is %v %v\n", p.fname, p.lname)
	return 0
}

func main() {
	p1 := Person{fname: "Robert", lname: "Thanu"}
	p1.speak()

}
