package main

import (
	"log"
	"net/http"
)

// APIDOC(Name): Some API
// APIDOC(Description): A nice synopsis of 'Some API'
// APIDOC(Host): http://httpbin.org

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleInfo)
	mux.HandleFunc("/greet", HandleGreeting)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
