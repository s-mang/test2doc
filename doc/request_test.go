package doc

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

	apiReq, err := NewRequest(req)
	t.Must(t.Nil(err))

	t.Equal(string(apiReq.Body.Content), testRequestBody)
}

func (t *suite) TestNewRequest_OriginalRequestBodyDoesNotChange() {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	t.Must(t.Nil(err))

	_, err = NewRequest(req)
	t.Must(t.Nil(err))

	httpReqBody, err := ioutil.ReadAll(req.Body)
	t.Must(t.Nil(err))
	t.Equal(string(httpReqBody), testRequestBody)
}

// TODO
func (t *suite) TestNewRequest_404DontRecord() {

}

// TODO
func (t *suite) TestNewRequest_PanicMidRequest() {

}
