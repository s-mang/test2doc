package test2doc

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gophergala/test2doc/out"
)

func HandlerWrapper(handler http.HandlerFunc, docFile *os.File) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := payload(r)
		if err != nil {
			panic(err.Error())
		}

		err = out.WriteHeaders(docFile, r.Header)
		if err != nil {
			log.Println(err.Error())
			return
		}

		err = out.WriteBody(docFile, string(body))
		if err != nil {
			log.Println(err.Error())
			return
		}

		handler(w, r)
	}
}

func payload(r *http.Request) (body []byte, err error) {
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}
