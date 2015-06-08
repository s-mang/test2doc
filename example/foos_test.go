package main

import "net/http"

func (t *mainSuite) TestHandleGetFoos() {
	resp, err := http.Get(server.URL + "/foos")
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	// b, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// t.Must(t.Nil(err))

	// TODO: finish test
}

func (t *mainSuite) TestHandleGetFoo() {
	resp, err := http.Get(server.URL + "/foos/1")
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	// b, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// t.Must(t.Nil(err))
	// TODO: finish test
}
