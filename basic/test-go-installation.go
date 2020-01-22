package main

import (
	"fmt"
	"os"

	"github.com/robertranjan/golang/mix"
)

func main() {
	var name string = "Robert"
	var age int = 45
	var sex byte = 'm'

	rn := mix.Random()
	fmt.Printf("Hello %v is %v years old and he is a '%c' \nrn: %v", name, age, sex, rn)
	fmt.Printf("\033[31m")
	mix.ArrayExample()
	mix.Ls()
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println(os.Stat("/Users/robertrt/go/code/src/github.com/robertranjan/golang/basic/test-go-installation.go"))

}
