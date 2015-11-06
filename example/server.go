package example

import (
	"log"
	"net/http"

	"github.com/adams-sarah/test2doc/example/foos"
	"github.com/adams-sarah/test2doc/example/widgets"
	"github.com/gorilla/mux"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/foos", foos.GetFoos).Methods("GET").Name("GetFoos")
	r.HandleFunc("/foos/{key}", foos.GetFoo).Methods("GET").Name("GetFoo")

	r.HandleFunc("/widgets", widgets.GetWidgets).Methods("GET").Name("GetWidgets")
	r.HandleFunc("/widgets", widgets.PostWidget).Methods("POST").Name("PostWidget")
	r.HandleFunc("/widgets/{id}", widgets.GetWidget).Methods("GET").Name("GetWidget")

	return r
}
