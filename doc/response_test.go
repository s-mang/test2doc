package doc

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
)

var testResponseBody = `{"foo": "bar"}`

func (t *suite) TestNewResponse_ResponseBodyIsCorrectlyCopied() {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	t.Must(t.Nil(err))

	w := httptest.NewRecorder()
	testResponseHandler(w, req)

	apiResp := NewResponse(w)

	t.Equal(string(apiResp.Body.Content), testResponseBody)
}

func (t *suite) TestNewResponse_OriginalResponseBodyDoesNotChange() {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	t.Must(t.Nil(err))

	w := httptest.NewRecorder()
	testResponseHandler(w, req)

	_ = NewResponse(w)

	t.Equal(w.Body.String(), testRequestBody)
}

func testResponseHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, testResponseBody)
}
