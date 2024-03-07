package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	// crsf handles security issues
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,                 // prevent cookies from being accessible to JavaScript
		Path:     "/",                  // sets the path
		Secure:   app.InProduction,     // sets the secure flag if app is not running in production
		SameSite: http.SameSiteLaxMode, // send cookies with GET requests
	})

	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
