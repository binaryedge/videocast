package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello " + GetData()))
	})
	r.Handle("/bar", barHandler{}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func GetData() string {
	return "joao"
}

type barHandler struct {
}

func (b barHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + GetData()))
}
