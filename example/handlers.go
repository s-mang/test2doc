package main

import (
	"fmt"
	"net/http"
)

// Foo is something cool
type Foo struct {
	B string
	A string
	R string
}

// HandleInfo serves basic server information to the client
func HandleInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		handleNotFound(w)
		return
	}

	fmt.Fprintf(w, "TODO")
}

// HandleGreeting is an API action for the Foo resource
func HandleGreeting(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		handleNotFound(w)
		return
	}

	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprintf(w, "Hello.")
	} else {
		fmt.Fprintf(w, "Hello, %s.", name)
	}

	return

}

func handleNotFound(w http.ResponseWriter) {
	http.Error(w, "Not found", http.StatusNotFound)
}
