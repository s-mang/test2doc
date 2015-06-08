package main

import "net/http"

// Widget is a thing
type Widget struct {
	Name string
	Role string
}

// HandleGetWidgets retrieves the collection of Wisdget
func HandleGetWidgets(w http.ResponseWriter, r *http.Request) {

}

// HandleGetWidget retrieves a single Widget
func HandleGetWidget(w http.ResponseWriter, r *http.Request) {

}

// HandlePostWidget adds a Widget to the collection
func HandlePostWidget(w http.ResponseWriter, r *http.Request) {

}
