package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (t *mainSuite) TestHandleInfo() {
	resp, err := http.Get(server.URL + InfoPath)
	t.Must(t.Nil(err))

	info, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	t.Must(t.Nil(err))
	t.Equal(string(info), "TODO")
}

func (t *mainSuite) TestHandleGreetingWithoutName() {
	resp, err := http.PostForm(server.URL+GreetingPath, nil)
	t.Must(t.Nil(err))

	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	t.Must(t.Nil(err))
	t.Equal(string(greeting), "Hello.")
}

func (t *mainSuite) TestHandleGreetingWithName() {
	name := "Sarah"

	v := url.Values{}
	v.Set("name", name)
	u := server.URL + GreetingPath + "?" + v.Encode()

	resp, err := http.PostForm(u, nil)
	t.Must(t.Nil(err))

	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	t.Must(t.Nil(err))
	t.Equal(string(greeting), fmt.Sprintf("Hello, %s.", name))
}
