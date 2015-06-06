package doc

import "strings"

// TODO
func (t *suite) TestRenderParameter() {

}

// TODO
func (t *suite) TestRenderParameter_ParameterIsNil() {

}

func (t *suite) TestQuoteParameterValue() {
	val := ParameterValue("param-value")

	quotedVal := val.Quote()
	t.Equal(len(val)+2, len(quotedVal))
	t.True(strings.Contains(quotedVal, string(val)))
	t.Equal(quotedVal[0], quotedVal[len(quotedVal)-1])
}

func (t *suite) TestQuoteParameterValue_EmptyValue() {
	var val ParameterValue

	quotedVal := val.Quote()
	t.Equal(string(val), quotedVal)
}
