package handler

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Age        int
	Name       string
	Occupation string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	p := Person{
		Age:        31,
		Name:       "Nestoras",
		Occupation: "Software Engineer",
	}
	data, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}