package main

import (
	"fmt"
	"sync"
	// "time"
)



var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

func sayHello(msg string) {
	fmt.Printf("%s %v\n", msg, counter)
	m.RUnlock()
	wg.Done()
}


func main() {

	for i :=0 ; i < 10 ; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello("Hello Robert")
		m.Lock()
		go increment()
	}
	wg.Wait()
}
