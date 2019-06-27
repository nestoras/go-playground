package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var validToken = "123456"

func Apply(next http.Handler) http.HandlerFunc {
	return middleware(next.ServeHTTP)
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if token == validToken {
			next(w, r)
			return
		}
		http.Error(w, http.StatusText(401), 401)
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Welcome to Golang course")
}

//http://localhost:3000/?token=123456
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", middleware(home)).Methods("GET")
	http.ListenAndServe(":3000", r)
}
