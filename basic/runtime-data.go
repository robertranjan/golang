package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {

	show := func(key string) {
		if envVal, ok := os.LookupEnv(key); !ok {
			fmt.Println("Key not found\n")
		} else {
			fmt.Printf("%s=%s\n", key, envVal)
		}
	}

	fmt.Println("Go max procs: ", runtime.GOMAXPROCS(-1))
	fmt.Println("Num of CPUs: ", runtime.NumCPU())
	fmt.Println("Go ROOT: ", runtime.GOROOT())
	fmt.Println("Go path: ", os.Getenv("GOPATH"))

	os.Setenv("TEMP1", "Temp1_value")
	os.Setenv("TEMP2", "")

	show("TEMP1")
	show("TEMP2")
	show("MISSING_KEY")

}
