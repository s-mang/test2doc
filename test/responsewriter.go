package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"runtime"
)

type ResponseWriter struct {
	HandlerInfo HandlerInfo
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
	pc, file, _, ok := runtime.Caller(3)
	if !ok {
		// TODO: handle this better?
		log.Println("setHandlerInfo: !ok")
		return
	}

	fn := runtime.FuncForPC(pc)

	rw.HandlerInfo = HandlerInfo{
		FileName: file,
		FuncName: fn.Name(),
	}
}
