package test

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/adams-sarah/test2doc/doc"
	"github.com/adams-sarah/test2doc/doc/parse"
)

// resources = map[uri]Resource
var resources = map[string]*doc.Resource{}

type Server struct {
	*httptest.Server
	doc *doc.Doc
}

// TODO: filter out 404 responses
func NewServer(handler http.Handler, pkgDir string) (s *Server, err error) {
	// check if url var extractor func is set
	if parse.Extractor == nil {
		panic("please set a URLVarExtractor.")
	}

	outDoc, err := doc.NewDoc(pkgDir)
	if err != nil {
		return s, err
	}

	httptestServer := httptest.NewServer(handleAndRecord(handler, outDoc))

	return &Server{
		httptestServer,
		outDoc,
	}, nil
}

func (s *Server) Finish() {
	s.Close()

	for _, r := range resources {
		s.doc.AddResource(r)
	}

	err := s.doc.Write()
	if err != nil {
		panic(err.Error())
	}
}

func handleAndRecord(handler http.Handler, outDoc *doc.Doc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// copy request body into Request object
		docReq, err := doc.NewRequest(req)
		if err != nil {
			log.Println("Error:", err.Error())
			return
		}

		// record response
		rw := httptest.NewRecorder()
		resp := NewResponseWriter(rw)

		handler.ServeHTTP(resp, req)

		// setup resource
		u := doc.NewURL(req)
		path := u.ParameterizedPath

		if resources[path] == nil {
			resources[path] = doc.NewResource(u)
		}

		// store response body in Response object
		docResp := doc.NewResponse(resp.W)

		// find action
		action := resources[path].FindAction(req.Method)
		if action == nil {
			// make new action
			action, err = doc.NewAction(req.Method, resp.HandlerInfo.FuncName)
			if err != nil {
				log.Println("Error:", err.Error())
				return
			}

			// add Action to Resource's list of Actions
			resources[path].AddAction(action)
		}

		// add request, response to action
		action.AddRequest(docReq, docResp)

		// copy response over to w
		doc.CopyHeader(w.Header(), resp.Header())
		w.WriteHeader(resp.W.Code)
		w.Write(resp.W.Body.Bytes())
	}
}
