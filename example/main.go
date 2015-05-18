package main

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
