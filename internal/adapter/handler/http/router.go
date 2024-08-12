package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Router struct {
	*chi.Mux
}

func NewRouter(healthHandler HealthHandler, userHandler UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {
		r.Get("/health", healthHandler.Health)
	})

	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", userHandler.Register)
		})
	})

	return r
}
