package doc

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testResponseBody = `{"foo": "bar"}`

func TestNewResponse_ResponseBodyIsCorrectlyCopied(t *testing.T) {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	w := httptest.NewRecorder()
	testResponseHandler(w, req)

	apiResp := NewResponse(w)

	if string(apiResp.Body.Content) != testResponseBody {
		t.Fatalf("expected 'string(apiResp.Body.Content)' (%v) to equal 'testResponseBody' (%v)", string(apiResp.Body.Content), testResponseBody)
	}
}

func TestNewResponse_OriginalResponseBodyDoesNotChange(t *testing.T) {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	w := httptest.NewRecorder()
	testResponseHandler(w, req)

	_ = NewResponse(w)

	if w.Body.String() != testRequestBody {
		t.Fatalf("expected 'w.Body.String()' (%v) to equal 'testRequestBody' (%v)", w.Body.String(), testRequestBody)
	}
}

func testResponseHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, testResponseBody)
}
