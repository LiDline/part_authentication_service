package router

import (
	"test/internal/constants"
	authrouter "test/internal/server/auth"
	"test/internal/server/healthcheck"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func MainRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Mount(constants.HEALTHCHECK, healthcheck.HealthCheckRouter())

	r.Mount(constants.AUTH, authrouter.AuthRouter())

	return r
}
