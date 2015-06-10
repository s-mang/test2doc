package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Widget is a thing
type Widget struct {
	Id   int64
	Name string
	Role string
}

var allWidgets []Widget

func init() {
	allWidgets = []Widget{
		Widget{0, "Nothing", "N/A"},
		Widget{1, "Jibjab", "Instrument"},
		Widget{2, "Pencil", "Utensil"},
		Widget{3, "Fork", "Utensil"},
		Widget{4, "Password", "Credential"},
		Widget{5, "SpanFrankisco", "Home"},
		Widget{6, "Doc", "Villain"},
		Widget{7, "Coff3e", "Hack"},
	}
}

// HandleGetWidgets retrieves the collection of Wisdget
func HandleGetWidgets(w http.ResponseWriter, req *http.Request) {
	widgetsJSON, err := json.Marshal(allWidgets)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(widgetsJSON))
}

// HandleGetWidget retrieves a single Widget
func HandleGetWidget(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	if id >= int64(len(allWidgets)) {
		handleError(w, err, http.StatusNotFound)
		return
	}

	widgetJSON, err := json.Marshal(allWidgets[id])
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(widgetJSON))
}

// HandlePostWidget adds a Widget to the collection
func HandlePostWidget(w http.ResponseWriter, req *http.Request) {
	var widget Widget
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&widget)
	if err != nil {
		log.Println(1)
		handleError(w, err, http.StatusBadRequest)
		return
	}

	if len(widget.Name) == 0 {
		err = errors.New("Widget name can't be blank.")
		handleError(w, err, http.StatusBadRequest)
		return
	}

	// not thread safe...
	widget.Id = int64(len(allWidgets))
	allWidgets = append(allWidgets, widget)

	widgetJSON, err := json.Marshal(widget)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(widgetJSON))

}
