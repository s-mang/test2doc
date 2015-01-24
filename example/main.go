package main

import (
	"fmt"
	"log"
	"net/http"

	t "github.com/gophergala/test2doc"
)

func main() {
	http.Handle("/foo", t.HandlerWrapper(fooHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fancy that!")
}
