package main

import (
	"fmt"
	"net/http"
)

// Retrieve server info
func HandleInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		handleNotFound(w)
		return
	}

	fmt.Fprintf(w, "TODO")
}

// Greet the server
func HandleGreeting(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		handleNotFound(w)
		return
	}

	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprintf(w, "Thanks")
	} else {
		fmt.Fprintf(w, "Thanks, "+name)
	}

	return

}

func handleNotFound(w http.ResponseWriter) {
	http.Error(w, "Not found", http.StatusNotFound)
}
