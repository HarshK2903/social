package main

import (
	"github.com/HarshK2903/social/internal/health"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRoutes(healthHandler *health.Handler) *chi.Mux {
	r := chi.NewRouter()

	//middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//health handler
	r.Get("/health", healthHandler.Check)
	return r
}
