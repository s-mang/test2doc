package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleInfo)
	mux.HandleFunc("/greet", HandleGreeting)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
