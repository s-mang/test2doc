package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

var testRequestBody = `{"foo": "bar"}`

func (t *suite) TestNewRequest_RequestBodyIsCorrectlyCopied() {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	t.Must(t.Nil(err))

	apiReq, err := NewRequest("Example POST", "POST some example JSON", req)
	t.Must(t.Nil(err))

	t.Equal(string(apiReq.Body), testRequestBody)
}

func (t *suite) TestNewRequest_OriginalRequestBodyDoesNotChange() {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	t.Must(t.Nil(err))

	_, err = NewRequest("Example POST", "POST some example JSON", req)
	t.Must(t.Nil(err))

	httpReqBody, err := ioutil.ReadAll(req.Body)
	t.Must(t.Nil(err))
	t.Equal(string(httpReqBody), testRequestBody)
}
