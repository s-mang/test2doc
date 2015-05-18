package main

import (
	"fmt"
	"net/http"
)

// Foo is something cool
// Foo is a resource of the Something API
type Foo struct {
	B string
	A string
	R string
}

// HandleInfo serves basic server information to the client
// HandleInfo is an API action for the Foo resource
func HandleInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		handleNotFound(w)
		return
	}

	fmt.Fprintf(w, "TODO")
}

// HandleGreeting is an API action for the Foo resource
// This endpoint handles greetings from a client
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
