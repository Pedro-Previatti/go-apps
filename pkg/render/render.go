package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Pedro-Previatti/go-apps/pkg/config"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate is a function to open and render go templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template

	// not to use in production, only needed to dev mode to test changes in the app
	if app.UseCache {
		// get template cache from cache config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested t, template from cache
	t, ok := tc[tmpl] // get template from the cache with its name
	if !ok {
		log.Fatal("Could not get template from template cache") // if ok == false then we kill application
	}

	// in the template cache
	buffer := new(bytes.Buffer) // var to hold bytes

	_ = t.Execute(buffer, nil)

	// render template
	_, err := buffer.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCache function creates a cache to store all templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	// cache := make(map[string]*template.Template)
	cache := map[string]*template.Template{} // other way to create a map for cache

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		// if we print page up to this point it will log the path so we get the filename
		name := filepath.Base(page) // .Base strips the path and leaves the last argument(name)
		// ts: template set, type *template.Template
		ts, err := template.New(name).ParseFiles(page) // populate ts parsing file
		if err != nil {
			return cache, err
		}

		// gets a slice of strings with all the layouts
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		// if we have layouts they have to be handled here
		if len(matches) > 0 {
			// adding the layout .tmpl files to the template set
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		// adds all together ot the template map cache
		cache[name] = ts
	}

	return cache, nil
}
