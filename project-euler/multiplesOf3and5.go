package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print(i, ", ")
			sum += i
		}
	}
	fmt.Println("\nSum of multiples of 3 and 5 between 1 and 1000 is: ", sum)
}
