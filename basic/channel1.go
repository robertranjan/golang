package main

import (
	"fmt"
	// "time"
)

func r1() {
	c1 <- "one"
	c2 <- "two"
}
// func r2() {
// 	c1 <- "one"
// 	c2 <- "two"
// }

var c1 = make(chan string)
var c2 = make(chan string)

func main() {
	go r1()

	for i :=0; i<2 ; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("Received {%v} from channel: c1\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("Received {%v} from channel: c2\n", msg2)
		}
	}

}
