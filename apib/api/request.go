package api

import (
	"net/http"
)

type Request struct {
	Header http.Header
	Body   []byte

	// TODO:
	// Attributes
	// Schema
}

func NewRequest(req *http.Request) (*Request, error) {
	body1, body2, err := cloneBody(req.Body)
	if err != nil {
		return nil, err
	}

	req.Body = nopCloser{body1}

	return &Request{
		Header: req.Header,
		Body:   body2.Bytes(),
	}, nil
}

func (r *Request) ContentType() string {
	return r.Header.Get("Content-Type")
}

func (r *Request) BodyStr() (body string) {
	body = string(r.Body)
	if r.ContentType() == "application/json" {
		var err error
		body, err = indentJSONBody(string(r.Body))
		if err != nil {
			panic(err.Error())
		}
	}

	return

}
