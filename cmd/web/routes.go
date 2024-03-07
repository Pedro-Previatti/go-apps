package main

import (
	"net/http"

	"github.com/Pedro-Previatti/go-apps/pkg/config"
	"github.com/Pedro-Previatti/go-apps/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// routes contains all application routes
func routes(app *config.AppConfig) http.Handler {
	// creating a new instance of the chi router
	mux := chi.NewRouter()

	// attach the Recoverer middleware to recover from panics
	mux.Use(middleware.Recoverer)
	// attach the nosurf middleware
	mux.Use(NoSurf)
	// attach the session load middleware
	mux.Use(SessionLoad)

	// defining routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// setting up file server
	fs := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
