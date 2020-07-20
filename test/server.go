package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"sort"

	"github.com/happyreturns/test2doc/doc"
	"github.com/happyreturns/test2doc/doc/parse"
)

// resources = map[uri]Resource
var resources = map[string]*doc.Resource{}

type ResourceSorter func(map[string]*doc.Resource) []*doc.Resource

type Server struct {
	*httptest.Server
	doc    *doc.Doc
	sortFn ResourceSorter
}

func defaultSorter(resourceMap map[string]*doc.Resource) []*doc.Resource {
	// sort resources by path
	var uris []string
	for k := range resourceMap {
		uris = append(uris, k)
	}
	sort.Strings(uris)

	sortedResources := make([]*doc.Resource, len(uris))
	for i, uri := range uris {
		sortedResources[i] = resources[uri]
	}
	return sortedResources
}

// TODO: filter out 404 responses
func NewServer(handler http.Handler) (s *Server, err error) {
	// check if url var extractor func is set
	if parse.Extractor == nil {
		panic("please set a URLVarExtractor.")
	}

	outDoc, err := doc.NewDoc(".")
	if err != nil {
		return s, err
	}

	httptestServer := httptest.NewServer(handleAndRecord(handler, outDoc))

	return &Server{
		httptestServer,
		outDoc,
		defaultSorter,
	}, nil
}

func (s *Server) SetResourceSorter(sortFunc ResourceSorter) {
	s.sortFn = sortFunc
}

func (s *Server) Finish() {
	s.Close()

	for _, resource := range s.sortFn(resources) {
		s.doc.AddResource(resource)
	}

	if err := s.doc.Write(); err != nil {
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
