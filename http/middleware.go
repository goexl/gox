package http

import (
	"net/http"
)

type Middleware interface {
	Handle(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc)
}
