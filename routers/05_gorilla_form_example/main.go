package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	return parseValues(r.PostForm, dst)
}

func parseURLParams(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	return parseValues(r.Form, dst)
}

func parseValues(values url.Values, dst interface{}) error {
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	if err := dec.Decode(dst, values); err != nil {
		return err
	}
	return nil
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type Data struct {
	Error  *Error      `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

type Error struct {
	Message string `json:"message"`
}

type User struct {
	Email    string
	Password string
}

// http://localhost:3000/?email=nestoras&password=123456
func LoginGet(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseURLParams(r, &form); err != nil {
		return
	}

	vd := Data{
		Result: &form,
	}

	response, err := json.Marshal(vd)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		return
	}

	user := User{
		Email:    form.Email,
		Password: form.Password,
	}

	vd := Data{
		Result: user,
	}
	response, err := json.Marshal(vd)

	if err != nil || user.Password == "" {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Write(response)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", LoginGet).Methods("GET")
	r.HandleFunc("/", Login).Methods("POST")
	http.ListenAndServe(":3000", r)
}
