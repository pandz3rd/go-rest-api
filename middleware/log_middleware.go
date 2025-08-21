package middleware

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

type LogMiddleware struct {
	handler http.Handler
	logger  *logrus.Logger
}

func NewLogMiddleware(handler http.Handler, logger *logrus.Logger) *LogMiddleware {
	return &LogMiddleware{
		handler: handler,
		logger:  logger,
	}
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

const TraceIdKey string = "traceId"

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	traceId := uuid.New().String()
	ctx := context.WithValue(r.Context(), TraceIdKey, traceId)

	// Add to response header (optional, useful for clients)
	w.Header().Set("X-Trace-Id", traceId)

	middleware.logger.WithFields(logrus.Fields{"traceId": traceId}).Info(doubleLine)
	middleware.logger.WithFields(logrus.Fields{"traceId": traceId}).Info(requestLine)
	middleware.logger.WithFields(logrus.Fields{"traceId": traceId}).Infof("%s : %s", r.Method, r.RequestURI)

	reqBodyBytes, _ := io.ReadAll(r.Body)
	middleware.logger.WithFields(logrus.Fields{"traceId": traceId}).Infof("Request Header : %s", r.Header)
	middleware.logger.WithFields(logrus.Fields{"traceId": traceId}).Infof("Request body: %s", string(reqBodyBytes))
	r.Body = io.NopCloser(bytes.NewBuffer(reqBodyBytes))

	crw := newCaptureResponseWriter(w)
	middleware.handler.ServeHTTP(crw, r.WithContext(ctx))
	end := time.Since(start).Milliseconds()
	responseBodyString := strings.TrimSpace(crw.body.String())
	middleware.logger.WithFields(logrus.Fields{"traceId": traceId}).Info(responseLine)
	middleware.logger.WithFields(logrus.Fields{"traceId": traceId}).Infof("Duration: %d millisecond", int(end))
	middleware.logger.WithFields(logrus.Fields{"traceId": traceId}).Infof("Response Status: %d", crw.statusCode)
	middleware.logger.WithFields(logrus.Fields{"traceId": traceId}).Infof("Response Header : %s", r.Header)
	middleware.logger.WithFields(logrus.Fields{"traceId": traceId}).Infof("Response Body : %s", responseBodyString)
	middleware.logger.WithFields(logrus.Fields{"traceId": traceId}).Info(doubleLine)
}
