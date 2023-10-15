package api

import (
	"bytes"
	"net/http"
)

type ContextKey string

type MiddlewareFunc func(next http.HandlerFunc) http.HandlerFunc

type Middleware func(next http.Handler) http.Handler

type Logger interface {
	Println(v ...interface{})
}

type loggerMW struct {
	next   http.Handler
	logger Logger
}

type customRW struct {
	rw         http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

type recoverMW struct {
	next http.Handler
}
