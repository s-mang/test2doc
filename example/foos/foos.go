package foos

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

var AllFoos map[string]Foo

func init() {
	AllFoos = map[string]Foo{ // map[key]Foo
		"ABeeSee":            Foo{"A", "Bee", "See"},
		"OneTwoThree":        Foo{"One", "Two", "Three"},
		"SomethingFunForYou": Foo{"Something", "Fun", "ForYou"},
		"2":                  Foo{"", "", "2"},
	}
}

// GetFoos retrieves the collection of Foos
func GetFoos(w http.ResponseWriter, req *http.Request) {
	foosJSON, err := json.Marshal(AllFoos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(foosJSON))
}

// GetFoo retrieves a single Foo
func GetFoo(w http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["key"]
	foo, ok := AllFoos[key]
	if !ok {
		err := errors.New("No Foo found.")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fooJSON, err := json.Marshal(foo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(fooJSON))
}
