package blueprint

import (
	"bytes"
	"net/http"
)

type Request struct {
	Name        string
	Description string
	Headers     http.Header
	Body        *bytes.Buffer

	// Todo:
	// Attributes
	// Schema
}

func NewRequest(name, description string, req *http.Request) (breq *Request, err error) {
	var body *bytes.Buffer

	body, err = copyBody(req.Body)
	if err != nil {
		return
	}

	return &Request{
		httpIO{
			Name:        name,
			Description: description,
			Header:      req.Header,
			ContentType: req.Header.Get("Content-Type"),
			Body:        body,
		},
		nil,
	}, nil
}
