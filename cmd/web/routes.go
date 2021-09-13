package main

import (
	"github.com/Tomonaga007/bookings/pkg/config"
	"github.com/Tomonaga007/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"

)

func routes(app *config.AppConfig) http.Handler{
	mux := chi.NewRouter()
	mux.Use(SessionLoad)
	mux.Use(NoSurf)


	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/apartment-1", http.HandlerFunc(handlers.Repo.Apartment_1))
	mux.Get("/apartment-2", http.HandlerFunc(handlers.Repo.Apartment_2))
	mux.Get("/search-availability", http.HandlerFunc(handlers.Repo.Availability))
	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))
	mux.Get("/make-reservation", http.HandlerFunc(handlers.Repo.Reservation))

	fileSrever := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileSrever))

	return mux
}