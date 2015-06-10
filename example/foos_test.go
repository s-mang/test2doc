package main

import (
	"encoding/json"
	"net/http"
)

func (t *mainSuite) TestHandleGetFoos() {
	urlPath, err := router.Get("HandleGetFoos").URL()
	t.Must(t.Nil(err))

	resp, err := http.Get(server.URL + urlPath.String())
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	foos := map[string]Foo{}
	err = decoder.Decode(&foos)
	t.Must(t.Nil(err))

	t.Equal(len(foos), len(allFoos))

	for k, foo := range allFoos {
		t.Equal(foos[k].B, foo.B)
		t.Equal(foos[k].A, foo.A)
		t.Equal(foos[k].R, foo.R)
	}
}

func (t *mainSuite) TestHandleGetFoo() {
	key := "ABeeSee"
	urlPath, err := router.Get("HandleGetFoo").URL("key", key)
	t.Must(t.Nil(err))

	resp, err := http.Get(server.URL + urlPath.String())
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var foo Foo
	err = decoder.Decode(&foo)
	t.Must(t.Nil(err))

	t.Equal(foo.B, allFoos[key].B)
	t.Equal(foo.A, allFoos[key].A)
	t.Equal(foo.R, allFoos[key].R)
}
