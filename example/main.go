package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	t "github.com/gophergala/test2doc"
)

var docFile *os.File

func init() {
	var err error
	docFile, err = os.Create("/Users/adamssarah/Desktop/apidoc.md")
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	http.Handle("/foo", t.HandlerWrapper(fooHandler, docFile))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fancy that!")
}
