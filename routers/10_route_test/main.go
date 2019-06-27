package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
	ID       uint
	Username string
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	user := User{
		ID:       1,
		Username: "nestoras",
	}

	response, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Write(response)
}

//go test routers/10_route_test/main_test.go routers/10_route_test/normalizer.go
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", UserHandler)
	http.ListenAndServe(":3000", r)
}
