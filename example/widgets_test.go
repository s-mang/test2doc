package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (t *mainSuite) TestHandleGetWidgets() {
	resp, err := http.Get(server.URL + "/widgets")
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	// b, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// t.Must(t.Nil(err))
	// TODO: finish test
}

func (t *mainSuite) TestHandlePostWidget() {
	widget := Widget{
		Name: "anotherwidget",
		Role: "controller",
	}

	jsonb, err := json.Marshal(widget)
	t.Must(t.Nil(err))
	buf := bytes.NewBuffer(jsonb)

	resp, err := http.Post(server.URL+"/widgets", "application/json", buf)
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	// b, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// t.Must(t.Nil(err))
	// TODO: finish test
}

func (t *mainSuite) TestHandleGetWidget() {
	resp, err := http.Get(server.URL + "/widgets/mywidget")
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	// b, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// t.Must(t.Nil(err))
	// TODO: finish test
}
