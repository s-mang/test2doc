package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	t "github.com/gophergala/test2doc"
	"github.com/gophergala/test2doc/out"
)

const docFileName = "example_doc.md"

var (
	docFile *out.APIDoc
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	docFilePath := fmt.Sprintf("%s/%s", wd, docFileName)

	desc := &out.APIDescription{
		APIBlueprintFormat: out.APIBlueprintFormat,
		HostURL:            "http://www.google.com",
		Title:              "Google Search",
		Description:        "The only search engine.",
	}

	docFile = out.NewAPIDoc(docFilePath, desc)
}

func main() {
	http.Handle("/foo", t.HandlerWrapper(fooHandler, docFile))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"Message": "Fancy that!"}`)
}
