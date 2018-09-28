package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/openware/rangda/pkg/log"
	"github.com/reconquest/karma-go"
)

const (
	CookieSession = "_rangda_session"
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

	cookie, err := request.Cookie(CookieSession)
	if err != nil && err != http.ErrNoCookie {
		log.Errorf(err, "unable to get session cookie: %s", CookieSession)
		return
	}

	if err == http.ErrNoCookie {
		// handle case when there is no cookie
		return
	}

	log.Debugf(nil, "user's cookie: %s", cookie.Value)

	writer.WriteHeader(http.StatusOK)
}
