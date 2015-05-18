package api

import (
	"net/http"
)

type Request struct {
	Name        string
	Description string
	Header      http.Header
	Body        []byte
	Response    *Response
	// Responses []*Response

	// Todo:
	// Attributes
	// Schema
}

func NewRequest(name, description string, req *http.Request) (*Request, error) {
	body1, body2, err := cloneBody(req.Body)
	if err != nil {
		return nil, err
	}

	req.Body = nopCloser{body1}

	return &Request{
		Name:        name,
		Description: description,
		Header:      req.Header,
		Body:        body2.Bytes(),
	}, nil
}
