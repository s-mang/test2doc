package widgets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (t *mainSuite) TestGetWidgets() {
	urlPath, err := router.Get("GetWidgets").URL()
	t.Must(t.Nil(err))

	resp, err := http.Get(server.URL + urlPath.String())
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var ws []Widget
	err = decoder.Decode(&ws)
	t.Must(t.Nil(err))

	t.Equal(len(ws), len(AllWidgets))
	t.Must(t.True(len(ws) > 2))

	t.Equal(ws[0].Id, AllWidgets[0].Id)
	t.Equal(ws[2].Name, AllWidgets[2].Name)
	t.Equal(ws[1].Role, AllWidgets[1].Role)
}

func (t *mainSuite) TestGetWidgetBadRequest() {
	idStr := "hello"

	urlPath, err := router.Get("GetWidget").URL("id", idStr)
	t.Must(t.Nil(err))

	resp, err := http.Get(server.URL + urlPath.String())
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusBadRequest))
}

func (t *mainSuite) TestGetWidget() {
	var id int64 = 2
	idStr := fmt.Sprintf("%d", id)

	urlPath, err := router.Get("GetWidget").URL("id", idStr)
	t.Must(t.Nil(err))

	resp, err := http.Get(server.URL + urlPath.String())
	t.Must(t.Nil(err))

	t.Must(t.Equal(resp.StatusCode, http.StatusOK))

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var widget Widget
	err = decoder.Decode(&widget)
	t.Must(t.Nil(err))

	t.Equal(widget.Id, AllWidgets[2].Id)
	t.Equal(widget.Name, AllWidgets[2].Name)
	t.Equal(widget.Role, AllWidgets[2].Role)
}

func (t *mainSuite) TestPostWidget() {
	urlPath, err := router.Get("PostWidget").URL()
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
