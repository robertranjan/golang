package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL: %v \nName: %v ", r.URL, r.FormValue("name"))
	// ( w, "Hello there!\n\n", "\n\nURL: ", r.URL, "\n\nName: ", r.FormValue("name"), "\nAge:", r.FormValue("age"))
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Robert! \n\nURL: %v ", r.URL )
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8100", nil)
}
