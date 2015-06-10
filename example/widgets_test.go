package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (t *mainSuite) TestHandleGetWidgets() {
	urlPath, err := router.Get("HandleGetWidgets").URL()
	t.Must(t.Nil(err))

	resp, err := http.Get(server.URL + urlPath.String())
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var widgets []Widget
	err = decoder.Decode(&widgets)
	t.Must(t.Nil(err))

	t.Equal(len(widgets), len(allWidgets))
	t.Must(t.True(len(widgets) > 2))

	t.Equal(widgets[0].Id, allWidgets[0].Id)
	t.Equal(widgets[2].Name, allWidgets[2].Name)
	t.Equal(widgets[1].Role, allWidgets[1].Role)
}

func (t *mainSuite) TestHandleGetWidget() {
	var id int64 = 2
	idStr := fmt.Sprintf("%d", id)

	urlPath, err := router.Get("HandleGetWidget").URL("id", idStr)
	t.Must(t.Nil(err))

	resp, err := http.Get(server.URL + urlPath.String())
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var widget Widget
	err = decoder.Decode(&widget)
	t.Must(t.Nil(err))

	t.Equal(widget.Id, allWidgets[2].Id)
	t.Equal(widget.Name, allWidgets[2].Name)
	t.Equal(widget.Role, allWidgets[2].Role)
}

func (t *mainSuite) TestHandlePostWidget() {
	urlPath, err := router.Get("HandlePostWidget").URL()
	t.Must(t.Nil(err))

	widget := Widget{
		Name: "anotherwidget",
		Role: "controller",
	}

	jsonb, err := json.Marshal(widget)
	t.Must(t.Nil(err))
	buf := bytes.NewBuffer(jsonb)

	resp, err := http.Post(server.URL+urlPath.String(), "application/json", buf)
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusCreated))

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var respWidget Widget
	err = decoder.Decode(&respWidget)
	t.Must(t.Nil(err))

	t.True(respWidget.Id > 0)
	t.Equal(respWidget.Name, widget.Name)
	t.Equal(respWidget.Role, widget.Role)
}
