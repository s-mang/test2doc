package api

import (
	"bytes"
	"io"
	"io/ioutil"
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

func cloneBody(r io.Reader) (*bytes.Buffer, *bytes.Buffer, error) {
	var clone1, clone2 bytes.Buffer

	rBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return &clone1, &clone2, err
	}

	mw := io.MultiWriter(&clone1, &clone2)
	_, err = mw.Write(rBytes)

	return &clone1, &clone2, err
}
