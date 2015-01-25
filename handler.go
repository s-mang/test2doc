package test2doc

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gophergala/test2doc/out"
)

func HandlerWrapper(handler http.HandlerFunc, docFile *out.APIDoc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := docFile.RecordRequest(r)
		if err != nil {
			log.Println(err.Error())
			return
		}

		resp, err := docFile.RecordResponse(r, handler)
		if err != nil {
			log.Println(err.Error())
			return
		}

		err = resp.Header().Write(w)
		if err != nil {
			log.Println(err.Error())
			return
		}

		w.WriteHeader(resp.Code)

		fmt.Fprint(w, resp.Body.String())
		return
	}
}
