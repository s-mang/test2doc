package doc

import (
	"strings"
	"testing"
)

// TODO
func TestRenderParameter(t *testing.T) {

}

// TODO
func TestRenderParameter_ParameterIsNil(t *testing.T) {

}

func TestQuoteParameterValue(t *testing.T) {
	val := ParameterValue("param-value")

	quotedVal := val.Quote()
	if len(val)+2 != len(quotedVal) {
		t.Fatalf("expected 'len(val)+2' (%v) to equal 'len(quotedVal)' (%v)", len(val)+2, len(quotedVal))
	}
	if !strings.Contains(quotedVal, string(val)) {
		t.Fatalf("expected 'strings.Contains(quotedVal, string(val))' (%v) be true", strings.Contains(quotedVal, string(val)))
	}
	if quotedVal[0] != quotedVal[len(quotedVal)-1] {
		t.Fatalf("expected 'quotedVal[0]' (%v) to equal 'quotedVal[len(quotedVal)-1]' (%v)", quotedVal[0], quotedVal[len(quotedVal)-1])
	}
}

func TestQuoteParameterValue_EmptyValue(t *testing.T) {
	var val ParameterValue

	quotedVal := val.Quote()
	if string(val) != quotedVal {
		t.Fatalf("expected 'string(val)' (%v) to equal 'quotedVal' (%v)", string(val), quotedVal)
	}
}

func TestQuoteParameterName(t *testing.T) {
	name := ParameterName("param-name")

	quotedName := name.Quote()
	if len(name)+2 != len(quotedName) {
		t.Fatalf("expected 'len(name)+2' (%v) to equal 'len(quotedName)' (%v)", len(name)+2, len(quotedName))
	}
	if !strings.Contains(quotedName, string(name)) {
		t.Fatalf("expected 'strings.Contains(quotedName, string(name))' (%v) be true", strings.Contains(quotedName, string(name)))
	}
	if quotedName[0] != quotedName[len(quotedName)-1] {
		t.Fatalf("expected 'quotedName[0]' (%v) to equal 'quotedName[len(quotedName)-1]' (%v)", quotedName[0], quotedName[len(quotedName)-1])
	}
}
