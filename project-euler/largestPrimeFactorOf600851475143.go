package main

import "fmt"

func main() {
	number := 600851475143
	target := number / 2
	for i := target; i >= 2; i-- {
		if number%i == 0 {
			fmt.Println("Number: ", number, " can be divided by : ", i)
			break
		}
		if i%100000000 == 0 {
			fmt.Printf("Currently working on %v\n", i)
		}
	}
}
