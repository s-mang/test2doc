package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func gistCompareHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	if len(body) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{
			"Status": "Bad Request",
			"Reason": "No request body"
		}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"Diffs": [
			{"Title": ["MyGist","Some other Gist"]}, 
			{"Id": [1,2]}
		]
	}`)
}

func gistHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{
		"Id": 1,
		"Title": "MyGist"
	}`)
}

func gistsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `[
		{
			"Id": 1,
			"Title": "MyGist"
		},
		{
			"Id": 2,
			"Title": "Some other Gist"
		}
	]`)
}
