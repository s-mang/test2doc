package main

import (
	"fmt"
	"net/http"
)

func repoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{
		"Id": 1,
		"Name": "my_repo",
		"Owner": "adams-sarah"
	}`)
}

func reposHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `[
		{
			"Id": 1,
			"Name": "my_repo",
			"Owner": "adams-sarah"
		},
		{
			"Id": 2,
			"Name": "my_other_repo",
			"Owner": "adams-sarah"
		},
		{
			"Id": 3,
			"Name": "someone_elses_repo",
			"Owner": "goody-twoshoes"
		}
	]`)
}
