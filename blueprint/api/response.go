package blueprint

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

type Response struct {
	StatusCode  int
	Description string
	Headers     http.Header
	Body        *bytes.Buffer

	// Todo:
	// Attributes
	// Schema
}

func NewResponse(name, description string, w httptest.ResponseRecorder) (resp *Response, err error) {
	var body *bytes.Buffer

	body, err = copyBody(w.Body)
	if err != nil {
		return
	}

	return &Response{
		w.Code,
		httpIO{
			Name:        name,
			Description: description,
			Header:      w.Header(),
			ContentType: w.Header().Get("Content-Type"),
			Body:        body,
		},
	}, nil
}
