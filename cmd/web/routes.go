package main

import (
	"net/http"

	"github.com/zroygbiv/bnb-res-app/internal/config"
	"github.com/zroygbiv/bnb-res-app/internal/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)

	mux.Get("/about", handlers.Repo.About)

	mux.Get("/kitchen", handlers.Repo.Kitchen)

	mux.Get("/common", handlers.Repo.Common)

	mux.Get("/bedroom", handlers.Repo.Bedroom)

	mux.Get("/search-availability", handlers.Repo.Availability);
	mux.Post("/search-availability", handlers.Repo.PostAvailability);

	mux.Get("/make-reservation", handlers.Repo.Reservation);
	mux.Post("/make-reservation", handlers.Repo.PostReservation);

	mux.Get("/contact", handlers.Repo.Contact);

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
