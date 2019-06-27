
package main

import (
	app "go-playground/testing/14_app"
	"net/http"

)

func main() {
	http.ListenAndServe(":3000", &app.Server{})
}