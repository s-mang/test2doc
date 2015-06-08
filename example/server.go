package main

// JSON Placeholder API
// Fake Online REST API for Testing and Prototyping
// http://jsonplaceholder.typicode.com

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := newRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleGetInfo).Methods("GET")

	r.HandleFunc("/foos", HandleGetFoos).Methods("GET")
	r.HandleFunc("/foos/{key}", HandleGetFoo).Methods("GET")

	r.HandleFunc("/widgets", HandleGetWidgets).Methods("GET")
	r.HandleFunc("/widgets", HandlePostWidget).Methods("POST")
	r.HandleFunc("/widgets/{id}", HandleGetWidget).Methods("GET")

	return r
}

// HandleGetInfo serves basic info about the Server resource to
// the client
func HandleGetInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		HandleNotFound(w)
		return
	}

	fmt.Fprintf(w, "TODO")
}

// HandleNotFound is the basic 404 handler
func HandleNotFound(w http.ResponseWriter) {
	http.Error(w, "Not found", http.StatusNotFound)
}
