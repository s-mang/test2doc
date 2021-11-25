package test

import (
	"github.com/adams-sarah/test2doc/doc/parse"
	"log"
	"net/http"
	"net/http/httptest"
	"runtime"
)

type ResponseWriter struct {
	HandlerInfo HandlerInfo
	URLVars     map[string]string
	W           *httptest.ResponseRecorder
}

type HandlerInfo struct {
	FileName string
	FuncName string
}

func NewResponseWriter(w *httptest.ResponseRecorder) *ResponseWriter {
	return &ResponseWriter{
		W: w,
	}
}

func (rw *ResponseWriter) Header() http.Header {
	return rw.W.Header()
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	rw.setHandlerInfo()
	return rw.W.Write(b)
}

func (rw *ResponseWriter) WriteHeader(c int) {
	rw.W.WriteHeader(c)
}

func (rw *ResponseWriter) setHandlerInfo() {
	i := 1
	max := 15

	var pc uintptr
	var file, fnName string
	var ok, fnInPkg, sawPkg bool

	// iterate until we find the top level func in this pkg (the handler)
	for i < max {
		pc, file, _, ok = runtime.Caller(i)
		if !ok {
			log.Println("test2doc: setHandlerInfo: !ok")
			return
		}

		fn := runtime.FuncForPC(pc)
		fnName = fn.Name()

		fnInPkg = parse.IsFuncInPkg(fnName)
		if sawPkg && !fnInPkg {
			pc, file, _, ok = runtime.Caller(i - 1)
			fn := runtime.FuncForPC(pc)
			fnName = fn.Name()
			break
		}

		sawPkg = fnInPkg
		i++
	}

	rw.HandlerInfo = HandlerInfo{
		FileName: file,
		FuncName: fnName,
	}
}
