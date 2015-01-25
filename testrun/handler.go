package testrun

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gophergala/test2doc/testdoc"
)

func NewTestServer(handler http.Handler, docFile *testdoc.APIDoc) *httptest.Server {
	return httptest.NewServer(handleAndRecord(handler, docFile))
}

func handleAndRecord(handler http.Handler, docFile *testdoc.APIDoc) http.HandlerFunc {
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
