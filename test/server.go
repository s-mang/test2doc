package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/adams-sarah/test2doc/doc"
)

// TODO: filter out 404 responses
func NewServer(handler http.Handler, outDir string) (s *httptest.Server, err error) {
	outDoc, err := doc.NewDoc(outDir)
	if err != nil {
		return s, err
	}

	return httptest.NewServer(handleAndRecord(handler, outDoc)), nil
}

func handleAndRecord(handler http.Handler, outDoc *doc.Doc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// err := doc.RecordRequest(outDoc, r)
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
