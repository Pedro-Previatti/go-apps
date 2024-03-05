package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pedro-Previatti/go-apps/pkg/config"
	"github.com/Pedro-Previatti/go-apps/pkg/handlers"
	"github.com/Pedro-Previatti/go-apps/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	// app configuration
	var app config.AppConfig

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
