package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5}
	intCh := toChannel(data)
	sqCh := square(intCh)
	for i := range sqCh {
		fmt.Printf("%v\t", i)
	}
}

func toChannel(data []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range data {
			out <- i
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
