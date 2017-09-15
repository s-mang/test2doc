package foos

import (
	"encoding/json"
	"net/http"
)

func (t *mainSuite) TestGetFoos() {
	urlPath, err := router.Get("GetFoos").URL()
	t.Must(t.Nil(err))

	resp, err := http.Get(server.URL + urlPath.String() + "?n=10")
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	fs := map[string]Foo{}
	err = decoder.Decode(&fs)
	t.Must(t.Nil(err))

	t.Equal(len(fs), len(AllFoos))

	for k, foo := range AllFoos {
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

	var foo Foo
	err = decoder.Decode(&foo)
	t.Must(t.Nil(err))

	t.Equal(foo.B, AllFoos[key].B)
	t.Equal(foo.A, AllFoos[key].A)
	t.Equal(foo.R, AllFoos[key].R)
}
