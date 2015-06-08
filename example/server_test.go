package main

import (
	"io/ioutil"
	"net/http"
)

func (t *mainSuite) TestHandleGetInfo() {
	resp, err := http.Get(server.URL)
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	info, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	t.Must(t.Nil(err))

	t.Equal(string(info), "TODO")
}
