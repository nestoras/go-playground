package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// http://localhost:3000/product/iphone/category/phones
	r.HandleFunc("/product/{product}/category/{category}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		product := vars["product"]
		category := vars["category"]

		fmt.Fprintf(w, "You've requested the product: %s from the category %s\n", product, category)
	})

	//http://localhost:3000/customer/1/
	r.HandleFunc("/customer/{id:[0-9]+}/", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		fmt.Fprintf(w, "Customer ID: %s\n", id)
	})

	http.ListenAndServe(":3000", r)
}
