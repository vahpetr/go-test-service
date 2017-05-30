package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":8080"
)

var calls = 0

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	// required noop
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	calls++
	fmt.Fprintf(w, "Hello, world! You have called me %d times.\n", calls)
}

func main() {
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		panic("ListenAndServe: " + err.Error())
	} else {
		fmt.Printf("Started server at http://localhost%v.\n", port)
	}
}
