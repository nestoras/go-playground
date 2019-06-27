package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "Welcome to Golang course")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "custom 404")
	}
}

func handlerFuncAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Welcome to admin")

}

func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc)
	mux.HandleFunc("/admin", handlerFuncAdmin)
	fmt.Println("Your webserver is running")
	http.ListenAndServe(":3000", mux)
}
