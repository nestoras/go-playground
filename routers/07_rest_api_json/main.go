package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Error is used to render API response
type Error struct {
	Message string `json:"message"`
}

// Data is the top level structure that views expect data
// to come in.
type Data struct {
	Error  *Error      `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

// Render is used to render the view with the predefined layout.
func Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	var vd Data
	switch d := data.(type) {
	case Data:
		vd = d
	default:
		vd = Data{
			Result: data,
		}
	}

	response, err := json.Marshal(vd)
	if err != nil {
		panic(err)
	}
	w.Write(response)
}

type Photo struct {
	ID       uint
	Filename string
	Path     string
}

func home(w http.ResponseWriter, r *http.Request) {
	photo := Photo{
		ID:       1,
		Filename: "my-photo.jpg",
		Path:     "/users/",
	}

	Render(w, r, photo)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	http.ListenAndServe(":3000", r)
}
