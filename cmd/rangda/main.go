package main

import (
	"net/http"

	"github.com/alexflint/go-arg"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/openware/rangda/pkg/log"
	"github.com/reconquest/karma-go"
)

var (
	// this variable is changed by runtime ldflags
	version = "[manual build]"
)

type args struct {
	Listen string
	Debug  bool
}

func (args *args) Version() string {
	return version
}

func main() {
	args := &args{}
	args.Listen = ":8080"
	arg.MustParse(args)

	if args.Debug {
		log.SetDebug(true)
	}

	router := chi.NewRouter()
	handler := NewHandler()

	setupRoutes(router, handler)

	log.Infof(
		karma.Describe("address", args.Listen).Describe("version", version),
		"starting listener",
	)

	err := http.ListenAndServe(args.Listen, router)
	if err != nil {
		log.Errorf(err, "unable to start listener at %s", args.Listen)
	}
}

func setupRoutes(router chi.Router, handler *Handler) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger: log.Logger,
	}))

	router.Get("/api/v1/auth", handler.HandleAuth)
}
