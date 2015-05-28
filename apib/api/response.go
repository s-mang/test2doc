package api

import (
	"net/http"
	"net/http/httptest"
)

type Response struct {
	StatusCode  int
	Description string
	Header      http.Header
	Body        []byte

	// TODO:
	// Attributes
	// Schema
}

func NewResponse(description string, w *httptest.ResponseRecorder) (*Response, error) {

	body1, body2, err := cloneBody(w.Body)
	if err != nil {
		return nil, err
	}

	w.Body = body1

	return &Response{
		StatusCode:  w.Code,
		Description: description,
		Header:      w.Header(),
		Body:        body2.Bytes(),
	}, nil
}

func (r *Response) ContentType() string {
	return r.Header.Get("Content-Type")
}

func (r *Response) BodyStr() string {
	return string(r.Body)
}
