package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/happyreturns/test2doc/example/foos"
	"github.com/happyreturns/test2doc/example/widgets"
)

func main() {
	router := mux.NewRouter()
	foos.AddRoutes(router)
	widgets.AddRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
