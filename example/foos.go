package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Foo is something cool
type Foo struct {
	B string
	A string
	R string
}

var allFoos map[string]Foo

func init() {
	allFoos = map[string]Foo{ // map[key]Foo
		"ABeeSee":            Foo{"A", "Bee", "See"},
		"OneTwoThree":        Foo{"One", "Two", "Three"},
		"SomethingFunForYou": Foo{"Something", "Fun", "ForYou"},
		"2":                  Foo{"", "", "2"},
	}
}

// HandleGetFoos retrieves the collection of Foos
func HandleGetFoos(w http.ResponseWriter, req *http.Request) {
	foosJSON, err := json.Marshal(allFoos)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(foosJSON))
}

// HandleGetFoo retrieves a single Foo
func HandleGetFoo(w http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["key"]
	foo, ok := allFoos[key]
	if !ok {
		err := errors.New("No Foo found.")
		handleError(w, err, http.StatusNotFound)
		return
	}

	fooJSON, err := json.Marshal(foo)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(fooJSON))
}
