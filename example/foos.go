package main

import "net/http"

// Foo is something cool
type Foo struct {
	B string
	A string
	R string
}

// HandleGetFoos retrieves the collection of Foos
func HandleGetFoos(w http.ResponseWriter, r *http.Request) {

}

// HandleGetFoo retrieves a single Foo
func HandleGetFoo(w http.ResponseWriter, r *http.Request) {

}
