package doc

import (
	"net/http"
	"testing"
)

func TestRenderHeader(t *testing.T) {

}

func TestRenderHeader_HeaderIsNil(t *testing.T) {

}

func TestContentType_OneContentType(t *testing.T) {
	ct := "text/plain"

	h := http.Header{}
	h.Add("Content-Type", ct)
	if NewHeader(h).ContentType != ct {
		t.Fatalf("expected 'NewHeader(h).ContentType' (%v) to equal 'ct' (%v)", NewHeader(h).ContentType, ct)
	}
}

func TestContentType_MultipleContentTypes_Mistakenly(t *testing.T) {
	ct := "text/plain"

	h := http.Header{}
	h.Add("Content-Type", ct)
	h.Add("Content-Type", "application/json")
	if NewHeader(h).ContentType != ct {
		t.Fatalf("expected 'NewHeader(h).ContentType' (%v) to equal 'ct' (%v)", NewHeader(h).ContentType, ct)
	}
}

func TestNewHeader_EmptyHeader(t *testing.T) {
	h := http.Header{}
	if NewHeader(h) != nil {
		t.Fatalf("expected 'NewHeader(h)' (%v) be nil", NewHeader(h))
	}
}
