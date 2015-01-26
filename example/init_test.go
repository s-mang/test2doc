package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gophergala/test2doc/testdoc"
	"github.com/gophergala/test2doc/testrun"
)

const docFileName = "apiary.apib"

var (
	docFile *testdoc.APIDoc
	server  *httptest.Server
)

type responseExpectation struct {
	StatusCode      int
	BodyContains    []string
	BodyNotContains []string
}

func init() {
	gopath := os.Getenv("GOPATH")

	docFilePath := fmt.Sprintf("%s/src/github.com/gophergala/test2doc/%s", gopath, docFileName)

	desc := &testdoc.APIDescription{
		APIBlueprintFormat: testdoc.APIBlueprintFormat,
		HostURL:            "http://gist.github.com",
		Title:              "Github Gists",
		Description:        "The best place to host your code snippits!",
	}

	docFile = testdoc.NewAPIDoc(docFilePath, desc)

	mux := http.NewServeMux()
	addRoutes(mux)
	server = testrun.NewTestServer(mux, docFile)
}

func helperTestPostRequest(url string, reqBody string, exp *responseExpectation) error {
	var buf bytes.Buffer
	buf.Write([]byte(reqBody))

	res, err := http.Post(url, "application/json", &buf)
	if err != nil {
		return errors.New("Error w/ POST request: " + err.Error())
	}

	return helperTestResponse(res, exp)
}

func helperTestGetRequest(url string, exp *responseExpectation) error {
	res, err := http.Get(url)
	if err != nil {
		return errors.New("Error w/ GET request:" + err.Error())
	}

	return helperTestResponse(res, exp)
}

func helperTestResponse(res *http.Response, exp *responseExpectation) error {
	if res.StatusCode != exp.StatusCode {
		return errors.New(fmt.Sprintf("Expected response code to be %d. Got %d.\n", res.StatusCode, exp.StatusCode))
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		errors.New("Error reading response body:" + err.Error())
	}

	body := string(bodyBytes)

	for _, str := range exp.BodyContains {
		if !strings.Contains(body, str) {
			return errors.New(fmt.Sprintf("Expected response body to include '%s'.", str))
		}
	}

	for _, str := range exp.BodyNotContains {
		if strings.Contains(body, str) {
			return errors.New(fmt.Sprintf("Expected response body NOT to include '%s'.", str))
		}
	}

	return nil
}
