package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	r, e := http.Get("https://www.washingtonpost.com/")
	bytes, _ := ioutil.ReadAll(r.Body)
	stringBody := string(bytes)
	defer r.Body.Close()
	fmt.Println("Error: %v \nResponse: %v \n\nbytes: %v \nstring body: %v", e, r, bytes, stringBody)
}
