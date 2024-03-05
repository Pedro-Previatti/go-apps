package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// renderTemplate is a function to open and render go templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create tc, a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err) // kills application if we cannot load the templates
	}

	// get requested t, template from cache
	t, ok := tc[tmpl] // get template from the cache with its name
	if !ok {
		log.Fatal(err) // if ok == false then we kill application
	}

	// this step is created to check if there is any error from the values stored
	// in the template cache
	buffer := new(bytes.Buffer) // var to hold bytes
	err = t.Execute(buffer, nil)
	if err != nil {
		log.Println(err)
	}

	// render template
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// createTemplateCache function creates a cache to store all templates
func createTemplateCache() (map[string]*template.Template, error) {
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
