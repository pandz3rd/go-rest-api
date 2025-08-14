package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type LogMiddleware struct {
	handler http.Handler
}

func NewLogMiddleware(handler http.Handler) *LogMiddleware {
	return &LogMiddleware{handler: handler}
}

type captureResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func newCaptureResponseWriter(w http.ResponseWriter) *captureResponseWriter {
	return &captureResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		body:           &bytes.Buffer{},
	}
}

func (r *captureResponseWriter) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func (crw *captureResponseWriter) Write(b []byte) (int, error) {
	crw.body.Write(b)
	return crw.ResponseWriter.Write(b)
}

var doubleLine = "===================================================================================="
var requestLine = "---------------------------------- REQUEST -----------------------------------------"
var responseLine = "---------------------------------- RESPONSE ----------------------------------------"

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(doubleLine)
	fmt.Println(requestLine)
	fmt.Printf("%s : %s\n", r.Method, r.RequestURI)

	reqBodyBytes, _ := io.ReadAll(r.Body)
	fmt.Printf("Request Header : %s\n", r.Header)
	fmt.Printf("Request body: %s\n", string(reqBodyBytes))
	r.Body = io.NopCloser(bytes.NewBuffer(reqBodyBytes))

	crw := newCaptureResponseWriter(w)
	middleware.handler.ServeHTTP(crw, r)
	fmt.Println(responseLine)
	fmt.Printf("Response Status: %d\n", crw.statusCode)
	fmt.Printf("Response Header : %s\n", r.Header)
	fmt.Printf("Response Body : %s\n", crw.body.String())
	fmt.Println(doubleLine)
}
