package test2doc

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func HandlerWrapper(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			body, err := payload(r)
			if err != nil {
				panic(err.Error())
			}

			log.Println("Header: ", r.Header)
			log.Println("Body: ", string(body))

			handler(w, r)
		},
	)

}

func payload(r *http.Request) (body []byte, err error) {
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}
