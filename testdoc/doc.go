package testdoc

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
)

type APIDoc struct {
	file *os.File
	desc *APIDescription
}

func NewAPIDoc(filePath string, desc *APIDescription) *APIDoc {
	os.Remove(filePath) // if there is an error, we don't really care -- file likely does not exist

	file, err := os.Create(filePath)
	if err != nil {
		panic(err.Error())
	}

	doc := &APIDoc{
		file: file,
		desc: desc,
	}

	err = doc.WriteDescription()
	if err != nil {
		panic(err.Error())
	}

	return doc
}

func (doc *APIDoc) RecordRequest(r *http.Request) (err error) {
	body, err := payload(r)
	if err != nil {
		return err
	}

	err = doc.WriteRequestTitle("")
	if err != nil {
		return err
	}

	err = doc.WriteHeaders(r.Header)
	if err != nil {
		return err
	}

	return doc.WriteBody(string(body))
}

func (doc *APIDoc) RecordResponse(r *http.Request, handler http.Handler) (resp *httptest.ResponseRecorder, err error) {
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, r)

	err = doc.WriteResponseTitle(resp.Code, resp.Header().Get("Content-Type"))
	if err != nil {
		return
	}

	err = doc.WriteHeaders(resp.Header())
	if err != nil {
		return
	}

	err = doc.WriteBody(string(resp.Body.String()))
	return
}

func payload(r *http.Request) (body []byte, err error) {
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}
