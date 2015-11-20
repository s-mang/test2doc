package widgets

import (
	"encoding/json"
	"errors"
	"fmt"
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

var AllWidgets []Widget

func init() {
	AllWidgets = []Widget{
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

// GetWidgets retrieves the collection of Wisdget
func GetWidgets(w http.ResponseWriter, req *http.Request) {
	widgetsJSON, err := json.Marshal(AllWidgets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(widgetsJSON))
}

// GetWidget retrieves a single Widget
func GetWidget(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id >= int64(len(AllWidgets)) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	widgetJSON, err := json.Marshal(AllWidgets[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(widgetJSON))
}

// PostWidget adds a Widget to the collection
func PostWidget(w http.ResponseWriter, req *http.Request) {
	var widget Widget
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&widget)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(widget.Name) == 0 {
		err = errors.New("Widget name can't be blank.")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// not thread safe...
	widget.Id = int64(len(AllWidgets))
	AllWidgets = append(AllWidgets, widget)

	widgetJSON, err := json.Marshal(widget)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(widgetJSON))

}
