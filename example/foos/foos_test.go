package foos_test

import (
	"encoding/json"
	"net/http"

	"github.com/adams-sarah/test2doc/example/foos"
)

func (t *mainSuite) TestGetFoos() {
	urlPath, err := router.Get("GetFoos").URL()
	t.Must(t.Nil(err))

	resp, err := http.Get(server.URL + urlPath.String() + "?n=10")
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	fs := map[string]foos.Foo{}
	err = decoder.Decode(&fs)
	t.Must(t.Nil(err))

	t.Equal(len(fs), len(foos.AllFoos))

	for k, foo := range foos.AllFoos {
		t.Equal(fs[k].B, foo.B)
		t.Equal(fs[k].A, foo.A)
		t.Equal(fs[k].R, foo.R)
	}
}

func (t *mainSuite) TestGetFoo() {
	key := "ABeeSee"
	urlPath, err := router.Get("GetFoo").URL("key", key)
	t.Must(t.Nil(err))

	resp, err := http.Get(server.URL + urlPath.String())
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var foo foos.Foo
	err = decoder.Decode(&foo)
	t.Must(t.Nil(err))

	t.Equal(foo.B, foos.AllFoos[key].B)
	t.Equal(foo.A, foos.AllFoos[key].A)
	t.Equal(foo.R, foos.AllFoos[key].R)
}
