package blueprint

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
)

type Action struct {
	HTTPMethod string
	URI        *url.URL
	Parameters []*Parameter
	Requests   []*Request

	// Todo:
	// - Relation
	// - Attributes
	// - Schema
}

func NewAction(uri *url.URL, httpMethod string) *Action {
	return &Action{
		HTTPMethod: httpMethod,
		URI:        uri,
	}
}

type httpIO struct {
	Name        string
	Description string
	Header      http.Header
	ContentType string
	Body        *bytes.Buffer

	// Todo:
	// - Schema
	// - Attributes
}

type Request struct {
	httpIO
	Responses []*Response
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

type Response struct {
	StatusCode int
	httpIO
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

func copyBody(r io.Reader) (buf *bytes.Buffer, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return buf, err
	}

	buf = bytes.NewBuffer(body)
	r = bytes.NewBuffer(body)

	return
}
