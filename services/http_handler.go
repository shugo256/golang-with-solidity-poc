package services

import (
	"github.com/rs/zerolog"
	genlog "github.com/shugo256/golang-with-solidity-poc/gen/log"
	"os"

	goahttp "goa.design/goa/v3/http"
	httpmiddleware "goa.design/goa/v3/http/middleware"
	"net/http"
)

type HttpHandler struct {
	goahttp.Muxer
}

type service interface {
	Use(func(handler http.Handler) http.Handler)
	Mount(goahttp.Muxer)
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{Muxer: goahttp.NewMuxer()}
}

func newPrettyLoggerMiddleware(name string) func(handler http.Handler) http.Handler {
	logger := genlog.New(name, true)
	prettyLogger := logger.Logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	logger.Logger = &prettyLogger

	return httpmiddleware.Log(logger)
}

func (h *HttpHandler) AddService(s service, name string) {
	s.Use(newPrettyLoggerMiddleware(name))
	s.Mount(h.Muxer)
}
