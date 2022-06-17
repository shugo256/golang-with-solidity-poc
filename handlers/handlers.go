package handlers

import (
	"net/http"
)

func HandlerFuncWithOptions(f func(http.ResponseWriter, *http.Request), opts ...func(http.Handler) http.Handler) http.Handler {
	handler := http.Handler(http.HandlerFunc(f))
	for _, opt := range opts {
		handler = opt(handler)
	}
	return handler
}
