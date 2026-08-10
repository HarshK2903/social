package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, rr *http.Request) {
			w.Write([]byte("welcome"))
		})
		r.Get("/health", app.healthCheckHandler)
	})
	return r
}

func (app *application) run(mux *chi.Mux) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 20,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Server Started Running %s", app.config.addr)
	return srv.ListenAndServe()

}
