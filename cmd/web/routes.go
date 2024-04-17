package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vantutran2k1/bookings/pkg/config"
	"github.com/vantutran2k1/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}