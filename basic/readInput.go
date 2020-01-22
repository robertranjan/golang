package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter you name: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("You have entered your name as: %s", text)
}
