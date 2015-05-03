package testrun

import (
	"net/http"
	"net/http/httptest"

	"github.com/adams-sarah/test2doc/blueprint"
)

// TODO: filter out 404 responses

func NewTestServer(handler http.Handler, apib *blueprint.APIBlueprint) *httptest.Server {
	return httptest.NewServer(handleAndRecord(handler, apib))
}

func handleAndRecord(handler http.Handler, apib *blueprint.APIBlueprint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// err := apib.RecordRequest(r)
		// if err != nil {
		// 	log.Println(err.Error())
		// 	return
		// }

		// resp, err := apib.RecordResponse(r, handler)
		// if err != nil {
		// 	log.Println(err.Error())
		// 	return
		// }

		// err = resp.Header().Write(w)
		// if err != nil {
		// 	log.Println(err.Error())
		// 	return
		// }

		// w.WriteHeader(resp.Code)

		// fmt.Fprint(w, resp.Body.String())
		return
	}
}
