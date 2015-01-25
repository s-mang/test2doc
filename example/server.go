package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	addRoutes(mux)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/gists/1/compare", gistCompareHandler)
	mux.HandleFunc("/gists/1", gistHandler)
	mux.HandleFunc("/gists", gistsHandler)

	mux.HandleFunc("/repos/1", repoHandler)
	mux.HandleFunc("/repos", reposHandler)
}
