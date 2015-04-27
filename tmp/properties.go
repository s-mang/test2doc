package testdoc

import (
	"fmt"
	"strings"
)

type PropertySet interface {
	List() string
}

type RequestProperties struct {
	Description string
}

func (p *RequestProperties) List() string {
	return p.Description
}

type ResponseProperties struct {
	StatusCode  int
	ContentType string
}

func (p *ResponseProperties) List() string {
	fmtdAttrs := []string{
		fmt.Sprintf("%d", p.StatusCode),
	}

	if len(p.ContentType) > 0 {
		fmtdAttrs = append(
			fmtdAttrs,
			fmt.Sprintf("(%s)", p.ContentType),
		)
	}

	return strings.Join(fmtdAttrs, " ")
}
