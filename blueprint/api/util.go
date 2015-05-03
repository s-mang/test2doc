package blueprint

import (
	"bytes"
	"io"
	"io/ioutil"
)

func copyBody(r io.Reader) (buf *bytes.Buffer, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return buf, err
	}

	buf = bytes.NewBuffer(body)
	r = bytes.NewBuffer(body)

	return
}
