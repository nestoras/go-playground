package main

import (
	"fmt"
	"net/http"
)

// https://golang.org/pkg/net/http/#Header
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "Welcome to Golang course")
	} else if r.URL.Path == "/admin" {
		fmt.Fprint(w, "Welcome to admin")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "custom 404")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Your webserver is running")
	http.ListenAndServe(":3000", nil)
}
