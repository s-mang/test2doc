package main

import (
	"log"
	"net/http"

	"github.com/adams-sarah/test2doc/example/foos"
	"github.com/adams-sarah/test2doc/example/widgets"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	foos.AddRoutes(router)
	widgets.AddRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
