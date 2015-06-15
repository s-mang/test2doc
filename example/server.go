// JSON Placeholder API
// Fake Online REST API for Testing and Prototyping
// http://jsonplaceholder.typicode.com
package main

import (
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
	r.HandleFunc("/foos", HandleGetFoos).Methods("GET").Name("HandleGetFoos")
	r.HandleFunc("/foos/{key}", HandleGetFoo).Methods("GET").Name("HandleGetFoo")

	r.HandleFunc("/widgets", HandleGetWidgets).Methods("GET").Name("HandleGetWidgets")
	r.HandleFunc("/widgets", HandlePostWidget).Methods("POST").Name("HandlePostWidget")
	r.HandleFunc("/widgets/{id}", HandleGetWidget).Methods("GET").Name("HandleGetWidget")

	return r
}

// handleError serves an error response to the client
func handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Printf("Error %d: %v\n", statusCode, err)
	http.Error(w, http.StatusText(statusCode), statusCode)
}
