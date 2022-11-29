package widgets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestGetWidgets(t *testing.T) {
	urlPath, err := router.Get("GetWidgets").URL()
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	resp, err := http.Get(server.URL + urlPath.String())
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 'resp.StatusCode' (%v) to equal 'http.StatusOK' (%v)", resp.StatusCode, http.StatusOK)
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var ws []Widget
	err = decoder.Decode(&ws)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	if len(ws) != len(AllWidgets) {
		t.Fatalf("expected 'len(ws)' (%v) to equal 'len(AllWidgets)' (%v)", len(ws), len(AllWidgets))
	}
	if len(ws) <= 2 {
		t.Fatalf("expected 'len(ws) > 2' (%v) be true", len(ws) > 2)
	}

	if ws[0].Id != AllWidgets[0].Id {
		t.Fatalf("expected 'ws[0].Id' (%v) to equal 'AllWidgets[0].Id' (%v)", ws[0].Id, AllWidgets[0].Id)
	}
	if ws[2].Name != AllWidgets[2].Name {
		t.Fatalf("expected 'ws[2].Name' (%v) to equal 'AllWidgets[2].Name' (%v)", ws[2].Name, AllWidgets[2].Name)
	}
	if ws[1].Role != AllWidgets[1].Role {
		t.Fatalf("expected 'ws[1].Role' (%v) to equal 'AllWidgets[1].Role' (%v)", ws[1].Role, AllWidgets[1].Role)
	}
}

func TestGetWidgetBadRequest(t *testing.T) {
	idStr := "hello"

	urlPath, err := router.Get("GetWidget").URL("id", idStr)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	resp, err := http.Get(server.URL + urlPath.String())
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 'resp.StatusCode' (%v) to equal 'http.StatusBadRequest' (%v)", resp.StatusCode, http.StatusBadRequest)
	}
}

func TestGetWidget(t *testing.T) {
	var id int64 = 2
	idStr := fmt.Sprintf("%d", id)

	urlPath, err := router.Get("GetWidget").URL("id", idStr)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	resp, err := http.Get(server.URL + urlPath.String())
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 'resp.StatusCode' (%v) to equal 'http.StatusOK' (%v)", resp.StatusCode, http.StatusOK)
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var widget Widget
	err = decoder.Decode(&widget)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	if widget.Id != AllWidgets[2].Id {
		t.Fatalf("expected 'widget.Id' (%v) to equal 'AllWidgets[2].Id' (%v)", widget.Id, AllWidgets[2].Id)
	}
	if widget.Name != AllWidgets[2].Name {
		t.Fatalf("expected 'widget.Name' (%v) to equal 'AllWidgets[2].Name' (%v)", widget.Name, AllWidgets[2].Name)
	}
	if widget.Role != AllWidgets[2].Role {
		t.Fatalf("expected 'widget.Role' (%v) to equal 'AllWidgets[2].Role' (%v)", widget.Role, AllWidgets[2].Role)
	}
}

func TestPostWidget(t *testing.T) {
	urlPath, err := router.Get("PostWidget").URL()
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	widget := Widget{
		Name: "anotherwidget",
		Role: "controller",
	}

	jsonb, err := json.Marshal(widget)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}
	buf := bytes.NewBuffer(jsonb)

	resp, err := http.Post(server.URL+urlPath.String(), "application/json", buf)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 'resp.StatusCode' (%v) to equal 'http.StatusCreated' (%v)", resp.StatusCode, http.StatusCreated)
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var respWidget Widget
	err = decoder.Decode(&respWidget)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	if respWidget.Id == 0 {
		t.Fatalf("expected respWidget.Id (%v) not to be 0", respWidget.Id)
	}
	if respWidget.Name != widget.Name {
		t.Fatalf("expected 'respWidget.Name' (%v) to equal 'widget.Name' (%v)", respWidget.Name, widget.Name)
	}
	if respWidget.Role != widget.Role {
		t.Fatalf("expected 'respWidget.Role' (%v) to equal 'widget.Role' (%v)", respWidget.Role, widget.Role)
	}
}
