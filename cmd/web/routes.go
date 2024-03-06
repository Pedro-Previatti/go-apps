package main

import (
	"net/http"

	"github.com/Pedro-Previatti/go-apps/pkg/config"
	"github.com/Pedro-Previatti/go-apps/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// creating a new instance of the chi router
	mux := chi.NewRouter()

	// attach the Recoverer middleware to recover from panics
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)

	// defining routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
