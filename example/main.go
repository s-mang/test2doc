package main

// JSON Placeholder API
// Fake Online REST API for Testing and Prototyping
// http://jsonplaceholder.typicode.com

import (
	"log"
	"net/http"
)

const (
	InfoPath     = "/"
	GreetingPath = "/greet"
)

func main() {
	mux := newMux()
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(InfoPath, HandleInfo)
	mux.HandleFunc(GreetingPath, HandleGreeting)
	return mux
}
