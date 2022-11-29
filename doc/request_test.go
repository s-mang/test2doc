package doc

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

var testRequestBody = `{"foo": "bar"}`

func TestNewRequest_RequestBodyIsCorrectlyCopied(t *testing.T) {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	apiReq, err := NewRequest(req)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	if string(apiReq.Body.Content) != testRequestBody {
		t.Fatalf("expected 'string(apiReq.Body.Content)' (%v) to equal 'testRequestBody' (%v)", string(apiReq.Body.Content), testRequestBody)
	}
}

func TestNewRequest_OriginalRequestBodyDoesNotChange(t *testing.T) {
	body := bytes.NewBuffer([]byte(testRequestBody))
	req, err := http.NewRequest("POST", "http://httpbin.org/post", body)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	_, err = NewRequest(req)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	httpReqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}
	if string(httpReqBody) != testRequestBody {
		t.Fatalf("expected 'string(httpReqBody)' (%v) to equal 'testRequestBody' (%v)", string(httpReqBody), testRequestBody)
	}
}

// TODO
func TestNewRequest_404DontRecord(t *testing.T) {

}

// TODO
func TestNewRequest_PanicMidRequest(t *testing.T) {

}
