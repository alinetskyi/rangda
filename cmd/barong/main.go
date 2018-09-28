package main

import (
	"net/http"

	"github.com/alexflint/go-arg"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/openware/barong/pkg/log"
)

func main() {
	var args struct {
		Listen string
	}

	args.Listen = ":80"
	arg.MustParse(&args)

	router := chi.NewRouter()
	handler := NewHandler()

	setupRoutes(router, handler)

	http.ListenAndServe(args.Listen, router)
}

func setupRoutes(router chi.Router, handler *Handler) {
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger: log.Logger,
	}))

	router.Get("/api/v1/auth", handler.HandleAuth)
}
