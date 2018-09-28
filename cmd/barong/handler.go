package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/openware/barong/pkg/log"
	"github.com/reconquest/karma-go"
)

const (
	CookieSession = "_barong_session"
)

type Handler struct {
	router chi.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) HandleAuth(
	writer http.ResponseWriter, request *http.Request,
) {
	context := karma.
		Describe("remote", request.RemoteAddr).
		Describe("request_id", middleware.GetReqID(request.Context()))

	log.Infof(context, "handling auth request")

	writer.WriteHeader(http.StatusOK)
}
