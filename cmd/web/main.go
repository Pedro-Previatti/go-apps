package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Pedro-Previatti/go-apps/pkg/config"
	"github.com/Pedro-Previatti/go-apps/pkg/handlers"
	"github.com/Pedro-Previatti/go-apps/pkg/render"
	"github.com/alexedwards/scs/v2"
)

// sets the default port
const portNumber = ":8080"

// app configuration
var app config.AppConfig

// sets the session
var session *scs.SessionManager

// main is the main application function
func main() {

	// change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// create template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// set template cache
	app.TemplateCache = tc
	app.UseCache = false

	// create new repo
	repo := handlers.NewRepo(&app)

	// create new handler
	handlers.NewHandler(repo)

	// set template configuration
	render.NewTemplates(&app)

	fmt.Println("application started")
	fmt.Println("running application on port : " + portNumber)

	// create http server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
