package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/openware/rails5session-go"
	"github.com/openware/rangda/pkg/log"
	"github.com/reconquest/karma-go"
)

const (
	CookieSession = "_barong_session"
)

type Server struct {
	chi.Router
	encryption *rails5session.Encryption
}

func NewServer(config *Config) *Server {
	encryption := rails5session.NewEncryption(
		config.Session.SecretKeyBase,
		rails5session.DefaultEncryptedCookieSalt,
		rails5session.DefaultEncryptedSignedCookieSalt,
	)

	return &Server{
		// Router is an interface, need to create a *chi.Mux
		Router: chi.NewRouter(),

		encryption: encryption,
	}
}

func (server *Server) SetupRoutes() {
	server.Use(middleware.RequestID)
	server.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger: log.Logger,
	}))

	server.Get("/api/v1/auth", server.HandleAuth)
}

func (server *Server) HandleAuth(
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
