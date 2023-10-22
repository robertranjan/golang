package main

import (
	"fmt"
	"strconv"
)

func demoBufChanWithForSelect() {
	fmt.Printf("\033[34mdemoBufChanWithForSelect\n")
	ch1 := make(chan string, 3)
	ch2 := make(chan string, 3)
	for i := 0; i < 3; i++ {
		ch1 <- strconv.Itoa(i)
	}
	for i := 10; i < 13; i++ {
		ch2 <- strconv.Itoa(i)
	}
	close(ch1)
	close(ch2)

	l := len(ch1) + len(ch2)
	// var ch1Item, ch2Item string
	for i := 0; i < l; i++ {
		select {
		case item, ok := <-ch1:
			if ok {
				fmt.Printf("read from channel ch1: %v\n", item)
			}
		case it2, ok := <-ch2:
			if ok {
				fmt.Printf("read from channel ch2: %v\n", it2)
			}
		}
	}
}

func main() {

	demoMultiChannels()
	demoBufChanWithForSelect()
	// fmt.Printf("reading data from channel: %v\n", <-c)
	// fmt.Printf("reading data from channel: %v\n", <-c)
	// fmt.Printf("reading data from channel: %v\n", <-c)

	// simple go routine
	// go someFunc("hello... 1")
	// go someFunc("hello... 2")
	// go someFunc("hello... 3")
	// fmt.Println("hi")
	// time.Sleep(10 * time.Millisecond)
}

// // simple go routine
// func someFunc(num string) {
// 	fmt.Println(num)
// }

func channelExample1(ch chan string, i int) {
	ch <- strconv.FormatInt(int64(i), 10)
}

func demoMultiChannels() {
	fmt.Printf("\033[33mdemoMultiChannels\n")
	c := make(chan string)
	for i := 0; i < 10; i++ {
		go channelExample1(c, i)
	}

	d := make(chan string)
	for i := 0; i < 10; i++ {
		go channelExample1(d, i)
	}

	i := 0
	for {
		i++
		if i > 20 {
			break
		}
		select {
		case d1 := <-c:
			fmt.Printf("reading data from channel c: %v, i: %v\n", d1, i)
		case d2 := <-d:
			fmt.Printf("reading data from channel d: %v, i: %v\n", d2, i)
		}
	}
}
