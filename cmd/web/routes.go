package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"github.com/Tomonaga007/bookings/pkg/config"
	"github.com/Tomonaga007/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler{
	mux := chi.NewRouter()
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}