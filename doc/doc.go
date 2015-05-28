package doc

import (
	"net/http"
	"net/http/httptest"

	"github.com/adams-sarah/test2doc/apib"
)

var doc *apib.APIBlueprint

// TODO: filter out 404 responses
func NewTestServer(handler http.Handler, outDir string) (s *httptest.Server, err error) {
	doc, err = apib.NewAPIBlueprint(outDir)
	if err != nil {
		return
	}

	return httptest.NewServer(handleAndRecord(handler, doc)), nil
}

func handleAndRecord(handler http.Handler, doc *apib.APIBlueprint) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// err := doc.RecordRequest(r)
		// if err != nil {
		// 	log.Println(err.Error())
		// 	return
		// }

		// resp, err := doc.RecordResponse(r, handler)
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

		// TODO: remove
		handler.ServeHTTP(w, req)
	}
}
