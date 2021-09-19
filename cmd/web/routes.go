package main

import (
	"github.com/Tomonaga007/bookings/internal/config"
	"github.com/Tomonaga007/bookings/internal/handlers"
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
	mux.Post("/search-availability-json", http.HandlerFunc(handlers.Repo.AvailabilityJSON))
	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))
	mux.Get("/make-reservation", http.HandlerFunc(handlers.Repo.Reservation))

	mux.Post("/make-reservation", http.HandlerFunc(handlers.Repo.PostReservation))
	mux.Post("/search-availability", http.HandlerFunc(handlers.Repo.PostAvailability))

	fileSrever := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileSrever))

	return mux
}